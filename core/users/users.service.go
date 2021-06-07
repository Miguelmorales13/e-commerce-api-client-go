package users

import (
	"crud/core/users/dto"
	"crud/entities"
	"crud/providers"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}
func (s *Service) create(userCreated *dto.CreateUserDto) (entities.User, error) {
	user := entities.User{
		Email:          userCreated.Email,
		LastName:       userCreated.LastName,
		Name:           userCreated.Name,
		Password:       userCreated.Password,
		SecondLastName: userCreated.SecondLastName,
	}
	dbResponse := providers.Database.Create(&user)
	if dbResponse.Error != nil {
		fmt.Println("Error in creation", dbResponse.Error.Error())
		return user, &fiber.Error{
			Message: "Upps... we dont have create the user becouse exis a error. Please contact with ",
			Code:    fiber.StatusInternalServerError,
		}
	}
	return user, nil
}
func (s *Service) update(userUpdated *dto.UpdateUserDto, id uint) entities.User {
	var user entities.User
	providers.Database.First(&user, id)

	user = entities.User{
		ID:             id,
		Email:          userUpdated.Email,
		Active:         userUpdated.Active,
		LastName:       userUpdated.LastName,
		Name:           userUpdated.Name,
		SecondLastName: userUpdated.SecondLastName,
	}
	providers.Database.Updates(&user)
	return user
}

func (s *Service) delete(id string) error {
	var user entities.User
	providers.Database.First(&user, id)
	if user.ID == 0 {
		return &fiber.Error{Message: "Opps... this item has been previously deleted", Code: fiber.StatusNotFound}
	}
	providers.Database.Delete(&user)
	return nil
}

func (s *Service) getAll() (users []entities.User) {
	providers.Database.Find(&users).Preload(clause.Associations)
	return
}

func (s *Service) getOne(id uint) (user entities.User) {
	providers.Database.First(&user, id)
	return
}
