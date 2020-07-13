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
// time   : 2020-05-14 9:09
// version: 1.0.0
// desc   : 鉴权拦截器

package interceptors

import (
	"encoding/json"
	"github.com/yhyzgn/gox/common"
	"github.com/yhyzgn/gox/of/adapter"
	"gox-simple/internal/cst"
	"gox-simple/internal/model"
	"gox-simple/internal/res"
	"net/http"
	"net/url"
)

type AuthInterceptor struct {
	adapter.Interceptor
}

func NewAuthInterceptor() *AuthInterceptor {
	return new(AuthInterceptor)
}

func (li *AuthInterceptor) PreHandle(writer http.ResponseWriter, request *http.Request, handler common.Handler) (bool, *http.Request, http.ResponseWriter) {
	tioAuth := request.Header.Get("tioUnionAuth")
	if tioAuth == "" {
		res.Failed("无访问权限").Response(writer)
		return false, request, writer
	}
	// 解码
	tioAuth, err := url.QueryUnescape(tioAuth)
	if err != nil {
		res.Failed("权限数据解码错误：" + err.Error()).Response(writer)
		return false, request, writer
	}

	auth := &model.Authority{}
	err = json.Unmarshal([]byte(tioAuth), auth)
	if err != nil {
		res.Failed("权限数据解码错误：" + err.Error()).Response(writer)
		return false, request, writer
	}

	// 将权限信息通过request传递
	request = li.SetReqAttr(request, cst.StaffAuthority, auth)
	return true, request, writer
}
