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
// time   : 2020-05-13 18:33
// version: 1.0.0
// desc   : 

package filters

import (
	"github.com/yhyzgn/gox/component/filter"
	"github.com/yhyzgn/gox/of"
	xUtil "github.com/yhyzgn/gox/util"
	"gox-simple/internal/logger"
	"gox-simple/util"
	"net/http"
)

type LogFilter struct{}

func NewLogFilter() *LogFilter {
	return new(LogFilter)
}

func (*LogFilter) DoFilter(writer http.ResponseWriter, request *http.Request, chain *filter.Chain) {
	rw := of.NewResponseWriter(writer)

	reqBody := xUtil.RecycleRequestBody(request)
	logger.Info(string(reqBody))

	path := request.URL.Path
	start := util.Now()

	chain.DoFilter(rw, request)

	end := util.Now()

	resBody := rw.ResponseBody()
	logger.Info(string(resBody))

	logger.InfoF("path: {}, start: {}, end: {}, lapse: {}, len: {}, reqHeader: {}, resHeader: {}", path, util.Format(start), util.Format(end), end.Sub(start).Milliseconds(), rw.ContentLength(), request.Header, writer.Header())
}
