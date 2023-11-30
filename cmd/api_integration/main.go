package main

import (
	"github.com/freddiemo/healthcare-api/api_integration"
	"github.com/freddiemo/healthcare-api/config"
)

func main() {
	params := config.Init()
	api_integration.Init(params)
}
