package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBMS     string         `yaml:"dbms"`
	Driver   string         `yaml:"driver"`
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Logger   LoggerConfig   `yaml:"logger"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	Dbname   string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"sslmode"`
}

type LoggerConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

func Load() (*Config, error) {
	config := &Config{
		DBMS:     os.Getenv("DBMS"),
		Driver:   os.Getenv("DRIVER"),
		Server:   ServerConfig{Host: os.Getenv("HOST"), Port: os.Getenv("PORT")},
		Database: DatabaseConfig{Dbname: os.Getenv("DB_NAME"), User: os.Getenv("DB_USER"), Password: os.Getenv("DB_PASSWORD"), SSLMode: os.Getenv("DB_SSLMODE")},
	}

	return config, nil
}

func (c *Config) GetDSN() string {

	switch c.DBMS {
	case "psql":
		return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
			c.Server.Host, c.Server.Port, c.Database.Dbname, c.Database.User, c.Database.Password, c.Database.SSLMode)
	case "mssql":
		return fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;",
			c.Server.Host, c.Database.User, c.Database.Password, c.Database.Dbname)
	default:
		return "wrong DBMS name in config"
	}

}
