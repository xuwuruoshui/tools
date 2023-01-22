package bootstrap

type Cache struct {
	Name     string `mapstructure:"name" yaml:"name"`
	Host     string `mapstructure:"host" yaml:"host"`
	Port     int32  `mapstructure:"port" yaml:"port"`
	Category string `mapstructure:"category" yaml:"category"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
}

type CacheClient interface {
	ConnectCache(*Config)
}

type RedisClient struct {
}

func (r *RedisClient) ConnectCache(cfg *Config) {

}
