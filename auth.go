package main

import "net/http"

type AuthInfo interface {
	Attach(req *http.Request)
}

type cookie struct {
	cookieName  string
	cookieValue string
}

type token struct {
	tokenPrefix string
	tokenValue  string
}

func NewAuthInfo(info map[string]string) AuthInfo {
	var authInfo AuthInfo
	if len(info) != 0 {
		if info["type"] == "Cookie" {
			authInfo = cookie{cookieName: info["name"], cookieValue: info["value"]}
		} else if info["type"] == "OAuth 2.0" {
			authInfo = token{tokenPrefix: info["prefix"], tokenValue: info["value"]}
		}
		return authInfo
	} else {
		return nil
	}
}

func (c cookie) Attach(req *http.Request) {
	cookie := &http.Cookie{Name: c.cookieName, Value: c.cookieValue}
	req.AddCookie(cookie)
}

func (t token) Attach(req *http.Request) {
	value := t.tokenPrefix + " " + t.tokenValue
	req.Header.Add("Authorization", value)
}
