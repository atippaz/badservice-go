package main

import (
	"bad-service-go/config"
	"bad-service-go/databases"
	"bad-service-go/server"
)

func main() {
	conf := config.ConfigGetting()
	db, err := databases.NewMongoDatabase(conf.Database)
	if err != nil {
		panic(err)
	}
	server := server.NewEchoServer(conf, db)
	server.Start()
}
