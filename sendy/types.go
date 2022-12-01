package sendy

type Subscribe struct {
	ApiKey    *string `url:"api_key"`
	ListId    string  `url:"list"`
	Email     string  `url:"email"`
	Name      *string `url:"name,omitempty"`
	Country   *string `url:"country,omitempty"`
	IPAddress *string `url:"ipaddress,omitempty"`
	Referrer  *string `url:"referrer,omitempty"`
	GDPR      bool    `url:"gdpr,omitempty"`
	Silent    *bool   `url:"silent,omitempty"`
	HoneyPot  *bool   `url:"hp,omitempty"`
	Boolean   bool    `url:"boolean,omitempty"`
}

type Unsubscribe struct {
	Email   string  `url:"email"`
	ListId  *string `url:"list"`
	Boolean bool    `url:"boolean"`
}
