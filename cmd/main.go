package main

import (
	"database/sql"
	"fmt"

	"github.com/ryancarlos88/golang-base/config"
	"github.com/ryancarlos88/golang-base/internal/app"
	"github.com/ryancarlos88/golang-base/internal/infra/web/webserver"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	cfg, err := config.LoadConfig(".")
	if err != nil{
		panic(err)
	}
	db, err := sql.Open(cfg.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))
	if err != nil {
		panic(err)
	}
	ws := webserver.NewWebServer(cfg.WebServerPort)
	app := app.NewApp(ws, db)
	app.Run()
}