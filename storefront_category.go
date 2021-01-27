package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type StorefrontCategorySettingsService interface {
	Get(...int) (StorefrontCategorySettings, error)
	Update(StorefrontCategorySettings, ...int) (StorefrontCategorySettings, error)
	Delete(int, []string) error
}

type StorefrontCategorySettings struct {
	DefaultProductSort string `json:"default_product_sort"`
	CategoryTreeDepth  string `json:"category_tree_depth"`
}

type StorefrontCategorySettingsResponse struct {
	Data StorefrontCategorySettings `json:"data"`
}

type StorefrontCategorySettingsOp struct {
	client *Client
}

func (s *StorefrontCategorySettingsOp) Get(channelID ...int) (StorefrontCategorySettings, error) {
	var categoryResponse StorefrontCategorySettingsResponse

	var queryString string
	if len(channelID) == 1 {
		queryString = fmt.Sprintf("?channel_id=%d", channelID[0])
	}

	body, reqErr := s.client.DoRequest(http.MethodGet, fmt.Sprintf("/v3/settings/storefront/category%s", queryString), nil)
	if reqErr != nil {
		return categoryResponse.Data, reqErr
	}

	jsonErr := json.Unmarshal(body, &categoryResponse)
	if jsonErr != nil {
		return categoryResponse.Data, jsonErr
	}

	return categoryResponse.Data, nil
}

func (s *StorefrontCategorySettingsOp) Update(category StorefrontCategorySettings, channelID ...int) (StorefrontCategorySettings, error) {
	var categoryResponse StorefrontCategorySettingsResponse

	var queryString string
	if len(channelID) == 1 {
		queryString = fmt.Sprintf("?channel_id=%d", channelID[0])
	}

	jsonBody, err := json.Marshal(category)
	if err != nil {
		return categoryResponse.Data, err
	}

	reqBody := bytes.NewReader(jsonBody)
	body, reqErr := s.client.DoRequest(http.MethodPut, fmt.Sprintf("/v3/settings/storefront/category%s", queryString), reqBody)
	if reqErr != nil {
		return categoryResponse.Data, reqErr
	}

	jsonErr := json.Unmarshal(body, &categoryResponse)
	if jsonErr != nil {
		return categoryResponse.Data, jsonErr
	}
	return categoryResponse.Data, nil
}

func (s *StorefrontCategorySettingsOp) Delete(channelID int, keys []string) error {
	keysCsv := strings.Join(keys[:], ",")
	_, reqErr := s.client.DoRequest(http.MethodDelete, fmt.Sprintf("/v3/settings/storefront/category?channel_id=%d&keys=%s", channelID, keysCsv), nil)
	if reqErr != nil {
		return reqErr
	}

	return nil
}
