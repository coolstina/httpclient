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

package rawquery

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewRawQueryWithMap(t *testing.T) {
	grids := []struct {
		queries  []Query
		expected string
	}{
		{
			queries: []Query{
				{
					Field: "username",
					Value: "helloshaohua",
				},
				{
					Field: "sex",
					Value: "male",
				},
				{
					Field: "age",
					Value: 18,
				},
				{
					Field: "sleep",
					Value: int(time.Microsecond * 30),
				},
			},
			expected: `username=helloshaohua&sex=male&age=18&sleep=30000`,
		},
	}

	for _, grid := range grids {
		actual := NewRawQueryWithQueries(grid.queries)
		assert.Equal(t, grid.expected, actual)
	}
}

func TestParseRawQuery(t *testing.T) {
	grids := []struct {
		rawquery string
		expected []Query
	}{
		{
			rawquery: `username=helloshaohua&sex=male&age=18&sleep=30000`,
			expected: []Query{
				{
					Field: "username",
					Value: "helloshaohua",
				},
				{
					Field: "sex",
					Value: "male",
				},
				{
					Field: "age",
					Value: "18",
				},
				{
					Field: "sleep",
					Value: "30000",
				},
			},
		},
	}

	for _, grid := range grids {
		actual := ParseRawQuery(grid.rawquery)
		assert.Equal(t, grid.expected, actual)
	}
}

func TestMergeURLRawQuery(t *testing.T) {
	grids := []struct {
		rawurl   string
		rawquery string
		expected string
	}{
		{
			rawurl:   `https://www.google.com/search?q=hello+world&sex=male&sourceid=chrome&ie=UTF-8`,
			rawquery: `username=helloshaohua&sex=male&age=18&sleep=30000`,
			expected: `q=hello+world&sex=male&sourceid=chrome&ie=UTF-8&username=helloshaohua&age=18&sleep=30000`,
		},
	}

	for _, grid := range grids {
		actual, err := MergeURLRawQuery(grid.rawurl, grid.rawquery)
		assert.NoError(t, err)
		assert.Equal(t, grid.expected, actual)
	}
}
