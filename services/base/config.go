package base

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

type Config interface {
	ProductName() string
	ServiceName() string
	Env() string
	Prefix() string
}

type defaultConfig struct {
	Config

	productName string
	serviceName string
	env         string
}

func NewConfig(serviceName string) (cfg *defaultConfig) {
	env := defaultEnvironment
	for _, key := range environmentKeys {
		value := os.Getenv(key)
		if value != "" {
			env = value
			break
		}
	}

	return &defaultConfig{
		productName: productName,
		serviceName: serviceName,
		env:         env,
	}
}

func NewMockConfig(serviceName string) (cfg *defaultConfig) {
	return &defaultConfig{
		productName: productName,
		serviceName: serviceName,
		env:         defaultEnvironment,
	}
}

func (c *defaultConfig) ProductName() string {
	return c.productName
}

func (c *defaultConfig) ServiceName() string {
	return c.serviceName
}

func (c *defaultConfig) Env() string {
	return c.env
}

func (c *defaultConfig) Prefix() string {
	return strings.Join([]string{c.ProductName(), c.Env(), c.ServiceName()}, "-")
}
