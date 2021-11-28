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
	"net/http"

	"github.com/gin-gonic/gin"
)

type MockUsers struct {
	users userStore
}

func (mock *MockUsers) GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, mock.users)
}

func (mock *MockUsers) AddUser(ctx *gin.Context) {
	var u *user
	if err := ctx.BindJSON(&u); err != nil {
		return
	}

	err := mock.users.add(u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"result": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"result": "add user successfully"})
}

func NewMockUsers() *MockUsers {
	store := userStore{
		&user{
			Username: "helloshaohua",
			Sex:      "male",
			Mobile:   "+8613700000000",
		},
		&user{
			Username: "zhangsan",
			Sex:      "male",
			Mobile:   "+8613700000001",
		},
		&user{
			Username: "kitty",
			Sex:      "female",
			Mobile:   "+8613700000002",
		},
	}

	return &MockUsers{users: store}
}
