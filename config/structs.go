package config

type Config struct {
	App AppConfig `mapstructure:"app"`
	Log LogConfig `mapstructure:"log"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type LogConfig struct {
	// Possible variables are (info,debug,error, and warn or warning)
	Level      string `mapstructure:"level"`
	ShowSource bool   `mapstructure:"show_source"`
}
