package repositories

import (
	"msvc-syahril-app/app/interfaces"
	"msvc-syahril-app/app/models"
	"msvc-syahril-app/config"

	"gorm.io/gorm"
)

type projectRepository struct {
    db *gorm.DB
}

func NewProjectRepository() interfaces.ProjectRepository {
    return &projectRepository{
        db: config.DB,
    }
}

func (r *projectRepository) GetAll() ([]models.Project, error) {
    var projects []models.Project
    err := r.db.Find(&projects).Error
    return projects, err
}

func (r *projectRepository) GetByID(id int) (*models.Project, error) {
    var project models.Project
    err := r.db.First(&project, id).Error
    if err != nil {
        return nil, err
    }
    return &project, nil
}

func (r *projectRepository) Create(project *models.Project) error {
    return r.db.Create(project).Error
}

func (r *projectRepository) Update(project *models.Project) error {
    return r.db.Save(project).Error
}

func (r *projectRepository) Delete(id int) error {
    return r.db.Delete(&models.Project{}, id).Error
}
