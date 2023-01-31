package listmonk

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Listmonk struct {
	Config ListmonkConfig
}

func NewListmonk(config ListmonkConfig) *Listmonk {
	return &Listmonk{
		Config: config,
	}
}

func (s *Listmonk) do(url string, vals url.Values) (body []byte, err error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	r, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(vals.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	r.SetBasicAuth(s.Config.Username, s.Config.Password)

	resp, err := client.Do(r)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	return
}
