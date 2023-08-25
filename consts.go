package alipanopen

const (
	ROOT_FOLDER_ID = "root" // 跟目录ID

	FILE_TYPE_FILE   = "file"   // 文件类型：文件
	FILE_TYPE_FOLDER = "folder" // 文件类型：文件夹

	CHECK_NAME_MODE_REFUSE = "refuse" // 重名检测策略：拒绝
)

const (
	ALIPAN_OPENAPI_HOST = "https://openapi.aliyundrive.com"

	HEADER_HOST          = "Host"
	HEADER_REFERER       = "Referer"
	HEADER_USER_AGENT    = "User-Agent"
	HEADER_RANGE         = "Range"
	HEADER_ACCEPT        = "Accept"
	HEADER_AUTHORIZATION = "Authorization"

	METHOD_GET  = "GET"
	METHOD_POST = "POST"

	API_OAUTH_USER_INFO        = "/oauth/users/info"
	API_GET_DRIVE_INFO         = "/adrive/v1.0/user/getDriveInfo"
	API_OAUTH_ACCESS_TOKEN     = "/oauth/access_token"
	API_OAUTH_AUTHORIZE_QRCODE = "/oauth/authorize/qrcode"
	API_FILE_LIST              = "/adrive/v1.0/openFile/list"
	API_FILE_CREATE            = "/adrive/v1.0/openFile/create"
	API_FILE_DELETE            = "/adrive/v1.0/openFile/delete"
	API_FILE_TRASH             = "/adrive/v1.0/openFile/recyclebin/trash"
	API_FILE_COMPLETE          = "/adrive/v1.0/openFile/complete"
	API_FILE_MOVE              = "/adrive/v1.0/openFile/move"
	API_FILE_UPDATE            = "/adrive/v1.0/openFile/update"
	API_FILE_GET_UPLOAD_URL    = "/adrive/v1.0/openFile/getUploadUrl"
	API_FILE_GET_DOWNLOAD_URL  = "/adrive/v1.0/openFile/getDownloadUrl"
)

// 二维码状态
const (
	QRCODE_STATUS_WAITLOGIN     = "WaitLogin"     // 等待扫码
	QRCODE_STATUS_SCANSUCCESS   = "ScanSuccess"   // 已扫描
	QRCODE_STATUS_LOGINSUCCESS  = "LoginSuccess"  // 已授权登录
	QRCODE_STATUS_QRCODEEXPIRED = "QRCodeExpired" // 已过期
)

// 权限信息
const (
	SCOPE_USER_BASE      = "user:base"      // 允许获取用户基础信息
	SCOPE_FILE_ALL_READ  = "file:all:read"  // 允许读用户文件
	SCOPE_FILE_ALL_WRITE = "file:all:write" // 允许写用户文件
)
