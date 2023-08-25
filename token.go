package alipanopen

import "context"

type RefreshTokenResp struct {
	AccessToken  string `json:"access_token"`  // 用户token
	RefreshToken string `json:"refresh_token"` // 刷新token
	ExpiresIn    int    `json:"expires_in"`    // 单位秒
}

// 使用 refresh_token 刷新获取新的用户 token
func (client *Client) RefreshToken(ctx context.Context, clientId, clientSecret string, refreshToken string) (*RefreshTokenResp, error) {
	reqBody := map[string]string{
		"client_id":     clientId,
		"client_secret": clientSecret,
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
	}
	respBody := &RefreshTokenResp{}
	_, err := client.request(METHOD_POST, API_OAUTH_ACCESS_TOKEN, nil, reqBody, respBody)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

// 使用 authCode 获取新的用户 token
func (client *Client) RefreshTokenByAuthCode(ctx context.Context, clientId, clientSecret string, authCode string) (*RefreshTokenResp, error) {
	reqBody := map[string]string{
		"client_id":     clientId,
		"client_secret": clientSecret,
		"grant_type":    "authorization_code",
		"code":          authCode,
	}

	respBody := &RefreshTokenResp{}
	_, err := client.request(METHOD_POST, API_OAUTH_ACCESS_TOKEN, nil, reqBody, respBody)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
