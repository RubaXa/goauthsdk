package oauth

import (
	"fmt"
	"strings"
	"time"
)

type OURL struct {
	service     Service
	base        string
	state       string
	redirectUri string
}

func NewAuthURL(s Service) *OURL {
	return &OURL{
		service: s,
		base: fmt.Sprintf(
			"%s/login?client_id=%s&response_type=%s&scope=%s",
			BaseURL,
			s.ClientID,
			"code",
			strings.Join(s.Scopes, " "),
		),
		state: fmt.Sprintf("%d", time.Now().Unix()),
	}
}

func (o *OURL) SetState(v string) *OURL {
	o.state = v
	return o
}

func (o *OURL) SetRedirectURI(v string) *OURL {
	o.redirectUri = v
	return o
}

func (o *OURL) String() string {
	u := o.base + "&state=" + o.state + "&redirect_uri="
	if o.redirectUri != "" {
		u += o.redirectUri
	} else {
		u += o.service.RedirectURI.Get()
	}

	return u
}
