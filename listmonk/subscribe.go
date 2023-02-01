package listmonk

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (s *Listmonk) Subscribe(in *Subscribe) error {
	jsonB, err := json.Marshal(in)
	if err != nil {
		fmt.Println("marshaling data error: ", err)
	}

	u, _ := url.ParseRequestURI(s.Config.ApiUrl)
	u.Path = "/api/subscribers"

	body, err := s.do(u.String(), jsonB)

	fmt.Println(string(body))
	return err
}
