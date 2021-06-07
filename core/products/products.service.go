package products

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
func (s *Service) create(product entities.Product) (entities.Product, error) {
	providers.Database.Create(&product)
	return product, nil
}
func (s *Service) update(productUpdated entities.Product, id int) {
	providers.Database.Updates(&productUpdated)
}
func (s *Service) delete() {

}
func (s *Service) getAll() (products []entities.Product) {
	providers.Database.Find(&products).Preload(clause.Associations)
	return
}

func (s *Service) getOne(id int) (product entities.Product) {
	providers.Database.Preload("Rol").Find(&product, id)
	return
}
