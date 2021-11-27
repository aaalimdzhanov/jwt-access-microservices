package apiserver



type Config struct {
	BindAdr  string `toml:"bind_addr"`
	Loglevel string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAdr:  ":9090",
		Loglevel: "debug",
	}
}
