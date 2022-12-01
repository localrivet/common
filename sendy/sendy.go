package sendy

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Sendy struct {
	Config SendyConfig
}

func NewSendy(config SendyConfig) *Sendy {
	return &Sendy{
		Config: config,
	}
}

func (s *Sendy) do(url string, vals url.Values) (body []byte, err error) {
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(vals.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	return
}
