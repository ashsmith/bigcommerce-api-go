package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type StorefrontRobotsTxtSettingsService interface {
	Get(...int) (StorefrontRobotsTxtSettings, error)
	Update(StorefrontRobotsTxtSettings, ...int) (StorefrontRobotsTxtSettings, error)
	Delete(int, []string) error
}

type StorefrontRobotsTxtSettings struct {
	RobotsTxtSSL string `json:"robots_txt_ssl"`
}

type StorefrontRobotsTxtSettingsResponse struct {
	Data StorefrontRobotsTxtSettings `json:"data"`
}

type StorefrontRobotsTxtSettingsOp struct {
	client *Client
}

func (s *StorefrontRobotsTxtSettingsOp) Get(channelID ...int) (StorefrontRobotsTxtSettings, error) {
	var robotsResponse StorefrontRobotsTxtSettingsResponse

	var queryString string
	if len(channelID) == 1 {
		queryString = fmt.Sprintf("?channel_id=%d", channelID[0])
	}

	body, reqErr := s.client.DoRequest(http.MethodGet, fmt.Sprintf("/v3/settings/storefront/robotstxt%s", queryString), nil)
	if reqErr != nil {
		return robotsResponse.Data, reqErr
	}

	jsonErr := json.Unmarshal(body, &robotsResponse)
	if jsonErr != nil {
		return robotsResponse.Data, jsonErr
	}

	return robotsResponse.Data, nil
}

func (s *StorefrontRobotsTxtSettingsOp) Update(robots StorefrontRobotsTxtSettings, channelID ...int) (StorefrontRobotsTxtSettings, error) {
	var robotsResponse StorefrontRobotsTxtSettingsResponse

	var queryString string
	if len(channelID) == 1 {
		queryString = fmt.Sprintf("?channel_id=%d", channelID[0])
	}

	jsonBody, err := json.Marshal(robots)
	if err != nil {
		return robotsResponse.Data, err
	}

	reqBody := bytes.NewReader(jsonBody)
	body, reqErr := s.client.DoRequest(http.MethodPut, fmt.Sprintf("/v3/settings/storefront/robotstxt%s", queryString), reqBody)
	if reqErr != nil {
		return robotsResponse.Data, reqErr
	}

	jsonErr := json.Unmarshal(body, &robotsResponse)
	if jsonErr != nil {
		return robotsResponse.Data, jsonErr
	}
	return robotsResponse.Data, nil
}

func (s *StorefrontRobotsTxtSettingsOp) Delete(channelID int, keys []string) error {
	keysCsv := strings.Join(keys[:], ",")
	_, reqErr := s.client.DoRequest(http.MethodDelete, fmt.Sprintf("/v3/settings/storefront/robotstxt?channel_id=%d&keys=%s", channelID, keysCsv), nil)
	if reqErr != nil {
		return reqErr
	}

	return nil
}
