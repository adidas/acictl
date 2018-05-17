package httpclient

type Client struct {
	Url         string
	Method      string
	ContentType string
	Password    string
	User        string
}

// NewClient create a new client
func NewClient(url, method, contentType, user, password string) Client {
	return Client{
		Url:         url,
		Method:      method,
		ContentType: contentType,
		Password:    password,
		User:        user,
	}
}
