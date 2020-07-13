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
// time   : 2020-05-13 18:25
// version: 1.0.0
// desc   : 测试

package debug

import (
	"github.com/yhyzgn/gox/core"
	"github.com/yhyzgn/gox/of"
	"gox-simple/internal/res"
	"net/http"
)

type TestController struct {
	of.Controller
}

func (tc TestController) Mapping(mapper *core.Mapper) {
	mapper.
		Get("/").HandlerFunc(tc.test).Mapping().
		Get("/req").HandlerFunc(tc.req).Mapping().
		Get("/hello").HandlerFunc(tc.hello).Mapping().
		Get("/get").HandlerFunc(tc.get).Param("name").Param("age").Mapping().
		Post("/post").HandlerFunc(tc.post).Param("name").Param("age").Mapping().
		Post("/body").HandlerFunc(tc.body).Body("body").Mapping()
}

func (tc TestController) req(req *http.Request) string {
	return "Hello Reporter."
}

func (TestController) test() string {
	return "Hello Reporter."
}

func (TestController) hello() *res.Response {
	return res.SuccessData(map[string]interface{}{
		"name": "张三",
		"age":  1003,
	})
}

func (TestController) get(name string, age int) *res.Response {
	return res.SuccessData(map[string]interface{}{
		"name": name,
		"age":  age,
	})
}

func (TestController) post(name string, age int) *res.Response {
	return res.SuccessData(map[string]interface{}{
		"name": name,
		"age":  age,
	})
}

func (TestController) body(body *map[string]interface{}) *res.Response {
	return res.SuccessData(body)
}
