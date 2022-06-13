package serve

import (
	"github.com/gaochuwuhan/goutils"
	"github.com/gaochuwuhan/goutils/logger"
	"github.com/gaochuwuhan/goutils/pkg/cafe"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	dsn := goutils.VP.GetString(cafe.JoinStr(goutils.ENV,".mysqlconn")) //连接数据库db
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		SkipDefaultTransaction: true, //跳过全局事务
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "cld_",//表名前缀，`User`结构体的表名应该是`cld_users`
			//SingularTable: true
		},
		Logger:gormlog.Default.LogMode(gormlog.Error), //日志等级
		DisableForeignKeyConstraintWhenMigrating: true,//逻辑外键（代码里面自动体现外键关系）加快数据库处理速度
	})

	if err != nil {
		logger.Log.Error("MySQL启动异常", zap.Error(err))
		panic(err)
		//return nil
	} else {
		//logger.Log.Info("Mysql连接建立！")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		return db
	}
}
