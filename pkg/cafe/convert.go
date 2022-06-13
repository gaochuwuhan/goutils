package cafe

import (
	"encoding/json"
	"errors"
	"github.com/gofrs/uuid"
	"time"
)

func UTCStrToTime(UTCTimestr string) (tm time.Time,err error){
	t, err := time.Parse(time.RFC3339, UTCTimestr)
	if err != nil {
		return time.Time{},err
	}
	return t,nil
}


func InterfaceToStr(in interface{}) (string,error){
	v,ok:=in.(string)
	if ok{
		return v,nil
	}else{
		err:=errors.New("type error!")
		return "",err
	}

}

func StructToMapSmp(in interface{}) (map[string]interface{},error){
	/*
	in: 结构体对象的指针
	 */
	m:=make(map[string]interface{})
	bt,_:=json.Marshal(in)
	err:=json.Unmarshal(bt,&m)
	if err!=nil{
		return map[string]interface{}{} , err
	}
	return m, nil
}


func BytesToMap(bt []byte) (map[string]interface{},error){
	m := make(map[string]interface{})
	err := json.Unmarshal(bt, &m)
	if err!=nil{
		return map[string]interface{}{},err
	}else{
		return m,nil
	}
}

func MapToByte(m map[string]interface{}) ([]byte,error){
	bt,err:=json.Marshal(m)
	if err!=nil{
		return []byte{},err
	}else{
		return bt,nil
	}
}


func UuidstrToUUid(s string) uuid.UUID {
	uuid:=uuid.Must(uuid.FromString(s))
	return uuid
}