package main

import (
	"github.com/freddiemo/healthcare-api/api"
	"github.com/freddiemo/healthcare-api/config"
	"github.com/freddiemo/healthcare-api/db"
)

func main() {
	params := config.Init()
	db := db.Init(params)
	api.Init(params, db)
}
