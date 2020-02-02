package logger

import (
	"errors"
	"testing"

	"go.uber.org/zap/zapcore"

	"github.com/stretchr/testify/assert"
)

const (
	logPath  = "./log.txt"
	stdout   = "true"
	logLevel = "debug"
)

//LOGGER_FILE_PATH=./log.txt LOGGER_STDOUT=true LOGGER_LEVEL=debug go test ./logger/...
func TestGetEnv(t *testing.T) {

	envVals, err := getEnv()
	assert.Nil(t, err)
	assert.Equal(t, envVals.filePath, logPath)
	assert.True(t, envVals.stdout)
	assert.Equal(t, envVals.level, zapcore.DebugLevel)
}

func TestLoggin(t *testing.T) {
	msg := "test msg"

	Log.Debug(msg)
	Log.Info(msg)
	err := errors.New("test error")
	Log.Error(err.Error())

	assert.True(t, exists(logPath))

}
