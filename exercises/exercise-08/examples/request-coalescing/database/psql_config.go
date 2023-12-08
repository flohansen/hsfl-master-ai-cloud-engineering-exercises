package database

import "fmt"

type PsqlConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func (cfg PsqlConfig) Dsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database)
}
