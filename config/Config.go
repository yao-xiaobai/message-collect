package config

type Config struct {
	Kafka struct {
		Host string `yaml:"host"`
	} `yaml:"kafka"`
	Port int `yaml:"port"`
}
