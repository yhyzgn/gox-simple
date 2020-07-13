// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2020-04-09 16:55
// version: 1.0.0
// desc   : 

package logger

import (
	"net/url"
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	jsn := "{\"userInfo\":{\"userCode\":\"exampleUserCode\",\"name\":\"example\",\"mobile\":\"13312345678\",\"code\":\"exampleCode\"},\"depts\":[{\"deptCode\":\"exampleDeptCode\",\"orgCode\":\"exampleOrgCode\",\"orgName\":\"exampleOrgName\",\"areaCode\":\"530000\",\"tagCodes\":[1],\"areaLevel\":\"1\",\"name\":\"\",\"roleCodes\":[\"exampleRoleCode\"]}]}";
	encoded := url.QueryEscape(jsn)
	Info(encoded)

	time.Sleep(3 * time.Second)
}
