package mlog

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	mLogger *zap.SugaredLogger
)

// InitLogger 初始化logger
func InitLogger(filePath string) {
	// 获取控制台输出和文件输出的写入器
	consoleWriter := zapcore.Lock(os.Stdout)
	fileWriter := getLogWriter(filePath)

	// 获取编码器
	encoder := getEncoder()

	// 配置控制台输出的核心
	consoleCore := zapcore.NewCore(encoder, consoleWriter, zapcore.DebugLevel)

	// 配置文件输出的核心
	fileCore := zapcore.NewCore(encoder, fileWriter, zapcore.DebugLevel)

	// 合并两个核心
	core := zapcore.NewTee(consoleCore, fileCore)

	// 创建logger
	logger := zap.New(core, zap.AddCaller())
	mLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogWriter 在zap中加入Lumberjack支持
func getLogWriter(filename string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    100,   // 以 MB 为单位
		MaxBackups: 5,     // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxAge:     30,    // 保留旧文件的最大天数
		Compress:   false, // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Error(args interface{}) {
	mLogger.Error(args)
}

func Errorf(format string, args ...interface{}) {
	mLogger.Errorf(format, args...)
}

func Info(args interface{}) {
	mLogger.Info(args)
}

func Infof(format string, args ...interface{}) {
	mLogger.Infof(format, args...)
}
