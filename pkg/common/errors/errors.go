package errors

type Unknown struct {
	Message string `json:"message" example:"unknown error: {reason}" format:"string"`
}

func (e Unknown) Error() string {
	return "unknown error: " + e.Message
}

type NotFoundURL struct {
	Message string `json:"message" example:"unsupported url: {reason}" format:"string"`
}

func (e NotFoundURL) Error() string {
	return "unsupported url: " + e.Message
}

type InvalidMethod struct {
	Message string `json:"message" example:"invalid method: {reason}" format:"string"`
}

func (e InvalidMethod) Error() string {
	return "invalid method: " + e.Message
}

type InvalidMLAppName struct {
	Message string `json:"message" example:"invalid ML App name: {reason}" format:"string"`
}

func (e InvalidMLAppName) Error() string {
	return "invalid ML App name: " + e.Message
}

type NotFoundMLApp struct {
	Message string `json:"message" example:"not found ML app: {reason}" format:"string"`
}

func (e NotFoundMLApp) Error() string {
	return "not found ML app: " + e.Message
}

type IOError struct {
	Message string `json:"message" example:"io error: {reason}" format:"string"`
}

func (e IOError) Error() string {
	return "io error: " + e.Message
}

type TimeoutError struct {
	Message string `json:"message" example:"time out error: {reason}" format:"string"`
}

func (e TimeoutError) Error() string {
	return "time out error: " + e.Message
}

type InternalServerError struct {
	Message string `json:"message" example:"internal server error: {reason}" format:"string"`
}

func (e InternalServerError) Error() string {
	return "internal server error: " + e.Message
}
