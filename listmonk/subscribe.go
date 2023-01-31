package listmonk

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

func (s *Listmonk) Subscribe(in *Subscribe) error {
	postData, _ := query.Values(in)
	u, _ := url.ParseRequestURI(s.Config.ApiUrl)
	u.Path = "/api/subscribers"

	body, err := s.do(u.String(), postData)

	fmt.Println(string(body))
	return err
}
