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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	var params = Params{
		{
			Key:   "username",
			Value: "helloshaohua",
		},
		{
			Key:   "address",
			Value: "北京",
		},
		{
			Key:   "sex",
			Value: "male",
		},
	}

	params = params.Remove("address")
	assert.Len(t, params, 2)
}

func TestRemoves(t *testing.T) {
	var params = Params{
		{
			Key:   "username",
			Value: "helloshaohua",
		},
		{
			Key:   "address",
			Value: "北京",
		},
		{
			Key:   "sex",
			Value: "male",
		},
	}

	params = params.Removes(ParamKeys{"address", "sex"})
	assert.Len(t, params, 1)
}
