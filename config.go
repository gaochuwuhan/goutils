package goutils

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	ENV string
	VP  *viper.Viper
	DB  *gorm.DB
)
