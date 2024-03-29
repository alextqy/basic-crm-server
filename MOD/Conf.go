package mod

type Conf struct {
	DbHost  string `json:"db_host"`
	DbPort  string `json:"db_port"`
	DbUser  string `json:"db_user"`
	DbPwd   string `json:"db_pwd"`
	DbDebug bool   `json:"db_debug"`
	TcpPort string `json:"tcp_port"`
	UdpPort string `json:"udp_port"`
	Lang    string `json:"lang"`
	Reg     string `json:"reg"`
}
