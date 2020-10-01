package cli

// Консольный клиент

type CLI struct {
	IP   string
	Port string
}

//Client ...
type Client struct {
	NickName string
	Login    string
	Password string
}

func New(ip string, port string) *CLI {
	return &CLI{
		IP:   ip,
		Port: port,
	}
}
