package vtu_ng

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	httpclient "github.com/CeoFred/vtu_ng/lib"
)

type VTUConnection struct {
	Username        string
	Password        string
	client          HttpClient
	ServerURL       string
	authCredentials string
}

const BaseURL = "https://vtu.ng/wp-json/api/v1/"

type BaseResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	BaseResponse
	Data interface{} `json:"data"`
}

func (e ErrorResponse) Error() string {
	return e.BaseResponse.Message
}

type BalanceResponse struct {
	BaseResponse
	Data struct {
		Balance  string `json:"balance"`
		Currency string `json:"currency"`
	} `json:"data"`
}

type AirtimeTopUpResponse struct {
	BaseResponse
	Data struct {
		Network string `json:"network"`
		Phone   string `json:"phone"`
		Amount  string `json:"amount"`
		OrderID int    `json:"order_id"`
	} `json:"data"`
}

func NewVTUConnection(username, password string) *VTUConnection {
	return &VTUConnection{
		Username:        username,
		Password:        password,
		client:          httpclient.NewAPIClient(BaseURL, ""),
		authCredentials: fmt.Sprintf("?username=%s&password=%s", username, password),
	}
}

func (c *VTUConnection) GetBalance(ctx context.Context) (*BalanceResponse, error) {

	resp, err := c.client.Get(ctx, "balance"+c.authCredentials)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}

		return nil, errorResponse
	}

	var balance BalanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&balance); err != nil {
		return nil, err
	}

	return &balance, nil

}

func (c *VTUConnection) AirtimeTopUP(ctx context.Context, phone string, network_id NetworkID, amount float64) (*AirtimeTopUpResponse, error) {

	path := fmt.Sprintf("airtime%s&phone=%s&network_id=%s&amount=%f", c.authCredentials, phone, network_id, amount)

	resp, err := c.client.Get(ctx, path)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println(resp.StatusCode, "<= status code")

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}

		return nil, errorResponse
	}

	var response AirtimeTopUpResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("error decoding response", err)
		return nil, err
	}

	return &response, nil
}

func (c *VTUConnection) BuyData(ctx context.Context, phone string, network_id NetworkID, variation_id string) (*AirtimeTopUpResponse, error) {

	path := fmt.Sprintf("data%s&phone=%s&network_id=%s&variation_id=%s", c.authCredentials, phone, network_id, variation_id)

	resp, err := c.client.Get(ctx, path)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}

		return nil, errorResponse
	}

	var response AirtimeTopUpResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("error decoding response", err)
		return nil, err
	}

	return &response, nil
}
