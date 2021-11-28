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
	"context"
	"log"

	"github.com/coolstina/fishserver"

	"github.com/gin-gonic/gin"
)

func Server(ctx context.Context, cancel context.CancelFunc, address string) {
	engine := gin.Default()

	users := NewMockUsers()

	group := engine.Group("/users")
	group.GET("", users.GetUsers)
	group.POST("", users.AddUser)

	ops := []fishserver.Option{
		fishserver.WithCancelFunc(cancel),
		fishserver.WithContext(ctx),
	}

	err := fishserver.NewServer(address, ops...).SetHandler(engine).Run()
	if err != nil {
		log.Printf("failed to server: %+v\n", err)
	}
}
