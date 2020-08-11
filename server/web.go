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
// time   : 2020-05-13 18:42
// version: 1.0.0
// desc   : Web应用配置

package server

import (
	"github.com/yhyzgn/gox/component/filter"
	"github.com/yhyzgn/gox/component/interceptor"
	"github.com/yhyzgn/gox/ctx"
	"github.com/yhyzgn/gox/of/adapter"
	"gox-simple/component/filters"
	"gox-simple/component/interceptors"
	"gox-simple/component/resolver"
	"net/http"
)

type webConfig struct {
	adapter.WebConfig
	contextPath string
}

func newWebConfig(contextPath string) *webConfig {
	return &webConfig{
		contextPath: contextPath,
	}
}

func (wc *webConfig) Context(ctx *ctx.GoXContext) {
	ctx.
		SetContextPath(wc.contextPath).
		SetNotFoundHandler(func(writer http.ResponseWriter, request *http.Request) {
			http.Error(writer, "服务器度假去啦~", http.StatusNotFound)
		}).
		SetErrorResolver(resolver.NewErrorResolver())
}

func (wc *webConfig) ConfigFilter(chain *filter.Chain) {
	chain.AddFilters("/**", filters.NewBuiltFilter(), filters.NewLogFilter())
}

func (wc *webConfig) ConfigInterceptor(register *interceptor.Register) {
	register.AddInterceptors("/mini/**", interceptors.NewAuthInterceptor())
}
