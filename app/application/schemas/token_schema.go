package schemas

type tokenSchema struct {
	Token string `json:"token"`
}

func NewTokenSchema(token string) tokenSchema {
	return tokenSchema{token}
}