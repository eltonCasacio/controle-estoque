package main

import (
	"fmt"

	"github.com/eltonCasacio/controle-estoque/configs"
)

func main() {
	config, _ := configs.LoadConfig("./cmd/server/.env")
	fmt.Printf("?? %v", config)
}
