package nsqutils

import (
	"github.com/gaochuwuhan/goutils"
	"github.com/gaochuwuhan/goutils/logger"
	"github.com/gaochuwuhan/goutils/pkg/cafe"
	"github.com/nsqio/go-nsq"
)


var TcpNsqdAddr string

func NsqdInfo() string{
	env := goutils.VP.GetString("env")
	TcpNsqdAddr = goutils.VP.GetString(cafe.JoinStr(env,".tcpNsqdAddr"))
	return TcpNsqdAddr
}

func NsqProducer() (*nsq.Producer,error){
	addr := NsqdInfo()
	config:=nsq.NewConfig()
	producer,err:=nsq.NewProducer(addr,config)
	if err != nil{
		logger.Log.Sugar().Errorf("nsq connection error:%v",err.Error())
		return nil,err
	}
	return producer,nil
}

