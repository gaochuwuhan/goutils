package serve



import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func CronServe(freq string,cmd func()) (cron.EntryID, error){
	cr := cron.New(cron.WithSeconds())
	_,err:=cr.AddFunc(freq,cmd)
	if err != nil{
		fmt.Println(err)
	}
	cr.Start()
	defer cr.Stop()
	select {
	//查询语句，保持程序运行，在这里等同于for{}
	}
}



