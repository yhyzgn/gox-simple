// Copyright 2019 yhyzgn enp-simple
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2020-05-14 15:48
// version: 1.0.0
// desc   : 

package remote

import (
	"encoding/json"
	"gox-simple/internal/res"
)

type ResultResolver interface {
	Resolve(bs []byte) (result *res.Response, err error)
}

type jsonResolver struct {
	bean interface{}
}

func (jr *jsonResolver) Resolve(bs []byte) (result *res.Response, err error) {
	if bs != nil && jr.bean != nil {
		r := &res.Response{
			Data: jr.bean,
		}
		// 默认用 res.Response解析
		err = json.Unmarshal(bs, r)
		result = r
	}
	return
}

func newJsonResolver(bean interface{}) *jsonResolver {
	return &jsonResolver{bean: bean}
}
