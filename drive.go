package alipanopen

import "context"

type GetDriveInfoResp struct {
	DefaultDriveId  string `json:"default_drive_id"`  // 默认空间ID
	ResourceDriveId string `json:"resource_drive_id"` // 资源库空间ID
	BackupDriveId   string `json:"backup_drive_id"`   // 备份盘空间ID
}

// 获取当前用户空间信息
func (client *Client) GetDriveInfo(ctx context.Context) (*GetDriveInfoResp, error) {
	reqBody := EmptyStruct{}

	respBody := GetDriveInfoResp{}
	_, err := client.requestWithAccessToken(METHOD_POST, API_GET_DRIVE_INFO, reqBody, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}
