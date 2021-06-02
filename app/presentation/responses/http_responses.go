package responses

import (
	"github.com/sptGabriel/banking/app"
	"net/http"
)

func Deleted(d interface{}) Response {
	return Response{
		Status: http.StatusNoContent,
		Data:   d,
	}
}

func Created(d interface{}) Response {
	return Response{
		Status: http.StatusCreated,
		Data:   d,
	}
}

func OK(d interface{}) Response {
	return Response{
		Status: http.StatusOK,
		Data:   d,
	}
}

func Updated(d interface{}) Response {
	return Response{
		Status: http.StatusNoContent,
		Data:   d,
	}
}

func IsError(e error) Response {
	return Response{
		Status: app.GetErrorCode(e),
		Data:   Error{e.Error()},
	}
}