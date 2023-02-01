package sendy

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

func (s *Sendy) Subscribe(in *Subscribe, attribs map[string]interface{}) error {
	in.ApiKey = &s.Config.ApiKey
	in.Boolean = true

	postData, _ := query.Values(in)
	u, _ := url.ParseRequestURI(s.Config.ApiUrl)
	u.Path = "/subscribe"

	for key, value := range attribs {
		postData.Set(key, value.(string))
	}

	body, err := s.do(u.String(), postData)

	fmt.Println(string(body))
	return err
}
