package crypt

type Config interface {
}

type DefaultCryptConfig struct {
	Config
}

func NewConfig() Config {
	return &DefaultCryptConfig{}
}

func NewMockConfig() Config {
	return NewConfig()
}
