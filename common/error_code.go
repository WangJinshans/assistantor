package common

const (
	Success    = "11000" // 成功
	Fail       = "11001" // 失败
	EmptyToken = "11002" // 空token
)

const (
	TokenKey = "tk"
)

// qrcode
const (
	InitStatus      = 1 // 创建
	ScanStatus      = 2 // 扫描
	UnConfirmStatus = 3 // 未确认
	CancelStatus    = 4 // 取消
	SuccessStatus   = 5 // 成功
	FailStatus      = 6 // 失败
	InValidStatus   = 7 // 过期
)
