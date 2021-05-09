package main

import (
	"code.lhc.org/we-blog/config"
	"code.lhc.org/we-blog/model"
	"code.lhc.org/we-blog/route"
)

func main() {
	r := route.InitRouter()

	config.InitConfig()

	model.InitDb()
	defer model.Db.Close()

	r.Run("127.0.0.1:8080")
}
