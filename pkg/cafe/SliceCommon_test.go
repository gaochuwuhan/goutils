package cafe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceIntersect(t *testing.T) {
	s1:=[]string{"1", "3","2","2",  "6", "8"}
	s2:=[]string{"8","2", "3", "5","4", "0"}
	result:=SliceIntersect(s1,s2)
	expected:=[]string{"8","2","3"}
	assert.Equal(t,expected,result)
}

func TestSliceDiff1(t *testing.T) {
	s1:=[]string{"1", "3","2","2",  "6", "8"}
	s2:=[]string{"8","2", "3", "5","4", "0"}
	result:=SliceDiff(s1,s2)
	expected:=[]string{"1","6"}
	assert.Equal(t,expected,result)
}

func TestSliceDiff2(t *testing.T) {
	s1:=[]string{}
	s2:=[]string{"1", "3"}
	result:=SliceDiff(s1,s2)
	expected:=[]string{}
	assert.Equal(t,expected,result)
}

func TestNegativeSliceDiff(t *testing.T) {
	s1:=[]string{"1", "3","2","2","6", "8"}
	s2:=[]string{"8","2", "3", "5","4", "0"}
	result:=SliceDiff(s1,s2)
	expected:=[]string{"1","6","5","4","0"}
	assert.NotEqual(t,expected,result)
}

func TestSliceRemoveDup(t *testing.T) {
	s1:=[]string{
		"1","1","2","3",
	}
	result:=SliceRemoveDup(s1)
	expected:=[]string{"1","2","3",}
	assert.Equal(t,expected,result)
}

