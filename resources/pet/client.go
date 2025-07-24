package pet

import (
	bytes "bytes"
	json "encoding/json"
	io "io"
	http "net/http"
	os "os"
	sdkcore "pets_go/core"
	types "pets_go/types"
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

// Deletes a pet.
//
// Delete a pet.
//
// DELETE /pet/{petId}
func (c *Client) Delete(request DeleteRequest, reqModifiers ...RequestModifier) (http.Response, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/pet/" + sdkcore.FmtStringParam(request.PetId))
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

// Finds Pets by status.
//
// Multiple status values can be provided with comma separated strings.
//
// GET /pet/findByStatus
func (c *Client) FindByStatus(request FindByStatusRequest, reqModifiers ...RequestModifier) (http.Response, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/pet/" + "findByStatus")
	if err != nil {
		return http.Response{}, err
	}

	// Query params
	params := targetUrl.Query()
	sdkcore.AddQueryParam(params, "status", request.Status, "form", true)
	targetUrl.RawQuery = params.Encode()

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

// Find pet by ID.
//
// Returns a single pet.
//
// GET /pet/{petId}
func (c *Client) Get(request GetRequest, reqModifiers ...RequestModifier) (http.Response, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/pet/" + sdkcore.FmtStringParam(request.PetId))
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

// Add a new pet to the store.
//
// Add a new pet to the store.
//
// POST /pet
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (http.Response, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/pet")
	if err != nil {
		return http.Response{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.Pet{
		Category:  request.Category,
		Id:        request.Id,
		Status:    request.Status,
		Tags:      request.Tags,
		Name:      request.Name,
		PhotoUrls: request.PhotoUrls,
	})
	if err != nil {
		return http.Response{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return http.Response{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

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

// Uploads an image.
//
// Upload image of the pet.
//
// POST /pet/{petId}/uploadImage
func (c *Client) UploadImage(request UploadImageRequest, reqModifiers ...RequestModifier) (types.ApiResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/pet/" + sdkcore.FmtStringParam(request.PetId) + "/uploadImage")
	if err != nil {
		return types.ApiResponse{}, err
	}

	// Query params
	params := targetUrl.Query()
	sdkcore.AddQueryParam(params, "additionalMetadata", request.AdditionalMetadata, "form", true)
	targetUrl.RawQuery = params.Encode()

	// Prep body
	reqBodyBuf := &os.File{}
	reqBodyBuf = &request.Data

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.ApiResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/octet-stream")

	// Add auth
	err = c.coreClient.AddAuth(req, "api_key")
	if err != nil {
		return types.ApiResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.ApiResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.ApiResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.ApiResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.ApiResponse{}, err
	}
	var bodyData types.ApiResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.ApiResponse{}, err
	}
	return bodyData, nil

}

// Update an existing pet.
//
// Update an existing pet by Id.
//
// PUT /pet
func (c *Client) Update(request UpdateRequest, reqModifiers ...RequestModifier) (http.Response, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/pet")
	if err != nil {
		return http.Response{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.Pet{
		Category:  request.Category,
		Id:        request.Id,
		Status:    request.Status,
		Tags:      request.Tags,
		Name:      request.Name,
		PhotoUrls: request.PhotoUrls,
	})
	if err != nil {
		return http.Response{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("PUT", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return http.Response{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

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
