package alipanopen

import (
	"context"
	"time"
)

// 文件模型
type File struct {
	FileName     string    `json:"name"`           // 名称
	FileSize     int64     `json:"size"`           // 大小
	UpdatedAt    time.Time `json:"updated_at"`     // 更新时间
	ContentHash  string    `json:"content_hash"`   // 文件内容sha1
	Type         string    `json:"type"`           // 文件类型
	DriveId      string    `json:"drive_id"`       // 所属空间ID
	FileId       string    `json:"file_id"`        // 文件ID
	ParentFileId string    `json:"parent_file_id"` // 父文件夹ID
}

type GetFileReq struct {
	DriveId string `json:"drive_id"`
	FileId  string `json:"file_id"`
}

func (client *Client) GetFile(ctx context.Context, reqBody *GetFileReq) (*File, error) {
	respBody := File{}
	err := client.requestWithAccessToken(METHOD_POST, API_FILE_GET, reqBody, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}

type GetFileByPathReq struct {
	DriveId  string `json:"drive_id"`
	FilePath string `json:"file_path"`
}

func (client *Client) GetFileByPath(ctx context.Context, reqBody *GetFileByPathReq) (*File, error) {
	respBody := File{}
	err := client.requestWithAccessToken(METHOD_POST, API_FILE_GET_BY_PATH, reqBody, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}

type ListFileReq struct {
	DriveId      string `json:"drive_id"`
	ParentFileId string `json:"parent_file_id"`

	Limit  int    `json:"limit"`  // 单页数量
	Marker string `json:"marker"` // 分页用
}

type ListFileResp struct {
	Items      []*File `json:"items"`
	NextMarker string  `json:"next_marker"` // 分页用
}

// 列举文件夹下文件
func (client *Client) ListFolder(ctx context.Context, reqBody *ListFileReq) (*ListFileResp, error) {
	respBody := ListFileResp{}
	err := client.requestWithAccessToken(METHOD_POST, API_FILE_LIST, reqBody, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}

type GetFileDownloadUrlReq struct {
	DriveId string `json:"drive_id"`
	FileId  string `json:"file_id"`
}

type GetFileDownloadUrlResp struct {
	Url        string    `json:"url"`
	Expiration time.Time `json:"expiration"`
}

// 获取文件下载地址
func (client *Client) GetDownloadUrl(ctx context.Context, reqBody *GetFileDownloadUrlReq) (*GetFileDownloadUrlResp, error) {
	respBody := GetFileDownloadUrlResp{}
	err := client.requestWithAccessToken(METHOD_POST, API_FILE_GET_DOWNLOAD_URL, reqBody, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}

type CreateFolderReq struct {
	Name          string `json:"name"`            // 文件夹名称
	DriveId       string `json:"drive_id"`        // 所属空间ID
	ParentFileId  string `json:"parent_file_id"`  // 父文件夹ID
	Type          string `json:"type"`            // 固定是 folder
	CheckNameMode string `json:"check_name_mode"` // 重名检测策略
}

// 创建文件夹
func (client *Client) CreateFolder(ctx context.Context, reqBody *CreateFolderReq) (*File, error) {
	reqBody.Type = FILE_TYPE_FOLDER

	fi := &File{}
	err := client.requestWithAccessToken(METHOD_POST, API_FILE_CREATE, reqBody, fi)
	if err != nil {
		return nil, err
	}

	return fi, nil
}

type CreateFileReq struct {
	Name            string `json:"name"`           // 文件名
	DriveId         string `json:"drive_id"`       // 文件所属空间ID
	ParentFileId    string `json:"parent_file_id"` // 文件ID
	Type            string `json:"type"`           // 文件类型
	ContentHash     string `json:"content_hash"`   // 文件内容sha1
	ContentHashName string `json:"content_hash_name"`
	CheckNameMode   string `json:"check_name_mode"` // 重名检测策略
	Size            int64  `json:"size"`            // 文件大小
}

// 创建文件
func (client *Client) CreateFile(ctx context.Context, reqBody *CreateFileReq) (*CreateFileResp, error) {
	reqBody.Type = FILE_TYPE_FILE

	respBody := CreateFileResp{}
	err := client.requestWithAccessToken(METHOD_POST, API_FILE_CREATE, reqBody, &respBody)
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

type DeleteFileReq struct {
	DriveId string `json:"drive_id"`
	FileId  string `json:"file_id"`
}

// 删除文件
func (client *Client) DeleteFile(ctx context.Context, reqBody *DeleteFileReq) error {
	respBody := EmptyStruct{}
	err := client.requestWithAccessToken(METHOD_POST, API_FILE_DELETE, reqBody, &respBody)
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

// 完成文件创建
func (client *Client) CompleteFile(ctx context.Context, reqBody *CompleteFileReq) (*CompleteFileResp, error) {
	respBody := &CompleteFileResp{}
	err := client.requestWithAccessToken(METHOD_POST, API_FILE_COMPLETE, reqBody, respBody)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

type TrashFileReq struct {
	DriveId string `json:"drive_id"`
	FileId  string `json:"file_id"`
}

// 将文件移入回收站
func (client *Client) TrashFile(ctx context.Context, reqBody *TrashFileReq) error {
	respBody := EmptyStruct{}
	err := client.requestWithAccessToken(METHOD_POST, API_FILE_TRASH, reqBody, &respBody)
	return err
}

type MoveFileReq struct {
	DriveId string `json:"drive_id"`
	FileId  string `json:"file_id"`
	NewName string `json:"new_name"` // 新文件名

	CheckNameMode  string `json:"check_name_mode"`
	Overwrite      bool   `json:"overwrite"`
	ToParentFileId string `json:"to_parent_file_id"` // 目的文件夹ID
}

// 移动文件
func (client *Client) MoveFile(ctx context.Context, reqBody *MoveFileReq) error {
	respBody := EmptyStruct{}
	err := client.requestWithAccessToken(METHOD_POST, API_FILE_MOVE, reqBody, &respBody)
	return err
}

type UpdateFileNameReq struct {
	DriveId       string `json:"drive_id"`
	FileId        string `json:"file_id"`
	Name          string `json:"name"`
	CheckNameMode string `json:"check_name_mode"`
}

// 更新文件名
func (client *Client) UpdateFileName(ctx context.Context, reqBody *UpdateFileNameReq) error {
	respBody := EmptyStruct{}
	err := client.requestWithAccessToken(METHOD_POST, API_FILE_UPDATE, reqBody, &respBody)
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

// 获取文件上传地址
func (client *Client) GetUploadUrl(ctx context.Context, reqBody *GetUploadUrlReq) (*GetUploadUrlResp, error) {
	respBody := &GetUploadUrlResp{}
	err := client.requestWithAccessToken(METHOD_POST, API_FILE_GET_UPLOAD_URL, reqBody, respBody)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
