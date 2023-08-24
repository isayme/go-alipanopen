package goalipanopen

import (
	"context"
	"time"
)

type File struct {
	FileName     string    `json:"name"`
	FileSize     int64     `json:"size"`
	UpdatedAt    time.Time `json:"updated_at"`
	ContentHash  string    `json:"content_hash"`
	Type         string    `json:"type"`
	DriveId      string    `json:"drive_id"`
	FileId       string    `json:"file_id"`
	ParentFileId string    `json:"parent_file_id"`
}

type ListFileReq struct {
	DriveId      string `json:"drive_id"`
	ParentFileId string `json:"parent_file_id"`

	Limit  int    `json:"limit"`
	Marker string `json:"marker"`
}

type ListFileResp struct {
	Items      []*File `json:"items"`
	NextMarker string  `json:"next_marker"`
}

func (client *Client) ListFolder(ctx context.Context, reqBody *ListFileReq) ([]*File, error) {
	respBody := ListFileResp{}
	_, err := client.requestWithAccessToken(METHOD_POST, API_FILE_LIST, reqBody, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Items, nil
}

type GetFileDownloadUrlResp struct {
	Url        string    `json:"url"`
	Expiration time.Time `json:"expiration"`
}

func (client *Client) GetDownloadUrl(ctx context.Context, driveId, fileId string) (*GetFileDownloadUrlResp, error) {
	reqBody := map[string]string{
		"drive_id": driveId,
		"file_id":  fileId,
	}

	respBody := GetFileDownloadUrlResp{}
	_, err := client.requestWithAccessToken(METHOD_POST, API_FILE_GET_DOWNLOAD_URL, reqBody, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}

type CreateFolderReq struct {
	Name          string `json:"name"`
	DriveId       string `json:"drive_id"`
	ParentFileId  string `json:"parent_file_id"`
	Type          string `json:"type"`
	CheckNameMode string `json:"check_name_mode"`
}

func (client *Client) CreateFolder(ctx context.Context, reqBody *CreateFolderReq) (*File, error) {
	fi := &File{}
	_, err := client.requestWithAccessToken(METHOD_POST, API_FILE_CREATE, reqBody, fi)
	if err != nil {
		return nil, err
	}

	return fi, nil
}

type CreateFileReq struct {
	Name            string `json:"name"`
	DriveId         string `json:"drive_id"`
	ParentFileId    string `json:"parent_file_id"`
	Type            string `json:"type"`
	ContentHash     string `json:"content_hash"`
	ContentHashName string `json:"content_hash_name"`
	CheckNameMode   string `json:"check_name_mode"`
	Size            int64  `json:"size"`
}

func (client *Client) CreateFile(ctx context.Context, reqBody *CreateFileReq) (*CreateFileResp, error) {
	respBody := CreateFileResp{}
	_, err := client.requestWithAccessToken(METHOD_POST, API_FILE_CREATE, reqBody, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}

type UploadPartInfo struct {
	PartNumber int    `json:"part_number"`
	PartSize   int64  `json:"part_size"`
	UploadUrl  string `json:"upload_url"`
}

type CreateFileResp struct {
	DriveId      string           `json:"drive_id"`
	FileId       string           `json:"file_id"`
	UploadId     string           `json:"upload_id"`
	PartInfoList []UploadPartInfo `json:"part_info_list"`
}

func (client *Client) DeleteFile(ctx context.Context, driveId, fileId string) error {
	reqBody := map[string]string{
		"drive_id": driveId,
		"file_id":  fileId,
	}
	respBody := EmptyStruct{}
	_, err := client.requestWithAccessToken(METHOD_POST, API_FILE_DELETE, reqBody, &respBody)
	if err != nil {
		return err
	}

	return nil
}

type CompleteFileReq struct {
	DriveId  string `json:"drive_id"`
	FileId   string `json:"file_id"`
	UploadId string `json:"upload_id"`
}

type CompleteFileResp struct {
	ContentHash string `json:"content_hash"`
	Size        int64  `json:"size"`
}

func (client *Client) CompleteFile(ctx context.Context, reqBody *CompleteFileReq) (*CompleteFileResp, error) {
	respBody := &CompleteFileResp{}
	_, err := client.requestWithAccessToken(METHOD_POST, API_FILE_COMPLETE, reqBody, respBody)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (client *Client) TrashFile(ctx context.Context, driveId, fileId string) error {
	reqBody := map[string]string{
		"drive_id": driveId,
		"file_id":  fileId,
	}

	respBody := EmptyStruct{}
	_, err := client.requestWithAccessToken(METHOD_POST, API_FILE_TRASH, reqBody, &respBody)
	return err
}

type MoveFileReq struct {
	DriveId string `json:"drive_id"`
	FileId  string `json:"file_id"`
	NewName string `json:"new_name"`

	CheckNameMode  string `json:"check_name_mode"`
	Overwrite      bool   `json:"overwrite"`
	ToParentFileId string `json:"to_parent_file_id"`
}

func (client *Client) MoveFile(ctx context.Context, reqBody *MoveFileReq) error {
	respBody := EmptyStruct{}
	_, err := client.requestWithAccessToken(METHOD_POST, API_FILE_MOVE, reqBody, &respBody)
	return err
}

type UpdateFileNameReq struct {
	DriveId       string `json:"drive_id"`
	FileId        string `json:"file_id"`
	Name          string `json:"name"`
	CheckNameMode string `json:"check_name_mode"`
}

func (client *Client) UpdateFileName(ctx context.Context, reqBody *UpdateFileNameReq) error {
	respBody := EmptyStruct{}
	_, err := client.requestWithAccessToken(METHOD_POST, API_FILE_UPDATE, reqBody, &respBody)
	return err
}

type GetUploadPartInfoReq struct {
	PartNumber int `json:"part_number"`
}
type GetUploadUrlReq struct {
	DriveId      string                 `json:"drive_id"`
	FileId       string                 `json:"file_id"`
	UploadId     string                 `json:"upload_id"`
	PartInfoList []GetUploadPartInfoReq `json:"part_info_list"`
}

type GetUploadUrlResp struct {
	DriveId      string           `json:"drive_id"`
	FileId       string           `json:"file_id"`
	UploadId     string           `json:"upload_id"`
	PartInfoList []UploadPartInfo `json:"part_info_list"`
}

func (client *Client) GetUploadUrl(ctx context.Context, reqBody *GetUploadUrlResp) (*GetUploadUrlResp, error) {
	respBody := &GetUploadUrlResp{}
	_, err := client.requestWithAccessToken(METHOD_POST, API_FILE_GET_UPLOAD_URL, reqBody, respBody)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
