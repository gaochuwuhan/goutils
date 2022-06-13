package feishu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gaochuwuhan/goutils"
	"github.com/gaochuwuhan/goutils/pkg/cafe"
	"github.com/gaochuwuhan/goutils/pkg/tools"
)

const (
	feishu_access_token_url = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal/"
	feishu_auth_user_url    = "https://open.feishu.cn/open-apis/authen/v1/access_token"
	feishu_contact_userinfo = "https://open.feishu.cn/open-apis/contact/v3/users"
)




type TenantAccessTokenApi struct {
	URL string `json:"url"`
	Method string `json:"method"`
	Body *TenantAccessTokenBody `json:"body"`
	Headers TenantAccessTokenHeaders `json:"headers"`
	Response TenantAccessTokenResponse
}
type TenantAccessTokenBody struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type TenantAccessTokenHeaders map[string]string

type TenantAccessTokenResponse struct {
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int    `json:"expire"`
}

func(self *TenantAccessTokenApi) NewCMDBTenantAccessTokenApi() *TenantAccessTokenApi{
	app_id := goutils.VP.GetString(cafe.JoinStr(goutils.ENV,".thirdpart.feishu.app_id"))
	app_secret := goutils.VP.GetString(cafe.JoinStr(goutils.ENV,".thirdpart.feishu.app_secret"))
	self.URL = feishu_access_token_url
	self.Method = "POST"
	self.Body = &TenantAccessTokenBody{
		AppId:app_id,
		AppSecret:app_secret,
	}
	headers:=make(TenantAccessTokenHeaders)
	headers["Content-Type"] = "application/json; charset=utf-8"
	self.Headers = headers
	return self
}

func(self *TenantAccessTokenApi) RetrieveTenantAccessToken() (string, error){
	api:=self.NewCMDBTenantAccessTokenApi()
	bodyByte,err:=json.Marshal(api.Body)
	if err != nil {
		return "",err
	}
	bodyBuffer:=bytes.NewBuffer(bodyByte)
	resBody,_,err:=tools.NewSimpleHttpRequest(api.URL,api.Method,bodyBuffer,api.Headers)
	if err != nil {
		return "", err
	}
	json.Unmarshal(resBody,&self.Response)
	if self.Response.Code !=0 {
		return "", fmt.Errorf(self.Response.Msg)
	}
	return self.Response.TenantAccessToken,nil
}

//获取登录用户身份
type UserAuthApi struct {
	URL string `json:"url"`
	Method string `json:"method"`
	Body *UserAuthBody `json:"body"`
	Headers UserAuthHeaders `json:"headers"`
	Response UserAuthResponse
}
type UserAuthBody struct {
	GrantType     string `json:"grant_type"`
	Code string `json:"code"`
}

type UserAuthHeaders map[string]string

type UserAuthResponse struct {
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	Data FeishuUser 		 `json:"data"`

}

//UserAuthResponse返回的部分用户信息，json格式要和飞书一致
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/authen-v1/authen/access_token
type FeishuUser struct {
	Email        string `json:"email"`
	UserCn       string `json:"name"`
	UserId  	 string `json:"user_id"`
}

func(self *UserAuthApi) NewUserAuthApi(code, token string) *UserAuthApi{
	self.URL = feishu_auth_user_url
	self.Method = "POST"
	self.Body = &UserAuthBody{
		GrantType:"authorization_code",
		Code:code,
	}
	headers:=make(UserAuthHeaders)
	headers["Content-Type"] = "application/json; charset=utf-8"
	headers["Authorization"] = cafe.JoinStr("Bearer ",token)
	self.Headers = headers
	return self
}

func(self *UserAuthApi) RetrieveUserInfo(code, token string) (*FeishuUser,int, error){
	api:=self.NewUserAuthApi(code,token)
	bodyByte,err:=json.Marshal(api.Body)
	if err != nil {
		return nil,-1,err
	}
	bodyBuffer:=bytes.NewBuffer(bodyByte)
	resBody,status,err:=tools.NewSimpleHttpRequest(api.URL,api.Method,bodyBuffer,api.Headers)
	if err != nil {
		return nil, status,err
	}
	json.Unmarshal(resBody,&self.Response)
	if self.Response.Code !=0 {
		return nil, -1,fmt.Errorf(self.Response.Msg)
	}
	return &self.Response.Data,status,nil
}

//获取用户详细信息接口
//https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/get
type ContactUser struct {
	User struct {
		UnionId       string `json:"union_id"`
		UserId        string `json:"user_id"`
		OpenId        string `json:"open_id"`
		Name          string `json:"name"`
		EnName        string `json:"en_name"`
		Nickname      string `json:"nickname"`
		Email         string `json:"email"`
		Mobile        string `json:"mobile"`
		MobileVisible bool   `json:"mobile_visible"`
		Gender        int    `json:"gender"`
		Avatar        struct {
			Avatar72     string `json:"avatar_72"`
			Avatar240    string `json:"avatar_240"`
			Avatar640    string `json:"avatar_640"`
			AvatarOrigin string `json:"avatar_origin"`
		} `json:"avatar"`
		Status struct {
			IsFrozen    bool `json:"is_frozen"`
			IsResigned  bool `json:"is_resigned"`
			IsActivated bool `json:"is_activated"`
			IsExited    bool `json:"is_exited"`
			IsUnjoin    bool `json:"is_unjoin"`
		} `json:"status"`
		DepartmentIds   []string `json:"department_ids"`
		LeaderUserId    string   `json:"leader_user_id"`
		City            string   `json:"city"`
		Country         string   `json:"country"`
		WorkStation     string   `json:"work_station"`
		JoinTime        int      `json:"join_time"`
		IsTenantManager bool     `json:"is_tenant_manager"`
		EmployeeNo      string   `json:"employee_no"`
		EmployeeType    int      `json:"employee_type"`
		Orders          []struct {
			DepartmentId    string `json:"department_id"`
			UserOrder       int    `json:"user_order"`
			DepartmentOrder int    `json:"department_order"`
		} `json:"orders"`
		CustomAttrs []struct {
			Type  string `json:"type"`
			Id    string `json:"id"`
			Value struct {
				Text        string `json:"text"`
				Url         string `json:"url"`
				PcUrl       string `json:"pc_url"`
				OptionValue string `json:"option_value"`
				Name        string `json:"name"`
				PictureUrl  string `json:"picture_url"`
				GenericUser struct {
					Id   string `json:"id"`
					Type int    `json:"type"`
				} `json:"generic_user"`
			} `json:"value"`
		} `json:"custom_attrs"`
		EnterpriseEmail string `json:"enterprise_email"`
		JobTitle        string `json:"job_title"`
	} `json:"user"`
}

type ContactUserInfoApi struct {
	URL string `json:"url"`
	Method string `json:"method"`
	URLPath string
	Query UserInfoQuery `json:"query"`
	Headers UserInfoHeaders `json:"headers"`
	Response UserInfoResponse
}


type UserInfoQuery map[string]string

type UserInfoHeaders map[string]string



type UserInfoResponse struct{
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	Data ContactUser 		 `json:"data"`
}

func(self *ContactUserInfoApi) NewContactUserInfoApi(tenant_access_token,userId string) *ContactUserInfoApi{
	self.URL = feishu_contact_userinfo
	self.Method = "GET"
	self.URLPath = userId
	headers:=make(UserInfoHeaders)
	headers["Authorization"] = cafe.JoinStr("Bearer ",tenant_access_token)
	self.Headers = headers
	query:=make(UserInfoQuery)
	query["user_id_type"] = ""
	self.Query = query
	return self
}

func (self *ContactUserInfoApi) RetrieveContactUserInfo(tenant_access_token,userId string) (*ContactUser,int,error){
	api:=self.NewContactUserInfoApi(tenant_access_token,userId)
	api.Query["user_id_type"]="user_id" //多个可选值，此处重写
	resBody,statusCode,err:=tools.NewFullHttpRequest(api.URL,api.Method,api.URLPath,api.Query,nil,api.Headers)
	if err != nil {
		return nil, statusCode,err
	}
	json.Unmarshal(resBody,&self.Response)
	if self.Response.Code !=0 {
		return nil, -1,fmt.Errorf(self.Response.Msg)
	}
	return &self.Response.Data,statusCode,nil
}
