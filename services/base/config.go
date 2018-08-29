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

type DefaultConfig struct {
	Config

	productName string
	serviceName string
	env         string
}

func NewConfig(serviceName string) (cfg *DefaultConfig) {
	env := defaultEnvironment
	for _, key := range environmentKeys {
		value := os.Getenv(key)
		if value != "" {
			env = value
			break
		}
	}

	return &DefaultConfig{
		productName: productName,
		serviceName: serviceName,
		env:         env,
	}
}

func NewMockConfig(serviceName string) (cfg *DefaultConfig) {
	return &DefaultConfig{
		productName: productName,
		serviceName: serviceName,
		env:         defaultEnvironment,
	}
}

func (c *DefaultConfig) ProductName() string {
	return c.productName
}

func (c *DefaultConfig) ServiceName() string {
	return c.serviceName
}

func (c *DefaultConfig) Env() string {
	return c.env
}

func (c *DefaultConfig) Prefix() string {
	return strings.Join([]string{c.ProductName(), c.Env(), c.ServiceName()}, "-")
}
