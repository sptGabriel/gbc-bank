package dtos

type SignInDTO struct {
	CPF    string `json:"cpf" validate:"required"`
	Secret string `json:"secret" validate:"required,min=8"`
}
