package cafe

import (
	"strings"
)

func JoinStr(strs ...string) string{
	//拼接字符串的函数
	var build strings.Builder
	for _,v :=range strs {
		build.WriteString(v)
	}
	ss:=build.String()
	return ss
}
