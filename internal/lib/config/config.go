package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Api struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"api"`

	Sql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		SslMode  string `yaml:"sslmode"`
		Database string `yaml:"database"`
	} `yaml:"sql"`

	ErrorsFile string `yaml:"errors"`
}

func NewConfig(path string) (*Config, error) {
	config := &Config{}
	return config.LoadConfig(path)
}

func (c *Config) LoadConfig(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Config) GetDBConnection() string {
	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Europe/Moscow", c.Sql.Host, c.Sql.User, c.Sql.Password, c.Sql.Database, c.Sql.Port, c.Sql.SslMode)
	return conn
}
