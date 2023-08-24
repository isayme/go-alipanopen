package alipanopen

import "context"

type GetDriveInfoResp struct {
	DefaultDriveId  string `json:"default_drive_id"`
	ResourceDriveId string `json:"resource_drive_id"`
	BackupDriveId   string `json:"backup_drive_id"`
}

func (client *Client) GetDriveInfo(ctx context.Context) (*GetDriveInfoResp, error) {
	reqBody := EmptyStruct{}

	respBody := GetDriveInfoResp{}
	_, err := client.requestWithAccessToken(METHOD_POST, API_GET_DRIVE_INFO, reqBody, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}
