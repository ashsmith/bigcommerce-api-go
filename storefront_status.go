package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type StorefrontStatusService interface {
	Get(...int) (StorefrontStatus, error)
	Update(StorefrontStatus, ...int) (StorefrontStatus, error)
	Delete(int, []string) error
}

type StorefrontStatus struct {
	DownForMaintenanceMessage string `json:"down_for_maintenance"`
	PrelaunchMessage          string `json:"prelaunch_message"`
	PrelaunchPassword         string `json:"prelaunch_message"`
}

type StorefrontStatusResponse struct {
	Data StorefrontStatus `json:"data"`
}

type StorefrontStatusOp struct {
	client *Client
}

func (s *StorefrontStatusOp) Get(channelID ...int) (StorefrontStatus, error) {
	var statusResponse StorefrontStatusResponse

	var queryString string
	if len(channelID) == 1 {
		queryString = fmt.Sprintf("?channel_id=%d", channelID[0])
	}

	body, reqErr := s.client.DoRequest(http.MethodGet, fmt.Sprintf("/v3/settings/storefront/status%s", queryString), nil)
	if reqErr != nil {
		return statusResponse.Data, reqErr
	}

	jsonErr := json.Unmarshal(body, &statusResponse)
	if jsonErr != nil {
		return statusResponse.Data, jsonErr
	}

	return statusResponse.Data, nil
}

func (s *StorefrontStatusOp) Update(status StorefrontStatus, channelID ...int) (StorefrontStatus, error) {
	var statusResponse StorefrontStatusResponse

	var queryString string
	if len(channelID) == 1 {
		queryString = fmt.Sprintf("?channel_id=%d", channelID[0])
	}

	jsonBody, err := json.Marshal(status)
	if err != nil {
		return statusResponse.Data, err
	}

	reqBody := bytes.NewReader(jsonBody)
	body, reqErr := s.client.DoRequest(http.MethodPut, fmt.Sprintf("/v3/settings/storefront/status%s", queryString), reqBody)
	if reqErr != nil {
		return statusResponse.Data, reqErr
	}

	jsonErr := json.Unmarshal(body, &statusResponse)
	if jsonErr != nil {
		return statusResponse.Data, jsonErr
	}

	return statusResponse.Data, nil
}

func (s *StorefrontStatusOp) Delete(channelID int, keys []string) error {
	keysCsv := strings.Join(keys[:], ",")
	_, reqErr := s.client.DoRequest(http.MethodDelete, fmt.Sprintf("/v3/settings/storefront/status?channel_id=%d&keys=%s", channelID, keysCsv), nil)
	if reqErr != nil {
		return reqErr
	}

	return nil
}
