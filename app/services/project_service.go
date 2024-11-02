package services

import (
	"msvc-syahril-app/app/interfaces"
	"msvc-syahril-app/app/models"
)

type ProjectService struct {
    repository interfaces.ProjectRepository
}

func NewProjectService(repo interfaces.ProjectRepository) *ProjectService {
    return &ProjectService{
        repository: repo,
    }
}

func (s *ProjectService) GetAllProjects() ([]models.Project, error) {
    return s.repository.GetAll()
}

func (s *ProjectService) GetProjectByID(id int) (*models.Project, error) {
    return s.repository.GetByID(id)
}

func (s *ProjectService) CreateProject(project *models.Project) error {
    return s.repository.Create(project)
}

func (s *ProjectService) UpdateProject(project *models.Project) error {
    return s.repository.Update(project)
}

func (s *ProjectService) DeleteProject(id int) error {
    return s.repository.Delete(id)
}
