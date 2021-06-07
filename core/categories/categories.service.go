package categories

import (
	"crud/entities"
	"crud/providers"
	"gorm.io/gorm/clause"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}
func (s *Service) create(category entities.Category) (entities.Category, error) {
	providers.Database.Create(&category)
	return category, nil
}
func (s *Service) update(categoryUpdated entities.Category, id int) {
	providers.Database.Updates(&categoryUpdated)
}
func (s *Service) delete() {

}
func (s *Service) getAll() (categories []entities.Category) {
	providers.Database.Find(&categories).Preload(clause.Associations)
	return
}

func (s *Service) getOne(id int) (category entities.Category) {
	providers.Database.Preload("Rol").Find(&category, id)
	return
}
