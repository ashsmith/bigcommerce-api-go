package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type StorefrontSearchSettingsService interface {
	Get(...int) (StorefrontSearchSettings, error)
	Update(StorefrontSearchSettings, ...int) (StorefrontSearchSettings, error)
	Delete(int, []string) error
}

type StorefrontSearchSettings struct {
	DefaultProductSort string `json:"default_product_sort"`
}

type StorefrontSearchSettingsResponse struct {
	Data StorefrontSearchSettings `json:"data"`
}

type StorefrontSearchSettingsOp struct {
	client *Client
}

func (s *StorefrontSearchSettingsOp) Get(channelID ...int) (StorefrontSearchSettings, error) {
	var searchResponse StorefrontSearchSettingsResponse

	var queryString string
	if len(channelID) == 1 {
		queryString = fmt.Sprintf("?channel_id=%d", channelID[0])
	}

	body, reqErr := s.client.DoRequest(http.MethodGet, fmt.Sprintf("/v3/settings/storefront/search%s", queryString), nil)
	if reqErr != nil {
		return searchResponse.Data, reqErr
	}

	jsonErr := json.Unmarshal(body, &searchResponse)
	if jsonErr != nil {
		return searchResponse.Data, jsonErr
	}

	return searchResponse.Data, nil
}

func (s *StorefrontSearchSettingsOp) Update(search StorefrontSearchSettings, channelID ...int) (StorefrontSearchSettings, error) {
	var searchResponse StorefrontSearchSettingsResponse

	var queryString string
	if len(channelID) == 1 {
		queryString = fmt.Sprintf("?channel_id=%d", channelID[0])
	}

	jsonBody, err := json.Marshal(search)
	if err != nil {
		return searchResponse.Data, err
	}

	reqBody := bytes.NewReader(jsonBody)
	body, reqErr := s.client.DoRequest(http.MethodPut, fmt.Sprintf("/v3/settings/storefront/search%s", queryString), reqBody)
	if reqErr != nil {
		return searchResponse.Data, reqErr
	}

	jsonErr := json.Unmarshal(body, &searchResponse)
	if jsonErr != nil {
		return searchResponse.Data, jsonErr
	}
	return searchResponse.Data, nil
}

func (s *StorefrontSearchSettingsOp) Delete(channelID int, keys []string) error {
	keysCsv := strings.Join(keys[:], ",")
	_, reqErr := s.client.DoRequest(http.MethodDelete, fmt.Sprintf("/v3/settings/storefront/search?channel_id=%d&keys=%s", channelID, keysCsv), nil)
	if reqErr != nil {
		return reqErr
	}

	return nil
}
