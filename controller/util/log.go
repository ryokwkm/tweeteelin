package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()
var filename = "debug.log"

func CreateMyLog() *logrus.Logger {
	var log = logrus.New()

	//debug用、ログファイルを実行の度に削除
	if _, err := os.Stat(filename); err == nil {
		os.Remove(filename)
	}
	logSetting(log)
	return log
}

func logSetting(log *logrus.Logger) {
	errorLogFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}
	log.Out = errorLogFile
	log.Formatter = &logrus.TextFormatter{}
}

func Vr(array interface{}, title string) {
	logSetting(Log)
	if len(title) > 0 {
		Log.Println(title, "\n")
	}
	Log.Printf("%#v ", array)

	Log.Print("")
}
