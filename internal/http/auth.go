package http

import (
	"net/http"
)

const (
	AuthTypeCookie = "Cookie"
	AuthTypeOAuth2 = "OAuth 2.0"
)

// AuthMethods is a store of prepared information of methods for authentication
type AuthMethods map[string]map[string]string

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

func (m AuthMethods) AuthInfo(method string) AuthInfo {
	var authInfo AuthInfo
	info := m[method]

	if len(info) != 0 {
		if info["type"] == AuthTypeCookie {
			authInfo = cookie{cookieName: info["name"], cookieValue: info["value"]}
		} else if info["type"] == AuthTypeOAuth2 {
			authInfo = token{tokenPrefix: info["prefix"], tokenValue: info["value"]}
		}
		return authInfo
	} else {
		return nil
	}
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
