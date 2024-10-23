package apiserver

import (
	"log"

	//"github.com/VitalyCone/kuznecov_communities_api/internal/app/apiserver/endpoints"
	"github.com/VitalyCone/kuznecov_communities_api/internal/app/store"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

var (
	mainPath string = "/communities"
)

type APIServer struct {
	config *Config
	router *gin.Engine
	store *store.Store
}

func NewAPIServer(config *Config, store *store.Store) *APIServer{
	return &APIServer{
		config: config,
		router: gin.Default(),
		store: store,
	}
}

func (s *APIServer) Start() error{
	log.Println("starting api server on ")

	s.configureEndpoints()
	
	return s.router.Run(s.config.BindAddr)
}

func (s *APIServer) configureEndpoints() {
	//endpoint := endpoints.NewEndpoints(s.store)

	docs.SwaggerInfo.BasePath = mainPath
	path := s.router.Group(mainPath)
	{
		path.GET("/groups") //получение всех групп для юзера, с query по поиску среди всех групп с приоритетом добавленных
		path.GET("/groups/:id") //получить инфу об определенной группе
		path.POST("/groups") //создание группы
		path.DELETE("/groups/:id") //удаление группы по id
		path.PUT("/groups/:id") //изменение группы по id

		path.GET("/news") //получение новостей определенного пользователя

		path.GET("/post/:id") //получение поста по его id
		path.POST("/post") //создание поста
		path.DELETE("/post/:id") //удаление поста
		path.PUT("/post/:id") //изменение поста по id
	}


	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}