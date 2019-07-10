package services

import (
	"superstar/dao"
	"superstar/models"
	"superstar/datasource"
)

// SuperstarService ...
type SuperstarService interface {
	GetAll() []models.StarInfo
	GetLimit(count int, start int) []models.StarInfo
	Search(country string) []models.StarInfo
	Get(id int) *models.StarInfo
	Delete(id int) error
	Update(user *models.StarInfo, columns []string) error
	Create(user *models.StarInfo) error
}

type superstarService struct {
	dao *dao.SuperstarDao
}

// NewSuperstarService ...
func NewSuperstarService() SuperstarService {
	return &superstarService{
		dao: dao.NewSuperstarDao(datasource.InstanceMaster()),
	}
}

func (s *superstarService)GetAll() []models.StarInfo {
	return s.dao.GetAll()
}

func (s *superstarService)GetLimit(count int, start int) []models.StarInfo {
	return s.dao.GetLimit(count, start)
}

func (s *superstarService)Search(country string) []models.StarInfo {
	return s.dao.Search(country)
}

func (s *superstarService)Get(id int) *models.StarInfo {
	return s.dao.Get(id)
}
func (s *superstarService)Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *superstarService)Update(user *models.StarInfo, columns []string) error {
	return s.dao.Update(user, columns)
}
func (s *superstarService)Create(user *models.StarInfo) error {
	return s.dao.Create(user)
}
