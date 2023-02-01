package n8n

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

type N8N struct {
	Config N8NConfig
}

func NewN8N(config N8NConfig) *N8N {
	return &N8N{
		Config: config,
	}
}

func (s *N8N) do(url string, jsonB []byte) (body []byte, err error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonB)) // URL-encoded payload
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
