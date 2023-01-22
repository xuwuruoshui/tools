package bootstrap

import "time"

type Config struct {
	Name    string        `mapstructure:"name" yaml:"name"`
	Version string        `mapstructure:"version" yaml:"version"`
	Timeout time.Duration `mapstructure:"timeout" yaml:"timeout"` // 一次请求的超时时间
	DB      *DB           `mapstructure:"db" yaml:"db"`
	Cache   *Cache        `mapstructure:"cache" yaml:"cache"`
}
