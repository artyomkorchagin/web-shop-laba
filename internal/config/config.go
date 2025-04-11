package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
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
	Port int    `yaml:"port"`
}

type DatabaseConfig struct {
	Dbname          string        `yaml:"dbname"`
	Name            string        `yaml:"name"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	SSLMode         string        `yaml:"sslmode"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

type LoggerConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

func Load(configPath string) (*Config, error) {
	config := &Config{}
	dir, err := filepath.Abs(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %v", err)
	}
	file, err := os.ReadFile(dir)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	return config, nil
}

func (c *Config) GetDSN() string {

	switch c.DBMS {
	case "pgsql":
		return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
			c.Server.Host, c.Server.Port, c.Database.Dbname, c.Database.User, c.Database.Password, c.Database.SSLMode)
	case "mssql":
		return fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;",
			c.Server.Host, c.Database.User, c.Database.Password, c.Database.Dbname)
	default:
		return "wrong DBMS name in config"
	}

}
