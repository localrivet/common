package sendy

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

func (s *Sendy) Subscribe(in *Subscribe) error {
	in.ApiKey = &s.Config.ApiKey
	in.Boolean = true

	postData, _ := query.Values(in)
	u, _ := url.ParseRequestURI(s.Config.ApiUrl)
	u.Path = "/subscribe"

	body, err := s.do(u.String(), postData)

	fmt.Println(string(body))
	return err
}
