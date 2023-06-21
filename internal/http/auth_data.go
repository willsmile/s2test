package http

const (
	AuthTypeCookie = "Cookie"
	AuthTypeOAuth2 = "OAuth 2.0"
)

// AuthData is data for authentication
type AuthData map[string]string

// AuthDataset is a dataset to store AuthData
type AuthDataset map[string]AuthData

func (s AuthDataset) Select(auth string) AuthData {
	return s[auth]
}

func (data AuthData) isCookie() bool {
	return data["type"] == AuthTypeCookie
}

func (data AuthData) isOAuth2() bool {
	return data["type"] == AuthTypeOAuth2
}

func (data AuthData) NewAuthInfo() AuthInfo {
	var info AuthInfo

	if data.isCookie() {
		info = cookie{cookieName: data["name"], cookieValue: data["value"]}
	} else if data.isOAuth2() {
		info = token{tokenPrefix: data["prefix"], tokenValue: data["value"]}
	}
	return info
}
