package n8n

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (s *N8N) Webhook(endpoint string, in map[string]interface{}) (body []byte, err error) {
	jsonB, err := json.Marshal(in)
	if err != nil {
		fmt.Println("marshaling data error: ", err)
	}

	fmt.Println("Webhook endpoint: " + endpoint)
	fmt.Println("in: ", string(jsonB))

	u, _ := url.ParseRequestURI(s.Config.ApiUrl)
	u.Path = endpoint

	return s.do(u.String(), jsonB)
}
