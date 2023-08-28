package alipanopen

import "context"

type RefreshTokenReq struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
	Code         string `json:"code"`
}

type RefreshTokenResp struct {
	AccessToken  string `json:"access_token"`  // 用户token
	RefreshToken string `json:"refresh_token"` // 刷新token
	ExpiresIn    int    `json:"expires_in"`    // 单位秒
}

// 刷新获取新的用户 token
func (client *Client) RefreshToken(ctx context.Context, reqBody *RefreshTokenReq) (*RefreshTokenResp, error) {
	reqBody.GrantType = "refresh_token"
	respBody := &RefreshTokenResp{}
	_, err := client.request(METHOD_POST, API_OAUTH_ACCESS_TOKEN, nil, reqBody, respBody)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
