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
// time   : 2020-05-14 9:12
// version: 1.0.0
// desc   : 

package res

import (
	xUtil "github.com/yhyzgn/gox/util"
	"gox-simple/util"
	"net/http"
)

const (
	success = 0 // 成功
	failed  = 1 // 失败
)

type Response struct {
	Code      int         `json:"errorcode"`
	Message   string      `json:"msg"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
	Ret       int         `json:"ret"`
}

func SuccessData(data interface{}) *Response {
	return &Response{
		Code:      success,
		Message:   "操作成功",
		Data:      data,
		Timestamp: util.NowUnix(),
		Ret:       success,
	}
}

func SuccessMsg(msg string) *Response {
	return &Response{
		Code:      success,
		Message:   msg,
		Timestamp: util.NowUnix(),
		Ret:       success,
	}
}

func Success(data interface{}, msg string) *Response {
	return &Response{
		Code:      success,
		Message:   msg,
		Data:      data,
		Timestamp: util.NowUnix(),
		Ret:       success,
	}
}

func Failed(msg string) *Response {
	return &Response{
		Code:      failed,
		Message:   msg,
		Timestamp: util.NowUnix(),
		Ret:       failed,
	}
}

func (r *Response) Response(writer http.ResponseWriter) {
	xUtil.ResponseJSON(writer, r)
}
