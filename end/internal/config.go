package internal

import "time"

type Config struct {
	Name    string        `mapstructure:"db" yaml:"db"`
	Version string        `mapstructure:"version" yaml:"version"`
	Timeout time.Duration `mapstructure:"timeout" yaml:"timeout"`
	DB      *DB           `mapstructure:"db" yaml:"db"`
	Cache   *Cache        `mapstructure:"cache" yaml:"cache"`
}
