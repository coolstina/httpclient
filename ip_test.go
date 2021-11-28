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

func TestIPv4ToInt64(t *testing.T) {
	grids := []struct {
		ipv4     string
		expected int64
	}{
		{
			ipv4:     "127.0.0.1",
			expected: 2130706433,
		},
		{
			ipv4:     "10.10.10.10",
			expected: 168430090,
		},
		{
			ipv4:     "0.0.0.0",
			expected: 0,
		},
		{
			ipv4:     "192.168.0.192",
			expected: 3232235712,
		},
		{
			ipv4:     "192.168.0.103",
			expected: 3232235623,
		},
		{
			ipv4:     "192.168.0.104",
			expected: 3232235624,
		},
	}

	for _, grid := range grids {
		actual := IPv4ToInt64(grid.ipv4)
		assert.Equal(t, grid.expected, actual,
			"want: %d, but got: %d\n", grid.expected, actual)
	}
}

func TestInt64ToIPv4(t *testing.T) {
	grids := []struct {
		ipv4     int64
		expected string
	}{
		{
			ipv4:     2130706433,
			expected: "127.0.0.1",
		},
		{
			ipv4:     168430090,
			expected: "10.10.10.10",
		},
		{
			ipv4:     0,
			expected: "0.0.0.0",
		},
		{
			ipv4:     3232235712,
			expected: "192.168.0.192",
		},
		{
			ipv4:     4294967040,
			expected: "192.168.0.192",
		},
	}

	for _, grid := range grids {
		actual := Int64ToIPv4(grid.ipv4)
		assert.Equal(t, grid.expected, actual,
			"want: %d, but got: %d\n", grid.expected, actual)
	}
}
