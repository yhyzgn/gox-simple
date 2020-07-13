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
// time   : 2020-05-13 18:38
// version: 1.0.0
// desc   : 

package util

import (
	"strconv"
	"time"
)

const (
	DatePattern = "2006-01-02 15:04:05.999"
)

// 当前时刻
//
// 东八区
func Now() time.Time {
	return time.Now().In(time.FixedZone("CST", 8*3600))
}

// 格式化日期
func Format(date time.Time) string {
	return date.Format(DatePattern)
}

func NowMillis() int64 {
	return Now().UnixNano() / 1e6
}

func NowUnix() int64 {
	return Now().Unix()
}

func NowUnixStr() string {
	return strconv.FormatInt(NowUnix(), 10)
}
