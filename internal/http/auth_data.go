package http

const (
	AuthTypeCookie = "Cookie"
	AuthTypeOAuth2 = "OAuth 2.0"
)

// AuthData is data for authentication
type AuthData map[string]string

// AuthDataset is a dataset to store AuthData
type AuthDataset map[string]AuthData

func (set AuthDataset) Select(auth []string) AuthDataset {
	result := make(AuthDataset)

	for _, a := range auth {
		if val, found := set[a]; found {
			result[a] = val
		}
	}

	return result
}

func (data AuthData) isCookie() bool {
	return data["type"] == AuthTypeCookie
}

func (data AuthData) isOAuth2() bool {
	return data["type"] == AuthTypeOAuth2
}

func (set AuthDataset) NewAuthInfo() []AuthInfo {
	var info []AuthInfo

	for _, d := range set {
		info = append(info, d.generate())
	}

	return info
}

func (data AuthData) generate() AuthInfo {
	var info AuthInfo

	if data.isCookie() {
		info = cookie{cookieName: data["name"], cookieValue: data["value"]}
	} else if data.isOAuth2() {
		info = token{tokenPrefix: data["prefix"], tokenValue: data["value"]}
	}

	return info
}
