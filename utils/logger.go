package utils

import (
	"log"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {
	writeSyncer := getLogWriter(Config.Log.Filename, Config.Log.MaxSize, Config.Log.MaxBackups, Config.Log.MaxAge)
	encoder := getEncoder()

	var l zapcore.Level
	err := l.UnmarshalText([]byte(Config.Log.Level))
	if err != nil {
		log.Fatalf("failed to init logger, %s", err.Error())
	}

	core := zapcore.NewCore(encoder, writeSyncer, l)
	Logger = zap.New(core, zap.AddCaller())
}

// 获取日志编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 获取日志写入器
func getLogWriter(filename string, maxSize, maxBackups, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize, // MB
		MaxBackups: maxBackups,
		MaxAge:     maxAge, // days
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
