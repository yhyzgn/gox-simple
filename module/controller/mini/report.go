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
// time   : 2020-05-15 21:07
// version: 1.0.0
// desc   : 小程序上报接口

package mini

import (
	"github.com/yhyzgn/gox/core"
	"gox-simple/internal/res"
	"gox-simple/module/controller/auth"
	"net/http"
)

type ReportController struct {
	auth.Controller
}

func (rc ReportController) Mapping(mapper *core.Mapper) {
	mapper.Get("/check").HandlerFunc(rc.check).Mapping()
}

// 检查最新的上报状态
func (rc ReportController) check(req *http.Request) (*res.Response, error) {
	ath, err := rc.Auth(req)
	if err != nil {
		return nil, err
	}
	return res.SuccessData(ath), nil
}
