package http

import (
	"net/http"
)

type AuthInfo interface {
	attach(req *Request)
}

type cookie struct {
	cookieName  string
	cookieValue string
}

type token struct {
	tokenPrefix string
	tokenValue  string
}

func (c cookie) attach(req *Request) {
	cookie := &http.Cookie{Name: c.cookieName, Value: c.cookieValue}
	req.Cookies = append(req.Cookies, cookie)
}

func (t token) attach(req *Request) {
	value := t.tokenPrefix + " " + t.tokenValue
	req.Headers.Add("Authorization", value)
}
