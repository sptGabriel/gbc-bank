package dtos

type SignInDTO struct {
	CPF   string `json:"cpf" validate:"required, min=9,max=12"`
	Secret    string `json:"cpf" validate:"required,min=8"`
}
