package responses

type Response struct {
	Status int
	Error  error
	Data   interface{}
}