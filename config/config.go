package config

import (
	"github.com/BurntSushi/toml"
)

// Config 对应配置文件结构
type Config struct {
	Active          string   `toml:"active"`
	MongoDbHost     []string `toml:"MongoDbHost"`
	MongoDbName     string   `toml:"MongoDbName"`
	MongoDbUser     string   `toml:"MongoDbUser"`
	MongoDbPassword string   `toml:"MongoDbPassword"`

	MailHost     string   `toml:"MailHost"`
	MailPort     int      `toml:"MailPort"`
	MailFrom     string   `toml:"MailFrom"`
	MailPassword string   `toml:"MailPassword"`
	MailTo       []string `toml:"MailTo"`
}

// 解析toml配置
func UnmarshalConfig(tomlfile string) (*Config, error) {
	c := &Config{}
	if _, err := toml.DecodeFile(tomlfile, c); err != nil {
		return c, err
	}
	return c, nil
}

func (c Config) GetActive() string {
	return c.Active
}

func (c Config) GetMongoDbHost() []string {
	return c.MongoDbHost
}

func (c Config) GetMongoDbName() string {
	return c.MongoDbName
}

func (c Config) GetMongoDbUser() string {
	return c.MongoDbUser
}

func (c Config) GetMongoDbPassword() string {
	return c.MongoDbPassword
}

func (c Config) GetMailHost() string {
	return c.MailHost
}

func (c Config) GetMailPort() int {
	return c.MailPort
}

func (c Config) GetMailFrom() string {
	return c.MailFrom
}

func (c Config) GetMailPassword() string {
	return c.MailPassword
}

func (c Config) GetMailTo() []string {
	return c.MailTo
}
