package app

import (
	"github.com/gaochuwuhan/goutils/pkg/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response 设定统一response的格式
func (g *Gin) Response(httpCode, errCode int, msg string,data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}




func OkWithDetailed(msg string,data interface{}, c *gin.Context) {
	appG := Gin{C: c}
	appG.Response(200, e.SUCCESS,msg,data) //200，有数据的返回，将获取的数据传入即可
}

func OkWithMsg(msg string,c *gin.Context) {
	appG := Gin{C: c}
	appG.Response(200, e.SUCCESS,msg,map[string]interface{}{}) //200，无数据的返回，data字段为{}
}

func OkWithCreated(msg string,c *gin.Context) {
	appG := Gin{C: c}
	appG.Response(201, e.SUCCESS,msg,map[string]interface{}{}) //201，创建成功，data字段为{}
}

func OkWithNoContent(msg string,c *gin.Context) {
	appG := Gin{C: c}
	appG.Response(204, e.SUCCESS,msg,map[string]interface{}{}) //204
}

func FailWithMsg(msg string,c *gin.Context) {
	appG := Gin{C: c}
	appG.Response(200,e.ERROR,msg,map[string]interface{}{}) //200，请求成功但是出现错误 无数据
}

func FailNotFound(msg string,c *gin.Context) {
	appG := Gin{C: c}
	appG.Response(404,e.INVALID_PARAMS,msg,map[string]interface{}{}) //404，无该数据
}

func FailBadRequestList(msg string,c *gin.Context){
	appG := Gin{C: c}
	appG.Response(400,e.INVALID_PARAMS,msg,[]string{}) //400，get list的error返回
}

func FailBadRequest(msg string,c *gin.Context){
	appG := Gin{C: c}
	appG.Response(400,e.INVALID_PARAMS,msg,map[string]interface{}{})  //400，获取单个数据的error返回
}

func FailWithIllegal(msg string,c *gin.Context) {
	Reload:=make(map[string]bool)
	Reload["reload"]=true
	appG := Gin{C: c}
	appG.Response(401,e.ERROR,msg,Reload) //401，认证失败
}

func FailWithWrongPermission(msg string,c *gin.Context) {
	Reload:=make(map[string]bool)
	Reload["reload"]=true
	appG := Gin{C: c}
	appG.Response(403,e.ERROR,msg,Reload) //403，鉴权失败
}

func FailedServerInsideError(msg string,c *gin.Context){
	appG := Gin{C: c}
	appG.Response(500,e.ERROR,msg,map[string]interface{}{})
}

var HttpOpreateMap = map[string]string{
	"POST": "add",
	"PUT": "update",
	"DELETE": "delete",
}