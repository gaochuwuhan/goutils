package serve

import (
	"fmt"
	logger2 "github.com/gaochuwuhan/goutils/logger"
	"github.com/gaochuwuhan/goutils/pkg/cafe"
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)






var level zapcore.Level

func Zap() (logger *zap.Logger) {
	if ok, _ := cafe.PathExists(logger2.ZapLogger.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", logger2.ZapLogger.Director)
		_ = os.Mkdir(logger2.ZapLogger.Director, os.ModePerm)
	}

	switch logger2.ZapLogger.Level { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}
	if logger2.ZapLogger.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  logger2.ZapLogger.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case logger2.ZapLogger.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case logger2.ZapLogger.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case logger2.ZapLogger.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case logger2.ZapLogger.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if logger2.ZapLogger.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore() (core zapcore.Core) {
	writer, err := GetWriteSyncer() // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(logger2.ZapLogger.Prefix + "2006/01/02 - 15:04:05.000"))
}


func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(logger2.ZapLogger.Director, "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName(logger2.ZapLogger.LinkName),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if logger2.ZapLogger.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}