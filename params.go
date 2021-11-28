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

// Param is a single URL parameter, consisting of a key and a value.
type Param struct {
	Key   string
	Value string
}

// ParamKeys is parameter Key slice. Used to delete multiple keys.
type ParamKeys []string


// Exists Check whether the Key parameter is saved in the ParamKeys slice.
func (pk ParamKeys) Exists(name string) bool {
	for _, v := range pk {
		if v == name {
			return true
		}
	}
	return false
}

// Params is a Param-slice, as returned by the router.
// The slice is ordered, the first URL parameter is also the first slice value.
// It is therefore safe to read values by the index.
type Params []Param

// Get returns the value of the first Param which key matches the given name.
// If no matching Param is found, an empty string is returned.
func (ps Params) Get(name string) (string, bool) {
	for _, entry := range ps {
		if entry.Key == name {
			return entry.Value, true
		}
	}
	return "", false
}

// ByName returns the value of the first Param which key matches the given name.
// If no matching Param is found, an empty string is returned.
func (ps Params) ByName(name string) (va string) {
	va, _ = ps.Get(name)
	return
}

func (ps Params) Remove(name string) Params {
	var r = make(Params, 0, len(ps))
	for _, entry := range ps {
		if entry.Key == name {
			continue
		}
		r = append(r, entry)
	}
	return r
}

func (ps Params) Removes(names ParamKeys) Params {
	var r = make(Params, 0, len(ps))
	for _, entry := range ps {
		if names.Exists(entry.Key) {
			continue
		}
		r = append(r, entry)
	}
	return r
}
