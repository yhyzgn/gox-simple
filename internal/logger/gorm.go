// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2020-04-09 18:27
// version: 1.0.0
// desc   : 

package logger

import (
	"github.com/yhyzgn/gog"
)

type Gorm struct {
}

func NewGorm() *Gorm {
	return &Gorm{}
}

func (g *Gorm) Print(v ...interface{}) {
	gog.Info(v...)
}
