// Copyright 2021 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestServerAPISuite(t *testing.T) {
	suite.Run(t, new(ServerAPISuite))
}

type ServerAPISuite struct {
	suite.Suite
	api *MockUsers
	ctx *gin.Context
	rec *httptest.ResponseRecorder
}

func (s *ServerAPISuite) BeforeTest(suiteName, testName string) {
	s.api = NewMockUsers()
	s.rec = httptest.NewRecorder()
	s.ctx, _ = gin.CreateTestContext(s.rec)
}

func (s *ServerAPISuite) AfterTest(suiteName, testName string) {
	s.api = nil
	s.rec = nil
	s.ctx = nil
}

func (s *ServerAPISuite) Test_GetUsers() {
	expected := `[{"username":"helloshaohua","sex":"male","mobile":"+8613700000000"},{"username":"zhangsan","sex":"male","mobile":"+8613700000001"},{"username":"kitty","sex":"female","mobile":"+8613700000002"}]`
	s.ctx.Request = httptest.NewRequest(http.MethodGet, "/users", nil)
	s.ctx.Request.Header.Set("Content-Type", "application/json")
	s.api.GetUsers(s.ctx)

	assert.Equal(s.T(), http.StatusOK, s.rec.Code)
	data, err := ioutil.ReadAll(s.rec.Body)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expected, fmt.Sprintf("%s", data))
}

func (s *ServerAPISuite) Test_AddUser_BadRequest() {
	params := `{"username":"helloshaohua","sex":"male"}`
	s.ctx.Request = httptest.NewRequest(http.MethodGet, "/test/get", strings.NewReader(params))
	s.ctx.Request.Header.Set("Content-Type", "application/json")
	s.api.AddUser(s.ctx)

	assert.Equal(s.T(), http.StatusBadRequest, s.rec.Code)
	assert.Contains(s.T(), s.ctx.Errors.String(), "Field validation for 'Mobile' failed on the 'required' tag")
}

func (s *ServerAPISuite) Test_AddUser_Successfully() {
	params := `{"username":"lili","sex":"male","mobile":"+8613700000005"}`
	expected := `{"result":"add user successfully"}`
	s.ctx.Request = httptest.NewRequest(http.MethodGet, "/test/get", strings.NewReader(params))
	s.ctx.Request.Header.Set("Content-Type", "application/json")
	s.api.AddUser(s.ctx)

	assert.Equal(s.T(), http.StatusOK, s.rec.Code)
	data, err := ioutil.ReadAll(s.rec.Body)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expected, fmt.Sprintf("%s", data))
}

func (s *ServerAPISuite) Test_AddUser_And_GetUsers() {
	params := `{"username":"lili","sex":"male","mobile":"+8613700000005"}`
	expectedAddUser := `{"result":"add user successfully"}`
	s.ctx.Request = httptest.NewRequest(http.MethodGet, "/test/get", strings.NewReader(params))
	s.ctx.Request.Header.Set("Content-Type", "application/json")
	s.api.AddUser(s.ctx)

	assert.Equal(s.T(), http.StatusOK, s.rec.Code)
	actualAddUser, err := ioutil.ReadAll(s.rec.Body)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedAddUser, fmt.Sprintf("%s", actualAddUser))

	expectedGetUsers := `[{"username":"helloshaohua","sex":"male","mobile":"+8613700000000"},{"username":"zhangsan","sex":"male","mobile":"+8613700000001"},{"username":"kitty","sex":"female","mobile":"+8613700000002"},{"username":"lili","sex":"male","mobile":"+8613700000005"}]`
	s.ctx.Request = httptest.NewRequest(http.MethodGet, "/users", nil)
	s.ctx.Request.Header.Set("Content-Type", "application/json")
	s.api.GetUsers(s.ctx)

	assert.Equal(s.T(), http.StatusOK, s.rec.Code)
	actualGetUsers, err := ioutil.ReadAll(s.rec.Body)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedGetUsers, fmt.Sprintf("%s", actualGetUsers))
}
