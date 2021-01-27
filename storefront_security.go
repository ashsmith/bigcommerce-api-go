package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type StorefrontSecuritySettingsService interface {
	Get(...int) (StorefrontSecuritySettings, error)
	Update(StorefrontSecuritySettings, ...int) (StorefrontSecuritySettings, error)
	Delete(int, []string) error
}

type HSTSOptions struct {
	Enabled           bool `json:"enabled"`
	MaxAgeMonths      int  `json:"max_age_months"`
	IncludeSubDomains bool `json:"includeSubDomains"`
}

type StorefrontSecuritySettings struct {
	SitewideHTTPSEnabled bool        `json:"sitewide_https_enabled"`
	CSPHeader            string      `json:"csp_header"`
	HSTS                 HSTSOptions `json:"hsts"`
}

type StorefrontSecuritySettingsResponse struct {
	Data StorefrontSecuritySettings `json:"data"`
}

type StorefrontSecuritySettingsOp struct {
	client *Client
}

func (s *StorefrontSecuritySettingsOp) Get(channelID ...int) (StorefrontSecuritySettings, error) {
	var securityResponse StorefrontSecuritySettingsResponse

	var queryString string
	if len(channelID) == 1 {
		queryString = fmt.Sprintf("?channel_id=%d", channelID[0])
	}

	body, reqErr := s.client.DoRequest(http.MethodGet, fmt.Sprintf("/v3/settings/storefront/security%s", queryString), nil)
	if reqErr != nil {
		return securityResponse.Data, reqErr
	}

	jsonErr := json.Unmarshal(body, &securityResponse)
	if jsonErr != nil {
		return securityResponse.Data, jsonErr
	}

	return securityResponse.Data, nil
}

func (s *StorefrontSecuritySettingsOp) Update(security StorefrontSecuritySettings, channelID ...int) (StorefrontSecuritySettings, error) {
	var securityResponse StorefrontSecuritySettingsResponse

	var queryString string
	if len(channelID) == 1 {
		queryString = fmt.Sprintf("?channel_id=%d", channelID[0])
	}

	jsonBody, err := json.Marshal(security)
	if err != nil {
		return securityResponse.Data, err
	}

	reqBody := bytes.NewReader(jsonBody)
	body, reqErr := s.client.DoRequest(http.MethodPut, fmt.Sprintf("/v3/settings/storefront/security%s", queryString), reqBody)
	if reqErr != nil {
		return securityResponse.Data, reqErr
	}

	jsonErr := json.Unmarshal(body, &securityResponse)
	if jsonErr != nil {
		return securityResponse.Data, jsonErr
	}
	return securityResponse.Data, nil
}

func (s *StorefrontSecuritySettingsOp) Delete(channelID int, keys []string) error {
	keysCsv := strings.Join(keys[:], ",")
	_, reqErr := s.client.DoRequest(http.MethodDelete, fmt.Sprintf("/v3/settings/storefront/security?channel_id=%d&keys=%s", channelID, keysCsv), nil)
	if reqErr != nil {
		return reqErr
	}

	return nil
}
