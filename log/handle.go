package mylog

import (
	"fmt"
	"time"

	"github.com/yangioc/bk_pack/util"
)

func Info(format interface{}) {
	if Level < Level_Info {
		return
	}

	logFormat := logFormat{
		Level:   Level_Info_str,
		Time:    time.Now().Format(time.RFC3339),
		Message: fmt.Sprint(format),
	}
	println(logFormat)
}

func Infof(format string, a ...interface{}) {
	if Level < Level_Info {
		return
	}

	logFormat := logFormat{
		Level:   Level_Info_str,
		Time:    time.Now().Format(time.RFC3339),
		Message: fmt.Sprintf(format, a...),
	}
	println(logFormat)
}

// func SprintInfof(format string, a ...interface{}) string {
// 	logFormat := logFormat{
// 		Level:   Level_Info_str,
// 		Time:    time.Now().Format(time.RFC3339),
// 		Message: fmt.Sprintf(format, a...),
// 	}
// 	js, _ := fmt.Marshal(logFormat)
// 	return util.Sprint(string(js))
// }

func Error(format interface{}) {
	logFormat := logFormat{
		Level:   Level_Error_str,
		Time:    time.Now().Format(time.RFC3339),
		Message: fmt.Sprint(format),
	}
	println(logFormat)
	// return errors.New(logFormat.Message)
}

func Errorf(format string, a ...interface{}) {
	logFormat := logFormat{
		Level:   Level_Error_str,
		Time:    time.Now().Format(time.RFC3339),
		Message: fmt.Sprintf(format, a...),
	}
	println(logFormat)
	// return errors.New(logFormat.Message)
}

func println(format logFormat) {
	js, _ := util.Marshal(format)
	fmt.Println(string(js))
}
