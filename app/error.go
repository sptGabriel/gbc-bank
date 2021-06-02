package app

type group string

const (
	domain   group = "domain"
	internal group = "internal"
)

type Error struct {
	group   group
	code    int
	message string
	parent  error
}

func (e Error) Error() string {
	return e.message
}

func (e Error) Parent() error {
	return e.parent
}

func (e Error) IsDomain() bool {
	return e.group == "domain"
}

func (e Error) IsInternal() bool {
	return e.group == "internal"
}

func (e Error) Code() int {
	if e.code != 0 {
		return e.code
	}
	return 500
}

func GetErrorCode(err error) int {
	e, ok := err.(*Error)
	if !ok {
		return 500
	}
	if e.code != 0 {
		return e.code
	}

	return GetErrorCode(e.parent)
}
func NewDomainError(m string) error {
	return &Error{
		group:   domain,
		message: m,
		code:    400,
	}
}

func NewInternalError(m string, p error) error {
	return &Error{
		group:   internal,
		message: m,
		parent:  p,
		code:    500,
	}
}

func NewConflictError(e string) error {
	entity := ""
	if e != "" {
		entity = e + " "
	}
	return &Error{
		group:   domain,
		code:    409,
		message: entity + "already exists",
	}
}

func NewNotFoundError(e string) error {
	entity := ""
	if e != "" {
		entity = e + " "
	}
	return &Error{
		group:   domain,
		code:    404,
		message: entity + "not found",
	}
}

func NewMalformedJSONError() error {
	return &Error{
		group:   domain,
		code:    400,
		message: "The request cannot be fulfilled due to bad syntax",
	}
}
