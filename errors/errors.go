package errors

var GlobalErrorsHandler = &ErrorHandler{
	Errors: make(chan Error),
}

type HandleError error

type Error struct {
	Err any `json:"err"`

	Location string `json:"location"`
}

type ErrorHandler struct {
	Errors chan Error
}

func (err ErrorHandler) SendErrors(errors []Error) {
	for _, el := range errors {
		err.Errors <- el
	}
}

func (err ErrorHandler) SendError(err_ Error) {
	err.Errors <- err_
}

func New() *ErrorHandler {
	return GlobalErrorsHandler
}
