package requestmodel

type CloudStorage struct {
	FileUrl string `yaml:"file_ep"`
	DockerHost string `yaml:"docker_host"`
	LocalHost string `yaml:"local_host"`
}

type FileResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Format    string `json:"format"`
	Body      string `json:"body"`
	CreatedAt string `json:"created"`
	UpdatedAt string `json:"updated"`
	Size      int64  `json:"size"`
}
