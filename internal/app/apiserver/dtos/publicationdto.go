package dtos

import (
	"github.com/VitalyCone/kuznecov_communities_api/internal/app/model"
	"github.com/VitalyCone/kuznecov_communities_api/internal/app/model/requestmodel"
)

type CreatePublicationDto struct {
	Text string `form:"text" json:"text"`
	//Likes     int
}

type CreatePublicationDetailsDto struct {
	Data  *model.Publication           `form:"data" json:"data"`
	Files []requestmodel.FileResponse `form:"files" json:"files"`
}
