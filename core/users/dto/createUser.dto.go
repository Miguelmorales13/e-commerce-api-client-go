package dto

type CreateUserDto struct {
	Name           string `validate:"required" json:"name"`
	LastName       string `validate:"required" json:"lastName"`
	SecondLastName string `validate:"required" json:"secondLastName"`
	Email          string `validate:"required" json:"email"`
	Password       string `validate:"required" json:"password"`
}
