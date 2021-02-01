package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type WidgetTemplateService interface {
	Get(string, ...interface{}) (WidgetTemplate, error)
	List(...interface{}) (ListWidgetTemplateResponse, error)
	Create(WidgetTemplate, ...interface{}) (WidgetTemplate, error)
	Update(WidgetTemplate, ...interface{}) (WidgetTemplate, error)
	Delete(string, ...interface{}) error
}

type GetWidgetTemplateResponse struct {
	Data WidgetTemplate `json:"data"`
}

type ListWidgetTemplateResponse struct {
	Data []WidgetTemplate `json:"data"`
	Meta MetaResult       `json:"meta"`
}

// WidgetTemplate structure.
type WidgetTemplate struct {
	UUID               string        `json:"uuid,omitempty"`
	Name               string        `json:"name,omitempty"`
	Schema             []interface{} `json:"schema,omitempty"`
	Template           string        `json:"template,omitempty"`
	StorefrontAPIQuery string        `json:"storefront_api_query,omitempty"`
	Kind               string        `json:"kind,omitempty"`
	DateCreated        string        `json:"date_created,omitempty"`
	DateModified       string        `json:"date_modified,omitempty"`
	CurrentVersionUUID string        `json:"current_version_uuid,omitempty"`
	IconName           string        `json:"icon_name,omitempty"`
}

type WidgetTemplateServiceOp struct {
	client *Client
}

// Get will retrieve a widget template by the UUID
func (s *WidgetTemplateServiceOp) Get(uuid string, options ...interface{}) (WidgetTemplate, error) {
	widgetTemplateResponse := GetWidgetTemplateResponse{}
	body, err := s.client.DoRequest(http.MethodGet, fmt.Sprintf("/content/widget-templates/%s", uuid), nil)
	if err != nil {
		return widgetTemplateResponse.Data, err
	}

	jsonErr := json.Unmarshal(body, &widgetTemplateResponse)
	if jsonErr != nil {
		return widgetTemplateResponse.Data, jsonErr
	}

	return widgetTemplateResponse.Data, nil
}

// List will return a page of widget templates.
func (s *WidgetTemplateServiceOp) List(options ...interface{}) (ListWidgetTemplateResponse, error) {
	listResult := ListWidgetTemplateResponse{}
	body, err := s.client.DoRequest(http.MethodGet, "/content/widget-templates/", nil)
	if err != nil {
		return listResult, err
	}

	jsonErr := json.Unmarshal(body, &listResult)
	if jsonErr != nil {
		return listResult, jsonErr
	}
	return listResult, nil
}

// Create should do a thing.
func (s *WidgetTemplateServiceOp) Create(widgetTemplate WidgetTemplate, options ...interface{}) (WidgetTemplate, error) {
	widgetTemplateResponse := GetWidgetTemplateResponse{}
	jsonBody, err := json.Marshal(widgetTemplate)
	if err != nil {
		return widgetTemplateResponse.Data, err
	}
	reqBody := bytes.NewReader(jsonBody)
	body, reqErr := s.client.DoRequest(http.MethodPost, "/content/widget-templates/", reqBody)
	if reqErr != nil {
		return widgetTemplateResponse.Data, reqErr
	}

	jsonErr := json.Unmarshal(body, &widgetTemplateResponse)
	if jsonErr != nil {
		return widgetTemplateResponse.Data, jsonErr
	}
	return widgetTemplateResponse.Data, nil
}

// Update should do a thing.
func (s *WidgetTemplateServiceOp) Update(widgetTemplate WidgetTemplate, options ...interface{}) (WidgetTemplate, error) {
	widgetTemplateResponse := GetWidgetTemplateResponse{}
	jsonBody, err := json.Marshal(widgetTemplate)
	if err != nil {
		return widgetTemplateResponse.Data, err
	}
	reqBody := bytes.NewReader(jsonBody)
	body, reqErr := s.client.DoRequest(http.MethodPut, fmt.Sprintf("/content/widget-templates/%s", widgetTemplate.UUID), reqBody)
	if reqErr != nil {
		return widgetTemplateResponse.Data, reqErr
	}

	jsonErr := json.Unmarshal(body, &widgetTemplateResponse)
	if jsonErr != nil {
		return widgetTemplateResponse.Data, jsonErr
	}
	return widgetTemplateResponse.Data, nil
}

// Delete should do a thing.
func (s *WidgetTemplateServiceOp) Delete(uuid string, options ...interface{}) error {
	_, reqErr := s.client.DoRequest(http.MethodDelete, fmt.Sprintf("/content/widget-templates/%s", uuid), nil)
	if reqErr != nil {
		return reqErr
	}
	return nil
}
