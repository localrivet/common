package listmonk

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

func (s *Listmonk) Tx(in *Transactional) error {
	postData, err := query.Values(in)
	if err != nil {
		fmt.Println("conversion error: ", err)
	}

	fmt.Println("Sending transactional email to " + in.SubscriberEmail)
	fmt.Println(postData)

	u, _ := url.ParseRequestURI(s.Config.ApiUrl)
	u.Path = "/api/tx"

	body, err := s.do(u.String(), postData)

	fmt.Println(string(body))
	return err
}
