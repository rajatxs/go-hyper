package log

import (
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	debugLog *log.Logger
	infoLog  *log.Logger
	warnLog  *log.Logger
	errLog   *log.Logger
	fatalLog *log.Logger
)

func init() {
	var flag int = log.Ltime | log.Ldate

	debugLog = log.New(
		os.Stdout,
		color.YellowString("DEBUG "),
		flag)

	infoLog = log.New(
		os.Stdout,
		color.New(color.FgHiGreen).Sprintf("INFO "),
		flag)

	warnLog = log.New(
		os.Stdout,
		color.New(color.Bold, color.FgHiYellow).Sprintf("WARN "),
		flag)

	errLog = log.New(
		os.Stdout,
		color.New(color.Bold, color.FgRed).Sprintf("ERROR "),
		flag)

	fatalLog = log.New(
		os.Stdout,
		color.New(color.Bold, color.BgRed, color.FgWhite).Sprintf("FATAL")+" ",
		flag)
}

func Debug(v ...interface{}) {
	debugLog.Println(v...)
}
func Debugf(f string, v ...any) {
	debugLog.Printf(f, v...)
}

func Info(v ...interface{}) {
	infoLog.Println(v...)
}
func Infof(f string, v ...any) {
	infoLog.Printf(f, v...)
}

func Fatal(v ...interface{}) {
	fatalLog.Fatal(v...)
}
func Fatalf(f string, v ...any) {
	fatalLog.Fatalf(f, v...)
}

func Warn(v ...interface{}) {
	warnLog.Println(v...)
}
func Warnf(f string, v ...any) {
	warnLog.Printf(f, v...)
}

func Error(v ...interface{}) {
	errLog.Println(v...)
}
func ErrorF(f string, v ...any) {
	errLog.Printf(f, v...)
}
