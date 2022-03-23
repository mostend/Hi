package client

type Brook struct {
	Server     string `json:"server"`
	ServerType string `json:"serverType"`
	Addr       string `json:"addr"`
	Password   string `json:"password"`
	Http       string `json:"http"`
	Socks5     string `json:"socks5"`
	NoBrook    string `json:"noBrook"`
}
