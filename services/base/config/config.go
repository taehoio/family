package config

import (
	"os"
	"strings"
)

const (
	productName        = "family"
	defaultEnvironment = "development"
)

var (
	environmentKeys = []string{"ENV", "APP_ENV", "NODE_ENV", "UP_STAGE"}
)

type IFace interface {
	ProductName() string
	ServiceName() string
	Env() string
	Prefix() string
}

type Config struct {
	IFace
	productName string
	serviceName string
	env         string
}

func New(serviceName string) (cfg *Config) {
	env := defaultEnvironment
	for _, key := range environmentKeys {
		value := os.Getenv(key)
		if value != "" {
			env = value
			break
		}
	}

	return &Config{
		productName: productName,
		serviceName: serviceName,
		env:         env,
	}
}

func NewMock(serviceName string) (cfg *Config) {
	return &Config{
		productName: productName,
		serviceName: serviceName,
		env:         defaultEnvironment,
	}
}

func (c *Config) ProductName() string {
	return c.productName
}

func (c *Config) ServiceName() string {
	return c.serviceName
}

func (c *Config) Env() string {
	return c.env
}

func (c *Config) Prefix() string {
	return strings.Join([]string{c.ProductName(), c.Env(), c.ServiceName()}, "-")
}
