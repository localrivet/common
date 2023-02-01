package n8n

import (
	"fmt"
	"strings"
)

type N8NConfig struct {
	InTestMode bool
	ApiUrl     string
	Username   string
	Password   string
}

func (c *N8NConfig) GetApiURL() string {
	apiUrl := fmt.Sprintf("%s/%s", c.ApiUrl, "webhook/")
	if c.InTestMode {
		apiUrl = fmt.Sprintf("%s/%s", c.ApiUrl, "webhook-test/")
	}
	return strings.ReplaceAll(apiUrl, "//", "/")
}
