package config

import (
	"os"
	"user-service/common/util"

	"github.com/sirupsen/logrus"
)

var config AppConfig

type AppConfig struct {
	Port                  int    `json:"port"`
	AppName               string `json:"appName"`
	AppEnv                string `json:"appEnv"`
	SignatureKey          string `json:"signatureKey"`
	Database              Database
	RateLimiterMaxRequest float64 `json:"rateLimiterMaxrequest"`
	RateLimiterTimeSecond int     `json:"rateLimiterTimeSecond"`
	JwtSecretKey          string  `json:"jwtSecretKey"`
	JwtExpirationTime     int     `json:"jwtExpirationTime"`
}

type Database struct {
	Host                  string `json:"host"`
	Port                  int    `json:"port"`
	Name                  string `json:"name"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	MaxOpenConnection     int    `json:"maxOpenConnection"`
	MaxLifeTimeConnection int    `json:"maxLifeTimeConnection"`
	MaxIdleConnection     int    `json:"maxIdleConnection"`
	MaxIdleTime           int    `json:"maxIdleTime"`
}

func Init() {
	err := util.BinFromJson(&config, "config.json", ".")
	if err != nil {
		logrus.Infof("failed to bind config %v", err)
		err = util.BinFromConsul(&config, os.Getenv("CONSUL_HTTP_URL"), os.Getenv("CONSUL_HTTP_KEY"))
		if err != nil {
			panic(err)
		}
	}
}
