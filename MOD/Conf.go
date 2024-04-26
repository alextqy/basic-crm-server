package mod

type ConfMod struct {
	DbHost  string `json:"db_host"`  // 数据库地址
	DbPort  string `json:"db_port"`  // 数据库接口
	DbUser  string `json:"db_user"`  // 数据库账号
	DbPwd   string `json:"db_pwd"`   // 数据库密码
	DbDebug bool   `json:"db_debug"` // 调试状态
	TcpPort string `json:"tcp_port"` // api端口
	UdpPort string `json:"udp_port"` // 广播端口
	Lang    string `json:"lang"`     // 系统语言
	Reg     string `json:"reg"`      // api正则过滤
	EncKey  string `json:"enc_key"`  // Token密钥
}
