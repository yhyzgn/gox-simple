package logger

import (
	"github.com/yhyzgn/gog"
)

func Init(profile string) {
	// 无论什么环境，都要在直接控制台输出日志
	if profile == "prod" {
		// 生产环境，添加es日志输出
		gog.SetLevel(gog.INFO)
		gog.SetFormatter(gog.NewJSONFormatter())
		gog.ShortFile(false)
	} else {
		// 其他
		if profile == "test" {
			// 测试环境，添加es日志输出
			gog.SetLevel(gog.DEBUG)
			gog.SetFormatter(gog.NewJSONFormatter())
			gog.ShortFile(false)
		} else {
			// 其他，开发环境，使用默认日志输出
			gog.SetLevel(gog.ALL)
			gog.ShortFile(true)
		}
	}

	gog.CallSkip(2)
	gog.Async(true)
}

func Trace(args ...interface{}) {
	gog.Trace(args...)
}

func TraceF(format string, args ...interface{}) {
	gog.TraceF(format, args...)
}

func Debug(args ...interface{}) {
	gog.Debug(args...)
}

func DebugF(format string, args ...interface{}) {
	gog.DebugF(format, args...)
}

func Info(args ...interface{}) {
	gog.Info(args...)
}

func InfoF(format string, args ...interface{}) {
	gog.InfoF(format, args...)
}

func Warn(args ...interface{}) {
	gog.Warn(args...)
}

func WarnF(format string, args ...interface{}) {
	gog.WarnF(format, args...)
}

func Error(args ...interface{}) {
	gog.Error(args...)
}

func ErrorF(format string, args ...interface{}) {
	gog.ErrorF(format, args...)
}

func Fatal(args ...interface{}) {
	gog.Fatal(args...)
}

func FatalF(format string, args ...interface{}) {
	gog.FatalF(format, args...)
}
