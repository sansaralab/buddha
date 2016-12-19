package repository

import "github.com/sansaralab/buddha/models"

func (r *Repository) AddProject(name, gitUrl string) *models.Project {
	return &models.Project{Id: 123, Name: name, GitUrl: gitUrl}
}
