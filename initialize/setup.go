package initialize

import (
	"fmt"
	"github.com/gaochuwuhan/goutils"
	"github.com/gaochuwuhan/goutils/logger"
	"github.com/gaochuwuhan/goutils/pkg/serve"
)

//读取

func InitService(){
	goutils.VP=serve.Viper()
	if goutils.VP == nil{
		fmt.Printf(">>>>>>>goutils VP is nil<<<<")
		panic("viper is nil ptr")
	}
	//当前环境
	goutils.ENV = goutils.VP.GetString("env")
	////初始化zap配置
	logger.ZapLoggerInit(goutils.VP)
	logger.Log=serve.Zap()
	logger.Log.Sugar().Infof("[ =========== current env:%v ===========]",goutils.ENV)
	//初始化全局mysql DB
	goutils.DB=serve.GormMysql()

}
