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

	"github.com/stretchr/testify/suite"
)

func TestURLSuite(t *testing.T) {
	suite.Run(t, new(URLSuite))
}

type URLSuite struct {
	suite.Suite
}

func (suite *URLSuite) Test_FixedURL() {
	url, err := FixedURL("://localhost:8000/users/id/22?username=helloshaohua&sex=male")
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), url)
}
