package NetCat

import (
	"net"
	"sync"
)

type Client struct {
	conn net.Conn
	name string
}

var (
	HistoryMessage []string
	clients        []Client
	ClientsNames   []string
	mutex          sync.Mutex
	WlcLogo        = []string{
		"          _nnnn_",
		"         dGGGGMMb",
		"        @p~qp~~qMb",
		"        M|@||@) M|",
		"        @,----.JM|",
		"       JS^\\__/  qKL",
		"      dZP        qKRb",
		"     dZP          qKKb",
		"    fZP            SMMb",
		"    HZM            MMMM",
		"    FqM            MMMM",
		"  __| \".        |\\dS\"qML",
		"  |    `.       | `' \\Zq",
		" _)      \\.___.,|     .'",
		" \\____   )MMMMMP|   .'",
		"      `-'       `--'",
	}

	colors = []string{
		"\033[31m", // Red
		"\033[32m", // Green
		"\033[33m", // Yellow
		"\033[34m", // Blue
		"\033[35m", // Magenta
		"\033[36m", // Cyan
		"\033[37m", // White
		"\033[90m", // Light Gray
		"\033[91m", // Light Red
		"\033[94m", // Light Blue
	}
)
