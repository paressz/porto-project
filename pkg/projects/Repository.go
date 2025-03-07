package projects

import (
	"errors"
	"gorm.io/gorm"
	"math"
	"porto-project/pkg/config/database"
	"porto-project/pkg/model"
)

type Repository interface {
	CreateProject(project *model.Project) (*model.Project, error)
	GetAllProjects(lastId int) ([]model.Project, int64, error)
	GetProjectById(id string) (*model.Project, error)
	EditProject(project *model.Project) (string, error)
	DeleteProject(id string) error
}
type repository struct {
	Db *gorm.DB
}

func NewRepository() Repository {
	db := database.Connect()
	return &repository{
		db,
	}
}

func (r *repository) CreateProject(project *model.Project) (*model.Project, error) {
	return project, r.Db.Create(project).Error
}

func (r *repository) GetAllProjects(last_int_id int) ([]model.Project, int64, error) {
	var projects []model.Project
	var count int64
	r.Db.Model(&projects).Count(&count)
	itemAmount := 9
	pageCount := int64(math.Ceil(float64(count)/float64(itemAmount)))
	err := r.Db.Where("int_id > ?", last_int_id).Limit(itemAmount).Order("int_id ASC").Find(&projects).Error
	return projects, pageCount, err
}

func (r *repository) GetProjectById(id string) (*model.Project, error) {
	var project model.Project
	err := r.Db.Where("id = ?", id).First(&project).Error
	return &project, err
}

func (r *repository) EditProject(project *model.Project) (string, error) {
	return project.Id, r.Db.Model(&project).Updates(model.Project{
		Name: project.Name,
		Description: project.Description,
		ImageUrl: project.ImageUrl,
	}).Error
}

func (r *repository) DeleteProject(id string) error {
	deleteQuery := r.Db.Delete(&model.Project{}, "id = ?", id)
	rowsAffected := deleteQuery.RowsAffected
	if rowsAffected < 1 {
		return errors.New("No record with id: " + id )
	}
	return deleteQuery.Error
}