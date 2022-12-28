package store

type Config struct {
	DataBaseURL    string `toml:"database_url"`
	DataBaseDriver string `toml:"db_driver"`
}

func NewConfig() *Config {
	return &Config{
		DataBaseDriver: "postgres",
	}
}
