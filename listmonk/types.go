package listmonk

type ContentType string

const (
	ContentTypeHtml     ContentType = "html"
	ContentTypeMarkdown ContentType = "markdown"
	ContentTypePlain    ContentType = "plain"
)

type Transactional struct {
	SubscriberEmail string      `url:"subscriber_email"`        //	String	Optional	E-mail of the subscriber. Either this or subscriber_id should be passed.
	SubscriberId    int64       `url:"subscriber_id,omitempty"` //	Number	Optional	ID of the subscriber. Either this or subscriber_email should be passed.
	TemplateId      int64       `url:"template_id"`             // 	Number	Required	ID of the transactional template to use in the message.
	FromEmail       string      `url:"from_email,omitempty"`    //	String	Optional	Optional from email. eg: Company <email@company.com>
	Data            string      `url:"data,omitempty"`          //	Map	Optional	Optional data in {} nested map. Available in the template as {{ .Tx.Data.* }}
	Headers         string      `url:"headers,omitempty"`       //	[]Map	Optional	Optional array of mail headers. [{"key": "value"}, {"key": "value"}]
	Messenger       string      `url:"messenger,omitempty"`     //	String	Optional	Messenger to use to send the message. Default value is email.
	ContentType     ContentType `url:"content_type,omitempty"`  //	String	Optional	html, markdown, plain
}

type Subscribe struct {
	Email                   string  `json:"email"`                              // 	String	Required	The email address of the new susbcriber.
	Name                    string  `json:"name"`                               // 	String	Required	The name of the new subscriber.
	Status                  string  `json:"status"`                             // 	String	Required	The status of the new subscriber. Can be enabled, disabled or blocklisted.
	ListIDs                 []int64 `json:"lists,omitempty"`                    // 	Numbers	Optional	Array of list IDs to subscribe to (marked as unconfirmed by default).
	Attribs                 string  `json:"attribs,omitempty"`                  // 	json	Optional	JSON list containing new subscriber's attributes.
	PreconfirmSubscriptions bool    `json:"preconfirm_subscriptions,omitempty"` //	Bool	Optional	If true, marks subscriptsions as confirmed and no-optin e-mails are sent for double opt-in lists.
}
