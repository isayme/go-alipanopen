package alipanopen

import (
	"context"
	"fmt"
)

type GetQrCodeResp struct {
	QrCodeUrl string `json:"qrCodeUrl"`
	Sid       string `json:"sid"`
}

/**
 * 获取二维码
 */
func (client *Client) GetQrCode(ctx context.Context, clientId, clientSecret string) (*GetQrCodeResp, error) {
	reqBody := map[string]interface{}{
		"client_id":     clientId,
		"client_secret": clientSecret,
		"scopes":        []string{"user:base", "file:all:read", "file:all:write"},
	}
	respBody := &GetQrCodeResp{}
	_, err := client.request(METHOD_POST, API_OAUTH_AUTHORIZE_QRCODE, nil, reqBody, respBody)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

type GetQrCodeStatusResp struct {
	Status   string `json:"status"`
	AuthCode string `json:"authCode"`
}

/**
 * 获取二维码状态
 */
func (client *Client) GetQrCodeStatus(ctx context.Context, sid string) (*GetQrCodeStatusResp, error) {
	respBody := &GetQrCodeStatusResp{}
	_, err := client.request(METHOD_GET, fmt.Sprintf("/oauth/qrcode/%s/status", sid), nil, nil, respBody)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
