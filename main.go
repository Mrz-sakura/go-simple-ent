package main

import (
	"dwbackend/config"
	"dwbackend/db"
	"dwbackend/global"
	"dwbackend/router"
	"fmt"
)

func main() {
	config.InitGlobalConfig()
	db.InitDB()

	r := router.InitRouter()

	r.Run(fmt.Sprintf(":%d", global.GlobalConfigInstance.Http.Port))
}
