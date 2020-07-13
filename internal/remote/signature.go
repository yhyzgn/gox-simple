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
// time   : 2020-05-15 16:33
// version: 1.0.0
// desc   : 接口签名

package remote

import (
	"crypto/sha1"
	"encoding/hex"
	"gox-simple/internal/logger"
	"sort"
	"strings"
)

// 签名
func Signature(appID, appKey, rand string) string {
	arr := []string{appID, appKey, rand}
	sort.Strings(arr)
	logger.InfoF("Sorted argument is {} which used create signature token.", arr)
	sum := sha1.Sum([]byte(strings.Join(arr, "")))
	sign := hex.EncodeToString(sum[:])
	logger.InfoF("Signature is {} of arguments [appID = {}, appKey = {}, rand = {}].", sign, appID, appKey, rand)
	return sign
}
