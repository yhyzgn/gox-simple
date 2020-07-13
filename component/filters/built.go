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
// time   : 2020-05-14 11:35
// version: 1.0.0
// desc   : 内置过滤器

package filters

import (
	"github.com/yhyzgn/gox/component/filter"
	"github.com/yhyzgn/gox/util"
	"gox-simple/internal/built"
	"net/http"
)

type BuiltFilter struct{}

func NewBuiltFilter() *BuiltFilter {
	return new(BuiltFilter)
}

func (bf *BuiltFilter) DoFilter(writer http.ResponseWriter, request *http.Request, chain *filter.Chain) {
	util.SetResponseWriterHeader(writer, "X-Built-By", built.FullNameWithBuildDate)
	util.SetResponseWriterHeader(writer, "X-Start-At", built.StartedAt)

	chain.DoFilter(writer, request)
}
