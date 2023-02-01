package sendy

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

func (s *Sendy) Status(in *Status) (response string, err error) {
	in.ApiKey = &s.Config.ApiKey

	postData, _ := query.Values(in)
	u, _ := url.ParseRequestURI(s.Config.ApiUrl)
	u.Path = "/api/subscribers/subscription-status.php"

	body, err := s.do(u.String(), postData)
	return string(body), err
}
