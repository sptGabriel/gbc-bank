package dtos

type CreateAccountDTO struct {
	Name   string `json:"name" validate:"required,min=10"`
	CPF    string `json:"cpf" validate:"required,min=9,max=12"`
	Secret string `json:"secret" validate:"required,min=8"`
}
