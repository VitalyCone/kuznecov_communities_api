package endpoints

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/VitalyCone/kuznecov_communities_api/internal/app/apiserver/dtos"
	"github.com/VitalyCone/kuznecov_communities_api/internal/app/model"
	"github.com/VitalyCone/kuznecov_communities_api/internal/app/model/requestmodel"
	"github.com/VitalyCone/kuznecov_communities_api/internal/app/serviceurl"
	"github.com/gin-gonic/gin"
)

// @Summary Get publication
// @Schemes
// @Description Get publication
// @Tags Publication
// @Accept json
// @Produce json
// @Param id path int true "publication id"
// @Router /publication/{id} [GET]
func (ep *Endpoints) GetPublication(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": error.Error(err)})
		return
	}

	publication, err := ep.store.Publication().GetById(id)
	if err != nil{
		g.JSON(http.StatusNotFound, gin.H{"message": "Failed to get publication: " + err.Error()})
		return
	}

	//g.JSON(http.StatusNotFound, publication.FileIds)
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
	respFiles, err := requestFilesInCloudStorage(publication.FileIds)
	if err != nil{
		g.JSON(http.StatusInternalServerError, gin.H{"message": "Failed send request to cloud storage: " + err.Error()})
		return
	}

	publicationDetails := dtos.CreatePublicationDetailsDto{
		Data: publication,
		Files: respFiles,
	}

	g.JSON(http.StatusOK, publicationDetails)
}

// @Summary Post publication
// @Schemes
// @Description Post publication
// @Tags Publication
// @Accept mpfd
// @Produce mpfd
// @Param files formData []file false "files"
// @Param data formData dtos.CreatePublicationDto true "<p>Publication data<p><p><u>created_at</u> <b>is not required</b><p>"
// @Router /publication [POST]
func (ep *Endpoints) PostPublication(g *gin.Context) {
	var createPublicationDto dtos.CreatePublicationDto
	
	respFiles:=  make([]requestmodel.FileResponse, 0)
	fileIds := make([]int,0)
	//var publication *model.Publication

	err := g.ShouldBind(&createPublicationDto)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind form (maybe data == nil): " + err.Error()})
		return
	}

	form, err := g.MultipartForm()
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form: " + err.Error()})
		return
	}
	files := form.File["files"]

	if len(files) > 0 {
		var b bytes.Buffer
		writer := multipart.NewWriter(&b)

		for _, file := range files {
			part, err := writer.CreateFormFile("files", file.Filename)
			if err != nil {
				g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create form file in cycle: " + err.Error()})
				return
			}

			src, err := file.Open()
			if err != nil {
				g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file in cycle: " + err.Error()})
				return
			}
			defer src.Close()

			_, err = io.Copy(part, src)

			if err != nil {
				g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy file in cycle: " + err.Error()})
				return
			}
		}
		writer.Close()
		resp, err := http.Post(serviceurl.Get().CloudStorage.FileUrl, writer.FormDataContentType(), &b)

		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed request form other api: " + err.Error()})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			g.JSON(resp.StatusCode, gin.H{"error": "Failed to upload files on other api"})
			return
		}

		if err := json.NewDecoder(resp.Body).Decode(&respFiles); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"message": "Files decode files from other api: " + err.Error()})
			return
		}

		for _, file := range respFiles{
			fileIds = append(fileIds, file.ID)
		}
	}


	publication := model.Publication{
		Text: createPublicationDto.Text,
		FileIds: fileIds,
	}

	err = ep.store.Publication().Create(&publication)
	if err != nil{
		g.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create publication: " + err.Error()})
		return
	}

	publicationDetails := dtos.CreatePublicationDetailsDto{
		Data: &publication,
		Files: respFiles,
	}

	g.JSON(http.StatusCreated, publicationDetails)
}

// @Summary Delete publication
// @Schemes
// @Description Delete publication
// @Tags Publication
// @Accept mpfd
// @Produce mpfd
// @Param id path int true "id"
// @Router /publication/{id} [DELETE]
func (ep *Endpoints) DeletePublication(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": error.Error(err)})
		return
	}

	ep.store.Publication().Delete(id)

	g.JSON(http.StatusNoContent, nil)
}
