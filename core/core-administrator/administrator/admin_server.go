package administrator

import (
	"github.com/jneubaum/honestvote/tests/logger"
)

func CreateServer(port string) {
	logger.Println("server.go", "main", "HTTP server running on port: "+port)
}
