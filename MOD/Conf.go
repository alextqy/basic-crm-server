package mod

type Conf struct {
	TcpPort string `json:"tcp_port"`
	UdpPort string `json:"udp_port"`
	Lang    string `json:"lang"`
}
