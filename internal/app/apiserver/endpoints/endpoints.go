package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VitalyCone/kuznecov_communities_api/internal/app/model/requestmodel"
	"github.com/VitalyCone/kuznecov_communities_api/internal/app/serviceurl"
	"github.com/VitalyCone/kuznecov_communities_api/internal/app/store"
	"github.com/gin-gonic/gin"
)

type Endpoints struct {
	store *store.Store
}

func NewEndpoints(s *store.Store) *Endpoints{
	return &Endpoints{
		store: s,
	}
}

func (ep *Endpoints) Ping(g *gin.Context){
	g.JSON(200, "PING")
}

func requestFilesInCloudStorage(fileIds []int) ([]requestmodel.FileResponse, error){
	respFiles := make([]requestmodel.FileResponse, 0)
	
	for _,fileId := range fileIds{
		resp, err := http.Get(fmt.Sprintf("%s/%d",serviceurl.Get().CloudStorage.FileUrl,fileId))
		if err!= nil{
			return nil, err
		}
		defer resp.Body.Close()

		file := requestmodel.FileResponse{}

		if err := json.NewDecoder(resp.Body).Decode(&file); err != nil {
			return nil, err
		}
		respFiles = append(respFiles, file)
	}
	return respFiles, nil
}