package models

import (
	"database/sql/driver"
	"fmt"
	"github.com/gofrs/uuid"
	"time"
)



//MyTime 自定义时间
type MyTime struct {
	time.Time
}


// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t MyTime) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero(){
		return []byte("null"), nil
	}
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}


// Value insert timestamp into mysql need this function.
func (t MyTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *MyTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = MyTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}


type BaseModel struct{
	ID        uint `gorm:"primarykey"`
	CreatedAt MyTime `gorm:"<-:create;comment:'创建时间';type:datetime;" json:"CreatedAt,omitempty"` //创建时间不可更改
	UpdatedAt MyTime  `gorm:"comment:'修改时间';type:datetime;" json:"UpdatedAt,omitempty"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
}

type UUIDBase struct {
	ID 		  uuid.UUID `gorm:"type:char(36);primaryKey"`
	CreatedAt MyTime `gorm:"<-:create;comment:'创建时间';type:datetime;" json:"CreatedAt,omitempty"`
	UpdatedAt MyTime  `gorm:"comment:'修改时间';type:datetime;" json:"UpdatedAt,omitempty"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
}


