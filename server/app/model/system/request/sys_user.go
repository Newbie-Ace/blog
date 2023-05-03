package request

// Register User register structure
type Register struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	NickName string `json:"nickname"` // 昵称
}

// Login structure
type Login struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}
