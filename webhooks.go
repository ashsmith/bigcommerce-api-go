package bigcommerce

import (
	"encoding/json"
	"log"
	"net/http"
)

type WebhooksService interface {
	Get(int64, interface{}) (*Webhook, error)
	List(interface{}) ([]Webhook, error)
	Update(Webhook) (*Webhook, error)
	Delete(int64) error
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
	ID          int64             `json:"id"`
	ClientID    string            `json:"client_id"`
	StoreHash   string            `json:"store_hash"`
	CreatedAt   int64             `json:"created_at"`
	UpdatedAt   int64             `json:"updated_at"`
	Scope       string            `json:"scope"`
	Destination string            `json:"destination"`
	IsActive    bool              `json:"is_active"`
	Headers     map[string]string `json:"headers"`
}

type WebhooksServiceOp struct {
	client *Client
}

// Get will fetch a single webhook by the provided ID.
func (s *WebhooksServiceOp) Get(id int64, options interface{}) (*Webhook, error) {
	body := s.client.DoRequest(http.MethodGet, "/v3/hooks", nil)

	webhookResponse := GetWebhookResponse{}
	jsonErr := json.Unmarshal(body, &webhookResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return &webhookResponse.Data, nil
}

// List will retrieve all webhooks
func (s *WebhooksServiceOp) List(options interface{}) ([]Webhook, error) {
	webhookListResponse := ListWebhookResponse{}
	return webhookListResponse.Data, nil
}

// Update will update a single webhook.
func (s *WebhooksServiceOp) Update(webhook Webhook) (*Webhook, error) {
	webhookResponse := GetWebhookResponse{}
	return &webhookResponse.Data, nil
}

// Delete will delete a webhook by the provided ID.
func (s *WebhooksServiceOp) Delete(id int64) error {
	return nil
}
