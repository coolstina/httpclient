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
	"fmt"
	"net/url"
	"strings"
)

func FixedURL(rawurl string) (*url.URL, error) {
	if !strings.HasPrefix(rawurl, "http") || !strings.HasPrefix(rawurl, "https") {
		// Clear scheme connector.
		if !strings.HasPrefix(rawurl, "://") {
			rawurl = fmt.Sprintf("://%s", rawurl)
		}
		rawurl = fmt.Sprintf("http%s", rawurl)
	}

	parse, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	return parse, nil
}
