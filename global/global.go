package global

import "dwbackend/ent"

type GlobalConfig struct {
	Http struct {
		Port int
	}
	MySQL struct {
		Host     string
		Port     int
		User     string
		Password string
		DB       string
	}
}

var GlobalConfigInstance *GlobalConfig

var DB *ent.Client
