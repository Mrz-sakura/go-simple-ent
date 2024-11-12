package db

import (
	"dwbackend/ent"
	"dwbackend/global"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() {
	if global.DB != nil {
		return
	}

	client, err := ent.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai", global.GlobalConfigInstance.MySQL.User, global.GlobalConfigInstance.MySQL.Password, global.GlobalConfigInstance.MySQL.Host, global.GlobalConfigInstance.MySQL.Port, global.GlobalConfigInstance.MySQL.DB))
	if err != nil {
		panic(err)
	}

	fmt.Println("数据库连接成功～")

	global.DB = client
}
