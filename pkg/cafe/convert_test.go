package cafe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestUTCStrToTime(t *testing.T) {
 input := "2006-01-02T15:04:05Z"
 expected,_ := UTCStrToTime(input)
 expectedType:=reflect.TypeOf(expected).String()
 assert.Equal(t,expectedType,"time.Time")
}

//func TestNegativeUTCStrToTime(t *testing.T) {
//	input := []string{"haha","haha","hha","hh"}
//	expected := []string{"haha","hha"}
//	assert.NotEqual(t,expected,UTCStrToTime(input))
//}
