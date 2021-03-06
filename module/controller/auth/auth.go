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
// time   : 2020-05-15 21:17
// version: 1.0.0
// desc   : 支持鉴权

package auth

import (
	"errors"
	"github.com/yhyzgn/gox/of"
	"gox-simple/internal/cst"
	"gox-simple/internal/model"
	"net/http"
)

type Controller struct {
	of.Controller
}

// 获取权限信息
func (c Controller) Auth(req *http.Request) (*model.Authority, error) {
	auth := c.GetReqAttr(req, cst.StaffAuthority)
	if auth == nil {
		return nil, errors.New("未接收到任何权限信息")
	}
	return auth.(*model.Authority), nil
}
