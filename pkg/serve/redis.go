package serve

import (
	"fmt"
	"github.com/gaochuwuhan/goutils"
	"github.com/gaochuwuhan/goutils/pkg/cafe"
	"github.com/go-redis/redis"
	"time"
)


type Redis struct{

}

func (re *Redis) ClientInit() *redis.Client{
	env:=goutils.VP.GetString("env")
	addr:=goutils.VP.GetString(cafe.JoinStr(env,".redisconn"))
	pass:=goutils.VP.GetString(cafe.JoinStr(env,".redispasswd"))
	db:=goutils.VP.GetInt(cafe.JoinStr(env,".redisdb"))
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass, // no password set
		DB:       db,        // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(pong,err)
	}
	return client
}

func (re *Redis) Set(key string,value interface{},expiration time.Duration)  error{
	/*redis的set方法，新建一个hash数据
		key：redis某个 key，
		value：为改key赋值，
		expiration：存在redis的超时时间
	*/

	client := re.ClientInit()
	err := client.Set(key,value, expiration).Err()
	if err != nil{
		return err
	}
	return nil

}

func (re *Redis) Drop(key string) error{
	err := re.ClientInit().Del(key).Err()
	return err
}

