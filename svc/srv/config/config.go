package config

type IFace interface {
	Namespace() string
}

type Config struct {
	IFace

	namespace string
}

func New(namespace string) (cfg *Config) {
	return &Config{
		namespace: namespace,
	}
}

func (c *Config) Namespace() string {
	return c.namespace
}
