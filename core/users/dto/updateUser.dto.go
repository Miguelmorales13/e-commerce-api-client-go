package dto

type UpdateUserDto struct {
	Name           string `validate:"required" json:"name"`
	LastName       string `validate:"required" json:"lastName"`
	SecondLastName string `validate:"required" json:"secondLastName"`
	Email          string `validate:"required" json:"email"`
	Active         bool   `validate:"required" json:"active"`
}
