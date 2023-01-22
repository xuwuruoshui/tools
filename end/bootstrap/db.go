package bootstrap

import (
	"end/domain"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type DB struct {
	Name     string `mapstructure:"name" yaml:"name"`
	Host     string `mapstructure:"host" yaml:"host"`
	Port     int32  `mapstructure:"port" yaml:"port"`
	Category string `mapstructure:"category" yaml:"category"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
}

type DBClient interface {
	ConnectDB(*Config)
}

type MySqlClient struct {
	*gorm.DB
}

func (m *MySqlClient) ConnectDB(cfg *Config) {
	config := logger.Config{
		SlowThreshold:             time.Second, // 慢 SQL 阈值
		LogLevel:                  logger.Info, // 日志级别
		IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  true,        // 禁用彩色打印
	}
	lg := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		config,
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: lg, NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		panic("Mysql connect faild:" + err.Error())
	}

	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		panic(err)
	}
	m.DB = db
}
