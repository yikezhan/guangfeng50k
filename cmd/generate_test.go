package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"testing"
)

/**
gen生成代码
*/
func Test_GenCode(t *testing.T) {
	g := gen.NewGenerator(gen.Config{OutPath: "../internal/model"})
	db, _ := gorm.Open(mysql.Open("root:220706Aa!@(127.0.0.1:3306)/guangfeng50k?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db)
	g.GenerateModelAs("room_tab", "RoomTab")
	g.GenerateModelAs("game_result_tab", "GameResultTab")
	g.Execute()
}
