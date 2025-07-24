package order

import (
	json "encoding/json"
	io "io"
	http "net/http"
	sdkcore "pets_go/core"
	types "pets_go/types"
	strings "strings"
)

type Client struct {
	coreClient *sdkcore.CoreClient
}
type RequestModifier = func(req *http.Request) error

// Instantiate a new resource client
func NewClient(coreClient *sdkcore.CoreClient) *Client {
	client := Client{
		coreClient: coreClient,
	}

	return &client
}

// Delete purchase order by identifier.
//
// For valid response try integer IDs with value < 1000. Anything above 1000 or non-integers will generate API errors.
//
// DELETE /store/order/{orderId}
func (c *Client) Delete(request DeleteRequest, reqModifiers ...RequestModifier) (http.Response, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/store/" + "order/" + sdkcore.FmtStringParam(request.OrderId))
	if err != nil {
		return http.Response{}, err
	}

	// Init request
	req, err := http.NewRequest("DELETE", targetUrl.String(), nil)
	if err != nil {
		return http.Response{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")

	// Add auth
	err = c.coreClient.AddAuth(req, "api_key")
	if err != nil {
		return http.Response{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return http.Response{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return http.Response{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return http.Response{}, sdkcore.NewApiError(*req, *resp)
	}

	return *resp, nil

}

// Find purchase order by ID.
//
// For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions.
//
// GET /store/order/{orderId}
func (c *Client) Get(request GetRequest, reqModifiers ...RequestModifier) (http.Response, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/store/" + "order/" + sdkcore.FmtStringParam(request.OrderId))
	if err != nil {
		return http.Response{}, err
	}

	// Init request
	req, err := http.NewRequest("GET", targetUrl.String(), nil)
	if err != nil {
		return http.Response{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")

	// Add auth
	err = c.coreClient.AddAuth(req, "api_key")
	if err != nil {
		return http.Response{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return http.Response{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return http.Response{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return http.Response{}, sdkcore.NewApiError(*req, *resp)
	}

	return *resp, nil

}

// Place an order for a pet.
//
// Place a new order in the store.
//
// POST /store/order
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.Order, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/store/" + "order")
	if err != nil {
		return types.Order{}, err
	}

	// Prep body
	reqBodyBuf := &strings.Reader{}
	reqBodyBuf, err = sdkcore.FormUrlEncodedBody(
		types.Order{
			Complete: request.Complete,
			Id:       request.Id,
			PetId:    request.PetId,
			Quantity: request.Quantity,
			ShipDate: request.ShipDate,
			Status:   request.Status,
		},
		map[string]string{
			"complete": "form",
			"id":       "form",
			"petId":    "form",
			"quantity": "form",
			"shipDate": "form",
			"status":   "form",
		},
		map[string]bool{
			"complete": true,
			"id":       true,
			"petId":    true,
			"quantity": true,
			"shipDate": true,
			"status":   true,
		},
	)
	if err != nil {
		return types.Order{}, err
	}

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.Order{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Add auth
	err = c.coreClient.AddAuth(req, "api_key")
	if err != nil {
		return types.Order{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.Order{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.Order{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.Order{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.Order{}, err
	}
	var bodyData types.Order
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.Order{}, err
	}
	return bodyData, nil

}
