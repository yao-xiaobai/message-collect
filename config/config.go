package config

type Config struct {
	Kafka struct {
		Host string `yaml:"host"`
	} `yaml:"kafka"`
	Port int `yaml:"port"`
}

type ConsumeConfig struct {
	Topic  string `yaml:"topic"`
	HOST   string `yaml:"host"`
	Group  string `yaml:"group"`
	Offset int64  `yaml:"offset"`
}
