package main

import (
	"log"
	"os"
	"strings"

	"github.com/VitalyCone/kuznecov_communities_api/internal/app/apiserver"
	"github.com/VitalyCone/kuznecov_communities_api/internal/app/serviceurl"
	"github.com/VitalyCone/kuznecov_communities_api/internal/app/store"
	"gopkg.in/yaml.v3"
)

var (
	configPath  string
	dockerCheck string
	emptyHost string
)

func init() {
	configPath = "config/apiserver.yaml"
	dockerCheck = "DOCKER_ENV"
	emptyHost = "EMPTY"
}

type configData struct {
	ApiAddr     string                 `yaml:"api_addr"`
	DbUrl       string                 `yaml:"database_url"`
	DbDockerUrl string                 `yaml:"database_docker_url"`
	ServiceURL  *serviceurl.ServiceURL `yaml:"service_url"`
}

func main() {
	var configStore *store.Config
	var configServer *apiserver.Config

	cfg := configData{}

	isDocker := os.Getenv(dockerCheck) == "true"

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	if isDocker {
		cfg.ServiceURL.CloudStorage.FileUrl = strings.Replace(cfg.ServiceURL.CloudStorage.FileUrl, emptyHost, cfg.ServiceURL.CloudStorage.DockerHost, 1)
		log.Println("App running in Docker. Using Docker database url")
		configStore = store.NewConfig(cfg.DbDockerUrl)
	} else {
		cfg.ServiceURL.CloudStorage.FileUrl = strings.Replace(cfg.ServiceURL.CloudStorage.FileUrl, emptyHost, cfg.ServiceURL.CloudStorage.LocalHost, 1)
		log.Println("App running without Docker. Using Local database url")
		configStore = store.NewConfig(cfg.DbUrl)
	}
	cfg.ServiceURL.Init()

	configServer = apiserver.NewConfig(cfg.ApiAddr)

	store := store.NewStore(configStore)
	server := apiserver.NewAPIServer(configServer, store)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
