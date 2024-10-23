package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/VitalyCone/kuznecov_communities_api/internal/app/apiserver"
	"github.com/VitalyCone/kuznecov_communities_api/internal/app/store"
)

var (
	configPath string
)

func init() {
	configPath = "config/apiserver.toml"
}

func main() {
	configServer := apiserver.NewConfig()
	configStore := store.NewConfig()

	if _, err := toml.DecodeFile(configPath, configServer); err != nil {
		log.Fatal(err)
	}

	if _, err := toml.DecodeFile(configPath, configStore); err != nil {
		log.Fatal(err)
	}

	store := store.NewStore(configStore)
	server := apiserver.NewAPIServer(configServer, store)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
