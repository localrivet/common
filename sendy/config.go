package sendy

type SendyConfig struct {
	ApiKey string
	ApiUrl string
	Lists  struct {
		NonBuyerLisId string
		BuyerListId   string
	}
}
