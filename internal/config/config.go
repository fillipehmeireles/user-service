package config

import (
	"log"

	customerrors "github.com/fillipehmeireles/user-service/pkg/custom-errors"
	"github.com/spf13/viper"
)

const (
	PG_DSN = "db.postgres.dsn"
)

type (
	// Dababase configs
	PgConfig struct {
		Dsn string
	}
)
type Config struct {
	PgConfig PgConfig
}

func NewConfig(configPath string) (*Config, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("[Config:NewConfig] Could not read config file.")
		return nil, err
	}

	pgConfig, err := setupPg()
	if err != nil {
		return nil, err
	}
	return &Config{
		PgConfig: *pgConfig,
	}, nil
}

func setupPg() (*PgConfig, error) {
	if !viper.IsSet(PG_DSN) {
		err := &customerrors.ConfigKeyNotFound{Key: PG_DSN}
		log.Println(err)
		return nil, err
	}
	dsn := viper.GetString(PG_DSN)

	return &PgConfig{dsn}, nil
}
