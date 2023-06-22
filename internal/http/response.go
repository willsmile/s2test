package http

// Response constains status and body of a HTTP Response
type Response struct {
	Body   string
	Status string
}

// newResponse creates an empty response
func newResponse() *Response {
	return &Response{}
}

// DefaultResponse creates a response with default values
func DefaultResponse() *Response {
	return &Response{
		Body:   "None",
		Status: "None",
	}
}
