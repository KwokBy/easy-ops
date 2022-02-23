package zlog

import (
	"fmt"
	"os"
	"time"

	"github.com/KwokBy/easy-ops/pkg/file"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zlog *zap.SugaredLogger

// 初始化zap日志配置
func init() {
	// TODO 路径后续接入配置文件
	// 判断是否有Director文件夹
	if ok, _ := file.PathExists("./log"); !ok {
		fmt.Printf("create %v directory\n", "./zlog")
		_ = os.Mkdir("./log/", os.ModePerm)
	}
	// 区分多种日志
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.ErrorLevel
	})
	core := [...]zapcore.Core{
		// TODO 路径后续接入配置文件
		getEncoderCore("./log/server_debug.log", debugPriority),
		getEncoderCore("./log/server_info.log", infoPriority),
		getEncoderCore("./log/server_warn.log", warnPriority),
		getEncoderCore("./log/server_error.log", errorPriority),
	}
	logger := zap.New(zapcore.NewTee(core[:]...), zap.AddCaller())
	logger = logger.WithOptions(zap.AddCaller())
	zlog = logger.Sugar()
}

func Debug(args ...interface{}) {
	zlog.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	zlog.Debugf(template, args...)
}

func Info(args ...interface{}) {
	zlog.Info(args...)
}

func Infof(template string, args ...interface{}) {
	zlog.Infof(template, args...)
}

func Warn(args ...interface{}) {
	zlog.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	zlog.Warnf(template, args...)
}

func Error(args ...interface{}) {
	zlog.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	zlog.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	zlog.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	zlog.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	zlog.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	zlog.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	zlog.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	zlog.Fatalf(template, args...)
}

// 	使用默认的编码器
func getEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// 自定义日志编码器配置
func getEncoderConfig() zapcore.EncoderConfig {
	config := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	return config
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	// 使用lumberjack进行日志分割
	// TODO 是否开启控制台log后续接入配置文件
	writer := GetLogWriter(fileName, true)
	return zapcore.NewCore(getEncoder(), writer, level)
}

// GetLogWriter lumberjack进行日志分割
func GetLogWriter(file string, consoleLog bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	if consoleLog {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}
