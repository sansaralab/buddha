package models

type Project struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	GitUrl string `json:"git_url"`
}

func NewProject() *Project {
	return &Project{}
}
