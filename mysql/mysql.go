// @Description  mysql
// @Author  	 jiangyang  
// @Created  	 2020/10/30 3:44 下午

// Example Config:
// mysql:
//   user: root
//   password: 123456
//   host: 127.0.0.1
//   port: 3306
//   dbname: demo
//   max_idle_conn: 10
//   max_open_conn: 100
//   debug: true

package mysql

import (
	"fmt"
	"gorm.io/gorm/logger"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// mysql连接
	db *gorm.DB
	// 保证只建立一次连接
	once sync.Once
)

// Mysql配置结构体
type Config struct {
	User        string `json:"user" yaml:"user"`                                                 // 用户名
	Password    string `json:"password" yaml:"password"`                                         // 密码
	Host        string `json:"host" yaml:"host"`                                                 // 主机地址
	Port        int    `json:"port" yaml:"port"`                                                 // 端口号
	Dbname      string `json:"dbname" yaml:"dbname"`                                             // 数据库名
	MaxIdleConn int    `json:"max_idle_conn" yaml:"max_idle_conn" mapstructure:"max_idle_conn"`  // 最大空闲连接
	MaxOpenConn int    `json:"max_open_conn" yaml:"max_open_conn" mapstructure:"max_open_conn" ` // 最大活跃连接
	Debug       bool   `json:"debug" yaml:"debug"`                                               // 是否开启Debug（开启Debug会打印数据库操作日志）
}

// 初始化数据库
func Init(mysqlConfig Config) {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
			mysqlConfig.User,
			mysqlConfig.Password,
			mysqlConfig.Host,
			mysqlConfig.Port,
			mysqlConfig.Dbname,
		)

		conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.New(
				logrus.StandardLogger(),
				logger.Config{
					SlowThreshold: time.Second,
				},
			),
		})

		if err != nil {
			logrus.Fatalf("mysql connect failed: %v", err)
		}

		sqlDB, err := conn.DB()
		if err != nil {
			logrus.Fatalf("mysql connPool failed: %v", err)
		}
		sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConn)
		sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConn)
		sqlDB.SetConnMaxLifetime(time.Hour)

		db = conn
		logrus.Info("mysql connect successfully")
	})
}

// 获取Mysql连接
func Conn() *gorm.DB {
	return db
}

// Close method
func Close() error {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return errors.WithStack(err)
		}
		if err := sqlDB.Close(); err != nil {
			return errors.WithStack(err)
		}
	}

	logrus.Info("mysql connect closed")
	return nil
}
