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
	"sync"
)

var mutex sync.Mutex

type userStore []*user

func (u *userStore) list() userStore {
	mutex.Lock()
	defer mutex.Unlock()
	return *u
}

func (u *userStore) add(user *user) error {
	mutex.Lock()
	defer mutex.Unlock()

	var already bool
	for _, item := range *u {
		if item.Username == user.Username {
			already = true
		}
	}
	if already {
		return fmt.Errorf("user already exists")
	}

	*u = append(*u, user)

	return nil
}
