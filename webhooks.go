package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type WebhooksService interface {
	Get(int64, ...interface{}) (Webhook, error)
	List(...interface{}) ([]Webhook, error)
	Create(Webhook, ...interface{}) (Webhook, error)
	Update(Webhook, ...interface{}) (Webhook, error)
	Delete(int64, ...interface{}) error
}

type PaginationResult struct {
	Offset     int64 `json:"offset"`
	Limit      int64 `json:"limit"`
	TotalItems int64 `json:"total_items"`
}

type MetaResult struct {
	Pagination PaginationResult `json:"pagination"`
}

type GetWebhookResponse struct {
	Data Webhook `json:"data"`
}

type ListWebhookResponse struct {
	Data []Webhook  `json:"data"`
	Meta MetaResult `json:"meta"`
}

// Webhook structure.
type Webhook struct {
	ID          int64             `json:"id,omitempty"`
	ClientID    string            `json:"client_id,omitempty"`
	StoreHash   string            `json:"store_hash,omitempty"`
	CreatedAt   int64             `json:"created_at,omitempty"`
	UpdatedAt   int64             `json:"updated_at,omitempty"`
	Scope       string            `json:"scope"`
	Destination string            `json:"destination"`
	IsActive    bool              `json:"is_active"`
	Headers     map[string]string `json:"headers"`
}

type WebhooksServiceOp struct {
	client *Client
}

// Get will fetch a single webhook by the provided ID.
func (s *WebhooksServiceOp) Get(id int64, options ...interface{}) (Webhook, error) {
	var webhookResponse GetWebhookResponse
	body, reqErr := s.client.DoRequest(http.MethodGet, fmt.Sprintf("/v3/hooks/%d", id), nil)
	if reqErr != nil {
		return webhookResponse.Data, reqErr
	}
	jsonErr := json.Unmarshal(body, &webhookResponse)
	if jsonErr != nil {
		return webhookResponse.Data, jsonErr
	}
	return webhookResponse.Data, nil
}

// List will retrieve all webhooks
func (s *WebhooksServiceOp) List(options ...interface{}) ([]Webhook, error) {
	webhookListResponse := ListWebhookResponse{}
	body, reqErr := s.client.DoRequest(http.MethodGet, "/v3/hooks", nil)
	if reqErr != nil {
		return webhookListResponse.Data, reqErr
	}
	jsonErr := json.Unmarshal(body, &webhookListResponse)
	if jsonErr != nil {
		return webhookListResponse.Data, jsonErr
	}
	return webhookListResponse.Data, nil
}

// Create will create a new webhook.
// The only fields required on a webhook are: Scope, Destination and IsActive
func (s *WebhooksServiceOp) Create(webhook Webhook, options ...interface{}) (Webhook, error) {
	var webhookResponse GetWebhookResponse
	jsonBody, err := json.Marshal(webhook)
	if err != nil {
		return webhookResponse.Data, err
	}
	reqBody := bytes.NewReader(jsonBody)
	body, reqErr := s.client.DoRequest(http.MethodPost, "/v3/hooks", reqBody)
	if reqErr != nil {
		return webhookResponse.Data, reqErr
	}

	jsonErr := json.Unmarshal(body, &webhookResponse)
	if jsonErr != nil {
		return webhookResponse.Data, jsonErr
	}

	return webhookResponse.Data, nil
}

// Update will update a single webhook.
func (s *WebhooksServiceOp) Update(webhook Webhook, options ...interface{}) (Webhook, error) {
	var webhookResponse GetWebhookResponse
	jsonBody, err := json.Marshal(webhook)
	if err != nil {
		return webhookResponse.Data, err
	}
	reqBody := bytes.NewReader(jsonBody)
	body, reqErr := s.client.DoRequest(http.MethodPut, fmt.Sprintf("/v3/hooks/%d", webhook.ID), reqBody)
	if reqErr != nil {
		return webhookResponse.Data, reqErr
	}

	jsonErr := json.Unmarshal(body, &webhookResponse)
	if jsonErr != nil {
		return webhookResponse.Data, jsonErr
	}

	return webhookResponse.Data, nil
}

// Delete will delete a webhook by the provided ID.
func (s *WebhooksServiceOp) Delete(id int64, options ...interface{}) error {
	_, reqErr := s.client.DoRequest(http.MethodDelete, fmt.Sprintf("/v3/hooks/%d", id), nil)
	if reqErr != nil {
		return reqErr
	}
	return nil
}
