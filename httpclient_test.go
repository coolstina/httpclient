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

package httpclient

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/coolstina/httpclient/test/server"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestHttpClientSuite(t *testing.T) {
	suite.Run(t, new(HttpClientSuite))
}

type HttpClientSuite struct {
	suite.Suite
	getUrl        string
	testServerUrl string
	ctx           context.Context
	cancel        context.CancelFunc
}

func (suite *HttpClientSuite) BeforeTest(suiteName, testName string) {
	suite.getUrl = "https://tpc.googlesyndication.com/daca_images/simgad/8396093335661461013"
	suite.testServerUrl = "localhost:30000"
	suite.ctx, suite.cancel = context.WithCancel(context.Background())

	go server.Server(suite.ctx, suite.cancel, suite.testServerUrl)
}

func (suite *HttpClientSuite) AfterTest(suiteName, testName string) {
	suite.cancel()
}

func (suite *HttpClientSuite) Test_Get_NetworkResource() {
	// Execute HTTP GET request.
	resp, err := Get(suite.getUrl).Do()
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	// Read response body.
	all, err := ioutil.ReadAll(resp.Body)
	assert.NoError(suite.T(), err)

	// Write to file.
	err = ioutil.WriteFile("test/data/test_download_image.png", all, os.ModePerm)
	assert.NoError(suite.T(), err)
}

func (suite *HttpClientSuite) Test_Get_For_QueryParams() {
	var params = Params{
		{
			Key:   "username",
			Value: "helloshaohua",
		},
	}

	resp, err := Get(suite.getUrl).QueryParams(params).Debug(true).Do()
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	all, err := ioutil.ReadAll(resp.Body)
	assert.NoError(suite.T(), err)

	err = ioutil.WriteFile("test/data/test_download_image_query_params.png", all, os.ModePerm)
	assert.NoError(suite.T(), err)
}

func (suite *HttpClientSuite) Test_AddUsers_With_Body() {

	data := `{
		"username": "user1",
		"sex": "female",
		"mobile": "+8613700001000"
	}`

	url := fmt.Sprintf("http://%s/users", suite.testServerUrl)
	resp, err := Post(url).Body(strings.NewReader(data)).Debug(true).Do()
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	actual, err := ioutil.ReadAll(resp.Body)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), actual)
}

func (suite *HttpClientSuite) Test_AddUsers_With_BodyWithJSON() {

	data := `{
		"username": "user1",
		"sex": "female",
		"mobile": "+8613700001000"
	}`

	url := fmt.Sprintf("http://%s/users", suite.testServerUrl)
	resp, err := Post(url).BodyWithJSON(data).Debug(true).Do()
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	actual, err := ioutil.ReadAll(resp.Body)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), actual)
}

func Test_AddUsers_With_Timeout_True(t *testing.T) {
	url := `http://httpbin.org/get`
	resp, err := Get(url).Debug(true).Timeout(time.Duration(time.Millisecond * 80)).Do()
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func Test_AddUsers_With_Timeout_False(t *testing.T) {
	url := `http://httpbin.org/get`
	resp, err := Get(url).Debug(true).Timeout(time.Duration(time.Millisecond * 800)).Do()
	assert.NoError(t, err)
	defer resp.Body.Close()

	actual, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, actual)
}
