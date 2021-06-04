package schemas

type TokenSchema struct {
	Token string `json:"token"`
}

func NewTokenSchema(token string) TokenSchema {
	return TokenSchema{token}
}