package listmonk

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (s *Listmonk) Tx(in *Transactional) error {
	jsonB, err := json.Marshal(in)
	if err != nil {
		fmt.Println("marshaling data error: ", err)
	}

	u, _ := url.ParseRequestURI(s.Config.ApiUrl)
	u.Path = "/api/tx"

	body, err := s.do(u.String(), jsonB)

	fmt.Println(string(body))
	return err
}
