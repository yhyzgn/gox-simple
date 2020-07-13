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
// time   : 2020-05-14 16:59
// version: 1.0.0
// desc   : 

package remote

import (
	"gox-simple/internal/logger"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestClient_JSON(t *testing.T) {
	result, err := POST().
		BaseURL("http://localhost:8099").
		Path("/simple/api/debug/test/post").
		Header("Mobile", "18313889251").
		Cookie("Token", "aabbccdd").
		Form("name", "张三").
		Form("age", 34).
		Response(&User{})
	logger.Info(result.Data)
	logger.Error(err)
}
