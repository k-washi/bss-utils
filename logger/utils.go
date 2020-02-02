package logger

import (
	"os"
	"strconv"
	"strings"

	"go.uber.org/zap/zapcore"
)

type envVals struct {
	filePath string
	stdout   bool
	level    zapcore.Level
}

//getEnv ログに関する環境変数を設定
func getEnv() (*envVals, error) {
	res := envVals{}
	res.filePath = os.Getenv("LOGGER_FILE_PATH")

	var err error
	res.stdout, err = strconv.ParseBool(os.Getenv("LOGGER_STDOUT"))
	if err != nil {
		res.stdout = true
	}

	level := os.Getenv("LOGGER_LEVEL")
	switch level {
	case "debug":
		res.level = zapcore.DebugLevel
	case "error":
		res.level = zapcore.ErrorLevel
	default:
		res.level = zapcore.InfoLevel
	}

	return &res, nil
}

//openFile ログを出力するファイルを設定する。
//ファイルが存在する場合、ファイルにログを追記。
//ファイルが存在しない場合、ファイルを作成し、ログを出力。
func openFile(fileName string) (*os.File, error) {
	if exists(fileName) {
		f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0777)
		return f, err
	}
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0777)
	return f, err
}

//formatFilePath ログに記載するファイル名の抽出
func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]

}

//exists　ファイルが存在するか確認する。
func exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
