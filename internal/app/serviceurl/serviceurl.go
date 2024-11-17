package serviceurl

import "github.com/VitalyCone/kuznecov_communities_api/internal/app/model/requestmodel"

var serviceURL *ServiceURL

type ServiceURL struct {
	CloudStorage *requestmodel.CloudStorage	`yaml:"cloud_storage"`
}


// func NewServiceURL() *ServiceURL{

// }

func (s *ServiceURL) Init(){
	serviceURL = s
}

func Get() *ServiceURL{
	return serviceURL
}