package http

import (
	"net/http"
)

type AuthInfo interface {
	Attach(req *Request)
}

type cookie struct {
	cookieName  string
	cookieValue string
}

type token struct {
	tokenPrefix string
	tokenValue  string
}

func (c cookie) Attach(req *Request) {
	cookie := &http.Cookie{Name: c.cookieName, Value: c.cookieValue}
	cookies := []*http.Cookie{cookie}
	req.Cookies = cookies
}

func (t token) Attach(req *Request) {
	value := t.tokenPrefix + " " + t.tokenValue
	req.Headers.Add("Authorization", value)
}
