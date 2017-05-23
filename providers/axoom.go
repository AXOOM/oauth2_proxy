package providers

import (
	"log"
	"fmt"
	"net/http"

	"github.com/bitly/oauth2_proxy/api"
)

type AxoomProvider struct {
	*ProviderData
}

func NewAxoomProvider(p *ProviderData) *AxoomProvider {
	p.ProviderName = "AXOOM"
	if p.LoginURL.String() == "" || p.RedeemURL.String() == "" || p.ProfileURL.String() == "" || p.ValidateURL.String() == "" {
		fmt.Errorf("Please provide all urls on start when using the AXOOM Provider")
		log.Printf("Please provide all urls on start when using the AXOOM Provider")
	}
	if p.Scope == "" {
		p.Scope = "openid profile"
	}
	return &AxoomProvider{ProviderData: p}
}

func getAxoomHeader(access_token string) http.Header {
	header := make(http.Header)
	header.Set("Authorization", fmt.Sprintf("Bearer %s", access_token))
	return header
}

func (p *AxoomProvider) GetEmailAddress(s *SessionState) (string, error) {
	req, err := http.NewRequest("GET",
		p.ProfileURL.String(), nil)
    req.Header = getAxoomHeader(s.AccessToken)
	if err != nil {
		log.Printf("failed building request %s", err)
		return "", err
	}
	json, err := api.Request(req)
	if err != nil {
		log.Printf("failed making request %s", err)
		return "", err
	}
	return json.Get("sub").String()
}
