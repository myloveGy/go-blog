package profile

const (
	// 状态
	StatusActive  = 10 // 启用
	StatusDisable = 5  // 停用
)

var (
	StatusMapNames = map[uint8]string{
		StatusActive:  "启用",
		StatusDisable: "停用",
	}
)
