package bootstrap

import (
	"github.com/spf13/viper"
)

type App struct {
	Config      *Config     `json:"config"`
	DBClient    DBClient    `json:"DBClient"`
	CacheClient CacheClient `json:"cacheClient"`
}

func NewApp() *App {
	// 1.获取配置
	var config Config
	v := viper.New()
	v.SetConfigFile("./application.yml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	// 2.设置各种客户端
	ap := &App{
		Config:      &config,
		DBClient:    &MySqlClient{},
		CacheClient: &RedisClient{},
	}

	// 连接数据库
	ap.DBClient.ConnectDB(&config)
	// 连接redis
	ap.CacheClient.ConnectCache(&config)

	return ap
}
