package responses

type Response struct {
	Status int
	Error  error
	Data   interface{}
}

type Error struct {
	Message string `json:"Message"`
}
