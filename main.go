package main

import (
	"fmt"

	"github.com/kidboy-man/mini-bank-rest/configs"
	"github.com/kidboy-man/mini-bank-rest/databases"
	"github.com/kidboy-man/mini-bank-rest/routers"
)

func main() {
	err := configs.AppConfig.Setup()
	if err != nil {
		panic(err)
	}

	err = databases.Init()
	if err != nil {
		panic(err)
	}

	server := routers.Setup()
	server.Run(fmt.Sprintf(":%s", configs.AppConfig.AppPort))
}
