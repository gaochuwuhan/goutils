package tools

import (
	"crypto/tls"
	logger2 "github.com/gaochuwuhan/goutils/logger"
	"github.com/gaochuwuhan/goutils/pkg/cafe"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"time"
)

func newHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Duration(10 * time.Second),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}

type HttpRequest struct {

}


//发送请求得到response body
func(self *HttpRequest) RetrieveResBody(urladdr, method string, body io.Reader,headers map[string]string) ([]byte,int, error){
	client:=newHttpClient()
	req,err:=http.NewRequest(method,urladdr,body)
	if err != nil {
		return nil,-1,err //-1代表http请求时出现错误
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res,err:=client.Do(req)
	if err != nil {
		return nil,-1,err
	}
	defer res.Body.Close()
	b,_:=io.ReadAll(res.Body)
	return b,res.StatusCode,nil
}

func(self *HttpRequest) RetrieveResBodyByFullParams(urladdr, method,path string,query map[string]string,body io.Reader,headers map[string]string) ([]byte,int,error){
	var (
		pathUrl,fullUrl string
	)
	if path != ""{
		pathUrl=cafe.JoinStr(urladdr,"/",path)
	}else{
		pathUrl=urladdr
	}
	if query != nil{
		va:=make(url.Values)
		for k,v := range query{
			va.Add(k,v)
		}
		queryStr:=va.Encode()
		fullUrl=cafe.JoinStr(pathUrl,"?",queryStr)
	}else{
		fullUrl=pathUrl
	}
	client:=newHttpClient()
	//fmt.Println("======url:",fullUrl)
	req,err:=http.NewRequest(method,fullUrl,body)
	if err != nil {
		logger2.Log.Sugar().Errorf("!!!",zap.Error(err))
		return nil,-1,err //-1代表http请求时出现错误
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res,err:=client.Do(req)
	if err != nil {
		logger2.Log.Sugar().Errorf("###",zap.Error(err))
		return nil,-1,err
	}
	defer res.Body.Close()
	b,_:=io.ReadAll(res.Body)
	return b,res.StatusCode,nil
}

func NewSimpleHttpRequest(url, method string, body io.Reader,headers map[string]string) ([]byte,int, error) {
	httpReq:=new(HttpRequest)
	return httpReq.RetrieveResBody(url, method,body,headers)
}

func NewFullHttpRequest(urladdr, method,path string,query map[string]string,body io.Reader,headers map[string]string) ([]byte,int,error){
	httpReq:=new(HttpRequest)
	return httpReq.RetrieveResBodyByFullParams(urladdr, method,path,query,body,headers)
}