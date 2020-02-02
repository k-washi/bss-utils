package logger

import (
	"fmt"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

/*zap
https://github.com/uber-go/zap#performance
*/

var (
	Log *zap.Logger
)

func init() {
	envVals, err := getEnv()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	outputPaths := []string{}
	if envVals.filePath != "" {
		outputPaths = append(outputPaths, envVals.filePath)
	}

	if envVals.stdout {
		outputPaths = append(outputPaths, "stdout")
	}

	logConfig := zap.Config{
		OutputPaths: outputPaths,
		Level:       zap.NewAtomicLevelAt(envVals.level),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			CallerKey:    "caller",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	if Log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

/*
//Debug debug msg
func Debug(msg string, tags ...zap.Field) {
	log.Debug(msg, tags...)
	log.Sync()
}

//Info info msg
func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	log.Sync() //バッファリングされたログエントリをflushする
}

//Error error msg
func Error(msg string, err error, tags ...zap.Field) {
	if err != nil {
		tags = append(tags, zap.NamedError("error", err))
	}
	log.Error(msg, tags...)
	log.Sync()
}
*/
