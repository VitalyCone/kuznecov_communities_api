package endpoints

import (
	"net/http"
	"strconv"

	"github.com/VitalyCone/kuznecov_communities_api/internal/app/apiserver/dtos"
	"github.com/gin-gonic/gin"
)

// @Summary Get News
// @Schemes
// @Description Get News
// @Tags Publication
// @Accept json
// @Produce json
// @Param offset query int false "offset from first responses"
// @Param limit query int false "restriction on return of publications"
// @Router /news [GET]
func (ep *Endpoints) GetNews(g *gin.Context) {
	offset := 0
	limit := 999
	publicationsDetails := make([]dtos.CreatePublicationDetailsDto, 0)
	queryOffset,isOffset := g.GetQuery("offset")
	if isOffset{
		if queryOffset != ""{
			num, err := strconv.Atoi(queryOffset)
			if err != nil {
				g.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect offset query " + error.Error(err)})
				return
			}

			offset = num
		}
	}
	queryLimit,isLimit := g.GetQuery("limit")
	if isLimit{
		if queryLimit != ""{
			num, err := strconv.Atoi(queryLimit)
			if err != nil {
				g.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect limit query " + error.Error(err)})
				return
			}
	
			limit = num
		}
	}

	publications, err := ep.store.Publication().GetAll_SortByCreatedTime(limit, offset)
	if err != nil{
		g.JSON(http.StatusNotFound, gin.H{"message": "Failed to get publications: " + err.Error()})
		return
	}
	for _,publication := range publications{
		respFiles, err := requestFilesInCloudStorage(publication.FileIds)
		if err != nil{
			g.JSON(http.StatusNotFound, gin.H{"message": "Failed send request to cloud storage: " + err.Error()})
			return
		}
		
		publicationDetails := dtos.CreatePublicationDetailsDto{
			Data: &publication,
			Files: respFiles,
		}

		publicationsDetails = append(publicationsDetails, publicationDetails)
	}

	g.JSON(http.StatusCreated, publicationsDetails)

	// for _,fileId := range publication.FileIds{
	// 	resp, err := http.Get(fmt.Sprintf("%s/%d",serviceurl.Get().CloudStorage.FileUrl,fileId))
	// 	if err!= nil{
	// 		g.JSON(http.StatusInternalServerError, gin.H{"message": "Failed send request to other api: " + err.Error()})
	// 		return
	// 	}
	// 	defer resp.Body.Close()

	// 	file := requestmodel.FileResponse{}

	// 	if err := json.NewDecoder(resp.Body).Decode(&file); err != nil {
	// 		g.JSON(http.StatusInternalServerError, gin.H{"message": "Files decode files from other api: " + err.Error()})
	// 		return
	// 	}
	// 	respFiles = append(respFiles, file)
	// }
	// publicationDetails := dtos.CreatePublicationDetailsDto{
	// 	Data: publication,
	// 	Files: respFiles,
	// }

	// g.JSON(http.StatusOK, publicationDetails)
}