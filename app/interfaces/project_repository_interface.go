package interfaces

import "msvc-syahril-app/app/models"

type ProjectRepository interface {
    GetAll() ([]models.Project, error)
    GetByID(id int) (*models.Project, error)
    Create(project *models.Project) error
    Update(project *models.Project) error
    Delete(id int) error
}
