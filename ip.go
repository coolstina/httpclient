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
	"math/big"
	"net"
)

// IPv4ToInt64 Convert IPv4 string to IP int64.
func IPv4ToInt64(ip string) int64 {
	return big.NewInt(0).SetBytes(net.ParseIP(ip).To4()).Int64()
}

// Int64ToIPv4 Convert IPv4 int64 to IPv4 string.
func Int64ToIPv4(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}
