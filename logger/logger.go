package logger

import (
	"github.com/gaochuwuhan/goutils"
	"github.com/gaochuwuhan/goutils/pkg/cafe"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Log *zap.Logger
)

type ZapFormat struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`
	LinkName      string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"showLine"`
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"`
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`
}

var ZapLogger ZapFormat

func ZapLoggerInit(vp *viper.Viper) {
	ZapLogger = ZapFormat{
		Level:         vp.GetString(cafe.JoinStr(goutils.ENV,".log.level")),
		Format:        "console",
		Prefix:        vp.GetString(cafe.JoinStr(goutils.ENV,".log.prefix")),
		Director:      vp.GetString(cafe.JoinStr(goutils.ENV,".log.dir")),
		LinkName:      "latest_log",
		ShowLine:      true,
		EncodeLevel:   "CapitalColorLevelEncoder",
		StacktraceKey: "stacktrace",
		LogInConsole:  true,
	}

}

