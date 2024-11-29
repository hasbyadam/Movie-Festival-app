package entity

import "time"

type Config struct {
	API struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"api"`
	Context struct {
		Timeout time.Duration
	}
	Database struct {
		Pg Pg `mapstructure:"pg"`
	} `mapstructure:"database"`
}

type Pg struct {
	Host                  string        `mapstructure:"host"`
	Port                  string        `mapstructure:"port"`
	Dbname                string        `mapstructure:"dbname"`
	User                  string        `mapstructure:"user"`
	Password              string        `mapstructure:"password"`
	Sslmode               string        `mapstructure:"sslmode"`
	MaxOpenConnection     int           `mapstructure:"max_open_connection"`
	MaxIdleConnection     int           `mapstructure:"max_idle_connection"`
	MaxConnectionLifetime time.Duration `mapstructure:"max_connection_lifetime"`
}
