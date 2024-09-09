package config

import (
	"github.com/pudthaiiii/golang-cms/src/types"
	"github.com/pudthaiiii/golang-cms/src/utils"
)

func GetPGConfig() types.PGConfig {
	return types.PGConfig{
		Host:       utils.RequireEnv("DB_HOST"),
		Port:       utils.RequireEnv("DB_PORT"),
		User:       utils.RequireEnv("DB_USERNAME"),
		Password:   utils.RequireEnv("DB_PASSWORD"),
		DBDatabase: utils.RequireEnv("DB_DATABASE"),
	}
}