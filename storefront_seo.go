package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type StorefrontSeoSettingsService interface {
	Get(...int) (StorefrontSeoSettings, error)
	Update(StorefrontSeoSettings, ...int) (StorefrontSeoSettings, error)
	Delete(int, []string) error
}

type StorefrontSeoSettings struct {
	PageTitle       string `json:"page_title"`
	MetaDescription string `json:"meta_description"`
	WwwRedirect     string `json:"www_redirect"`
}
type StorefrontSeoSettingsResponse struct {
	Data StorefrontSeoSettings `json:"data"`
}
type StorefrontSeoSettingsOp struct {
	client *Client
}

func (s *StorefrontSeoSettingsOp) Get(channelID ...int) (StorefrontSeoSettings, error) {
	var seoResponse StorefrontSeoSettingsResponse

	var queryString string
	if len(channelID) == 1 {
		queryString = fmt.Sprintf("?channel_id=%d", channelID[0])
	}

	body, reqErr := s.client.DoRequest(http.MethodGet, fmt.Sprintf("/v3/settings/storefront/seo%s", queryString), nil)
	if reqErr != nil {
		return seoResponse.Data, reqErr
	}

	jsonErr := json.Unmarshal(body, &seoResponse)
	if jsonErr != nil {
		return seoResponse.Data, jsonErr
	}

	return seoResponse.Data, nil
}

func (s *StorefrontSeoSettingsOp) Update(settings StorefrontSeoSettings, channelID ...int) (StorefrontSeoSettings, error) {
	var seoResponse StorefrontSeoSettingsResponse

	var queryString string
	if len(channelID) == 1 {
		queryString = fmt.Sprintf("?channel_id=%d", channelID[0])
	}

	jsonBody, err := json.Marshal(settings)
	if err != nil {
		return seoResponse.Data, err
	}

	reqBody := bytes.NewReader(jsonBody)
	body, reqErr := s.client.DoRequest(http.MethodPut, fmt.Sprintf("/v3/settings/storefront/seo%s", queryString), reqBody)
	if reqErr != nil {
		return seoResponse.Data, reqErr
	}

	jsonErr := json.Unmarshal(body, &seoResponse)
	if jsonErr != nil {
		return seoResponse.Data, jsonErr
	}
	return seoResponse.Data, nil
}

func (s *StorefrontSeoSettingsOp) Delete(channelID int, keys []string) error {
	keysCsv := strings.Join(keys[:], ",")
	_, reqErr := s.client.DoRequest(http.MethodDelete, fmt.Sprintf("/v3/settings/storefront/seo?channel_id=%d&keys=%s", channelID, keysCsv), nil)
	if reqErr != nil {
		return reqErr
	}

	return nil
}
