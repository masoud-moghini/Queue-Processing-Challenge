package dto

import (
	"net/http"

	"github.com/go-chi/render"
)

type UserRequest struct {
	DataToBeProcess string
}

// Bind implements render.Binder.
func (u *UserRequest) Bind(r *http.Request) error {
	return nil
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error, statusCode int, statusText string) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: statusCode,
		StatusText:     statusText,
		ErrorText:      err.Error(),
	}
}

type UserResponse struct {
	Data string
}

func (u UserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
