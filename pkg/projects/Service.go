package projects

import "porto-project/pkg/model"

type Service interface {
	CreateProject(project *model.Project) (*model.Project, error)
	GetAllProjects(lastId int) ([]model.Project, int64, error)
	GetProjectById(id string) (*model.Project, error)
	EditProject(project *model.Project) (string, error)
	DeleteProject(id string) error
}
type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo,
	}
}

func (s *service) CreateProject(project *model.Project) (*model.Project, error) {
	return s.repository.CreateProject(project)
}

func (s *service) GetAllProjects(lastId int) ([]model.Project, int64, error) {
	return s.repository.GetAllProjects(lastId)
}

func (s *service) GetProjectById(id string) (*model.Project, error) {
	return s.repository.GetProjectById(id)
}

func (s *service) EditProject(project *model.Project) (string, error) {
	return s.repository.EditProject(project)
}

func (s *service) DeleteProject(id string) error {
	return s.repository.DeleteProject(id)
}