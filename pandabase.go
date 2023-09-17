package pandabase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func NewClient(secret, version string) *Client {
	return &Client{
		Secret:  secret,
		Version: version,
	}
}

func (c *Client) makeRequest(method, path string, data interface{}) ([]byte, error) {
	url := fmt.Sprintf("https://api.pandabase.io/%s/%s", c.Version, path)
	reqBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.Secret)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		var baseError BaseError
		if err := json.Unmarshal(respBody, &baseError); err != nil {
			return nil, errors.New("HTTP error: " + resp.Status)
		}
		return nil, errors.New(baseError.Message)
	}

	return respBody, nil
}

type Pandabase struct {
	Client *Client
}

func NewPandabase(secret, version string) *Pandabase {
	return &Pandabase{
		Client: NewClient(secret, version),
	}
}

func (p *Pandabase) GetUser() (GetUserResponse, error) {
	var userResponse GetUserResponse
	endpoint := "user"
	responseBody, err := p.Client.makeRequest("GET", endpoint, nil)
	if err != nil {
		return GetUserResponse{}, err
	}

	if err := json.Unmarshal(responseBody, &userResponse); err != nil {
		return GetUserResponse{}, err
	}

	return userResponse, nil
}

func (p *Pandabase) GetCustomers() (GetCustomersResponse, error) {
	var customersResponse GetCustomersResponse
	endpoint := "customers"
	responseBody, err := p.Client.makeRequest("GET", endpoint, nil)
	if err != nil {
		return GetCustomersResponse{}, err
	}

	if err := json.Unmarshal(responseBody, &customersResponse); err != nil {
		return GetCustomersResponse{}, err
	}

	return customersResponse, nil
}

func (p *Pandabase) GetCustomer(id int) (GetCustomerResponse, error) {
	var customerResponse GetCustomerResponse
	endpoint := fmt.Sprintf("customer/%d", id)
	responseBody, err := p.Client.makeRequest("GET", endpoint, nil)
	if err != nil {
		return GetCustomerResponse{}, err
	}

	if err := json.Unmarshal(responseBody, &customerResponse); err != nil {
		return GetCustomerResponse{}, err
	}

	return customerResponse, nil
}

func (p *Pandabase) CreateCustomer(data CreateCustomerRequest) (GetCustomerResponse, error) {
	var customerResponse GetCustomerResponse
	endpoint := "customer"
	responseBody, err := p.Client.makeRequest("POST", endpoint, data)
	if err != nil {
		return GetCustomerResponse{}, err
	}

	if err := json.Unmarshal(responseBody, &customerResponse); err != nil {
		return GetCustomerResponse{}, err
	}

	return customerResponse, nil
}

func (p *Pandabase) UpdateCustomer(data UpdateCustomerRequest) (GetCustomerResponse, error) {
	var customerResponse GetCustomerResponse
	endpoint := fmt.Sprintf("customer/%d", data.ID)
	responseBody, err := p.Client.makeRequest("PUT", endpoint, data)
	if err != nil {
		return GetCustomerResponse{}, err
	}

	if err := json.Unmarshal(responseBody, &customerResponse); err != nil {
		return GetCustomerResponse{}, err
	}

	return customerResponse, nil
}

func (p *Pandabase) DeleteCustomer(id int) (GetCustomerResponse, error) {
	var customerResponse GetCustomerResponse
	endpoint := fmt.Sprintf("customer/%d", id)
	responseBody, err := p.Client.makeRequest("DELETE", endpoint, nil)
	if err != nil {
		return GetCustomerResponse{}, err
	}

	if err := json.Unmarshal(responseBody, &customerResponse); err != nil {
		return GetCustomerResponse{}, err
	}

	return customerResponse, nil
}

func (p *Pandabase) GetShop() (GetShopResponse, error) {
	var shopResponse GetShopResponse
	endpoint := "shop"
	responseBody, err := p.Client.makeRequest("GET", endpoint, nil)
	if err != nil {
		return GetShopResponse{}, err
	}

	if err := json.Unmarshal(responseBody, &shopResponse); err != nil {
		return GetShopResponse{}, err
	}

	return shopResponse, nil
}

func (p *Pandabase) GetProducts() (GetProductResponse, error) {
	var productsResponse GetProductResponse
	endpoint := "products"
	responseBody, err := p.Client.makeRequest("GET", endpoint, nil)
	if err != nil {
		return GetProductResponse{}, err
	}

	if err := json.Unmarshal(responseBody, &productsResponse); err != nil {
		return GetProductResponse{}, err
	}

	return productsResponse, nil
}

func (p *Pandabase) GetProduct(id int) (GetProductResponse, error) {
	var productResponse GetProductResponse
	endpoint := fmt.Sprintf("product/%d", id)
	responseBody, err := p.Client.makeRequest("GET", endpoint, nil)
	if err != nil {
		return GetProductResponse{}, err
	}

	if err := json.Unmarshal(responseBody, &productResponse); err != nil {
		return GetProductResponse{}, err
	}

	return productResponse, nil
}

func (p *Pandabase) CreateProduct(data CreateProductRequest) (GetProductResponse, error) {
	var productResponse GetProductResponse
	endpoint := "product"
	responseBody, err := p.Client.makeRequest("POST", endpoint, data)
	if err != nil {
		return GetProductResponse{}, err
	}

	if err := json.Unmarshal(responseBody, &productResponse); err != nil {
		return GetProductResponse{}, err
	}

	return productResponse, nil
}

func (p *Pandabase) UpdateProduct(data UpdateProductRequest) (GetProductResponse, error) {
	var productResponse GetProductResponse
	endpoint := fmt.Sprintf("product/%d", data.ID)
	responseBody, err := p.Client.makeRequest("PUT", endpoint, data)
	if err != nil {
		return GetProductResponse{}, err
	}

	if err := json.Unmarshal(responseBody, &productResponse); err != nil {
		return GetProductResponse{}, err
	}

	return productResponse, nil
}

func (p *Pandabase) DeleteProduct(id int) (GetProductResponse, error) {
	var productResponse GetProductResponse
	endpoint := fmt.Sprintf("product/%d", id)
	responseBody, err := p.Client.makeRequest("DELETE", endpoint, nil)
	if err != nil {
		return GetProductResponse{}, err
	}

	if err := json.Unmarshal(responseBody, &productResponse); err != nil {
		return GetProductResponse{}, err
	}

	return productResponse, nil
}

func (p *Pandabase) CreateTransaction(data CreateTransactionRequest) (GetCreatedTransactionResponse, error) {
	var transactionResponse GetCreatedTransactionResponse
	endpoint := "transaction"
	responseBody, err := p.Client.makeRequest("POST", endpoint, data)
	if err != nil {
		return GetCreatedTransactionResponse{}, err
	}

	if err := json.Unmarshal(responseBody, &transactionResponse); err != nil {
		return GetCreatedTransactionResponse{}, err
	}

	return transactionResponse, nil
}

func (p *Pandabase) GetTransaction(id int) (GetTransactionResponse, error) {
	var transactionResponse GetTransactionResponse
	endpoint := fmt.Sprintf("transaction/%d", id)
	responseBody, err := p.Client.makeRequest("GET", endpoint, nil)
	if err != nil {
		return GetTransactionResponse{}, err
	}

	if err := json.Unmarshal(responseBody, &transactionResponse); err != nil {
		return GetTransactionResponse{}, err
	}

	return transactionResponse, nil
}
