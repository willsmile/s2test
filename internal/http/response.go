package http

// Response constains status and body of a HTTP Response
type Response struct {
	Body   string
	Status string
}

// DefaultResponse creates a response with default values
func DefaultResponse() *Response {
	return &Response{
		Body:   "None",
		Status: "None",
	}
}
