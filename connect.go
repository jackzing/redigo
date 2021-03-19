package redis

import (
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

var DB *redigo.Pool

type Config struct {
	Host        string
	Password    string
	MaxIdle     int
	IdleTimeout time.Duration
	MaxActive   int
}

func ConRedis(conf Config) {
	DB = &redigo.Pool{
		MaxIdle:     conf.MaxIdle, //空闲数
		IdleTimeout: conf.IdleTimeout * time.Second,
		MaxActive:   conf.MaxActive, //最大数
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", conf.Host)
			if err != nil {
				return nil, err
			}
			if conf.Password != "" {
				if _, err := c.Do("AUTH", conf.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
