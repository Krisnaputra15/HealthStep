package main

import (
	"github.com/Krisnaputra15/gsc-solution/config"
	"github.com/Krisnaputra15/gsc-solution/db"
)

func init() {
	config.LoadEnv()
	db.ConnectToDB()
}

func main() {
	db.DB.AutoMigrate()
}