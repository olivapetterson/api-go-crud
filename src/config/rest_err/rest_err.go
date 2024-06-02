package rest_err

import "golang.org/x/text/message"

// struct of json error message (404, 503, etc)
type RestErr struct {
	Message string   `json:"message"` // [rint the message of request
	Err     string   `json:"error"` // print the error of request
	Code    int    `json:"code"` // print the code of request
	Causes  []Causes `json:"causes"` // print the error causes inside of the api
}

// this struct will send the message about the field in use, example bellow
type Causes struct {
	Field   string `json:"field"` // the user send a incorrect password, here we will send a message about it
	Message string `json:"message` // here too ^
}

// a builder to our object REST error
type NewRestErr (message, err string, code int64, causes []Causes) *RestErr{
	return &RestErr{
		Message: message,
		Err: err,
		Code: code,
		Causes: causes,
	}
}

// func to satisfact any lib whos receive error has parameter
func (r *RestErr) Error() string{
	return r.Message
}

// func returning error 400
func NewBadRequestError(message string) *RestErr{
	return &RestErr{
		Message: message,
		Err: "bad_request",
		Code: http.StatusBadRequest,
	}
}

// func returning error 400, but the focus its validate fields
func NewBadRequestValidationError(message string, causes []Causes) *RestErr{
	return &RestErr{
		Message: message,
		Err: "bad_request",
		Code: http.StatusBadRequest,
		Causes: causes, // cause its here because, he will validate fields at app
	}
}

// func returning error 403 - focus on JWT
func NewForbiddenRequestError(message string) *RestErr{
	return &RestErr{
		Message: message
		Err: "forbidden_request",
		Code: http.StatusForbidden,
	}
}

// func returning error 404
func NewNotFoundRequestError(message string) *RestErr{
	return &RestErr{
		Message: message
		Err: "not_found_request",
		Code: http.StatusNotFound,
	}
}

// func returning error 500
func NewInternalServerError(message string) *RestErr{
	return &RestErr{
		Message: message,
		Err: "internal_server_error",
		Code: http.StatusInternalServerError,
	}
}
