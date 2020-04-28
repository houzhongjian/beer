package log

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

//Info.
func Info(msg ...string) {
	logPath := GetLogpath()
	//sArr := strings.Split(logPath, "happyApp")
	//logPath = sArr[1]

	logMsg := "\033[32m[INFO]\033[0m [" + logPath + "] " + strings.Join(msg, "")
	println(logMsg)
}

//Info.
func Println(msg ...string) {
	logMsg := "\033[32m[INFO]\033[0m " + strings.Join(msg, "")
	println(logMsg)
}

func Debug(msg interface{}) {
	log.Printf("%+v\n", msg)
}

//Error 错误日志.
func Error(msg ...string) {
	logPath := GetLogpath()
	//sArr := strings.Split(logPath, "happyApp")
	//logPath = sArr[1]

	logMsg := "\033[31m[ERROR]\033[0m [" + logPath + "] " + strings.Join(msg, "")
	println(logMsg)
}

//WARNING .
func WARNING(msg ...string) {
	logPath := GetLogpath()
	//sArr := strings.Split(logPath, "happyApp")
	//logPath = sArr[1]

	logMsg := "\033[33m[WARNING]\033[0m [" + logPath + "] " + strings.Join(msg, "")
	println(logMsg)
}

//getLogpath 获取产生日志的路径.
func GetLogpath() string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s:%d", file, line)
}
