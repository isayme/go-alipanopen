package goalipanopen

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	host        string
	accessToken string
	c           *resty.Client
}

func NewClient() *Client {
	return &Client{
		host: ALIPAN_OPENAPI_HOST,
		c:    resty.New(),
	}
}

type EmptyStruct struct{}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (client *Client) SetRestyClient(restyClient *resty.Client) {
	client.c = restyClient
}

func (client *Client) SetHost(host string) {
	client.host = host
}

func (client *Client) SetAccessToken(accessToken string) {
	client.accessToken = accessToken
}

func (client *Client) GetAccessToken() string {
	return client.accessToken
}

func (client *Client) requestWithAccessToken(method, path string, body, out interface{}) (errResp *ErrorResponse, err error) {
	accessToken := client.GetAccessToken()

	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("Bearer %s", accessToken)

	return client.request(method, path, headers, body, out)
}

func (client *Client) request(method, path string, headers map[string]string, body, out interface{}) (errResp *ErrorResponse, err error) {
	url := fmt.Sprintf("%s%s", ALIPAN_OPENAPI_HOST, path)

	req := client.c.R()
	if headers != nil {
		req = req.SetHeaders(headers)
	}

	resp, err := req.SetDoNotParseResponse(true).SetBody(body).Execute(method, url)
	if err != nil {
		return nil, err
	}

	defer resp.RawBody().Close()

	bs, err := io.ReadAll(resp.RawBody())
	if err != nil {
		return nil, err
	}

	if resp.IsSuccess() {
		json.Unmarshal(bs, out)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	statusCode := resp.StatusCode()

	errResp = &ErrorResponse{}
	err = json.Unmarshal(bs, errResp)
	if err != nil {
		return nil, err
	}

	return errResp, fmt.Errorf("requestFail, code: %s/%d", errResp.Code, statusCode)
}
