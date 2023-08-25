package alipanopen

import "context"

type User struct {
	Id   string `json:"id"`   // 用户ID
	Name string `json:"name"` // 用户昵称
}

/**
 * 获取当前用户信息
 */
func (client *Client) GetCurrentUser(ctx context.Context) (*User, error) {
	reqBody := EmptyStruct{}

	respBody := User{}
	_, err := client.requestWithAccessToken(METHOD_GET, API_OAUTH_USER_INFO, reqBody, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}
