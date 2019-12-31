package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

// どこからでも使えるようにグローバルで宣言
var log = logrus.New()

func main() {
	logSetting()
	log.Println("ok!")
}

func logSetting() {
	filename := "debug.log"

	//debug用、ログファイルを実行の度に削除
	if _, err := os.Stat(filename); err == nil {
		os.Remove(filename)
	}

	errorLogFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}
	log.Out = errorLogFile
	log.Formatter = &logrus.JSONFormatter{}
}
