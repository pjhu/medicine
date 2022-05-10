package datasource

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	once              sync.Once
	datetimePrecision = 2
)

var db *gorm.DB

func Builder() *gorm.DB {
	once.Do(func() {
		initMysql()
	})
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}

func GetDB() *gorm.DB {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      false,
		})
	session := db.Session(&gorm.Session{Logger: newLogger})
	return session
}

// Init DBConnect for db connection
func initMysql() {

	gormConfig := gorm.Config{
		PrepareStmt:                              false,
		Logger:                                   logger.Default,
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
		DisableForeignKeyConstraintWhenMigrating: true,
		CreateBatchSize:                          1000,
		DisableNestedTransaction:                 true,
		SkipDefaultTransaction:                   true,
	}

	var err error
	dns := viper.GetString("datasource.master.jdbcUrl")
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,                // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,                // string 类型字段的默认长度
		DisableDatetimePrecision:  true,               // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DefaultDatetimePrecision:  &datetimePrecision, // default datetime precision
		DontSupportRenameIndex:    true,               // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,               // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,              // 根据当前 MySQL 版本自动配置
	}), &gormConfig)

	if err != nil {
		logrus.Info("MySQL Connected Failure: %v", err)
	}
	logrus.Info("MySQL Connected: %s", dns)
}
