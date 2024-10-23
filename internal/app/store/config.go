package store

type Config struct{
	databaseURL string `toml:"database_url"`
}

func NewConfig() *Config{
	return &Config{
		databaseURL: "host=localhost port=5322 user=admin dbname=kuznecov_communities password=admin sslmode=disable",
	}
}