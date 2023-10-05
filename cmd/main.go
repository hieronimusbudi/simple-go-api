package main

import (
	"fmt"

	"github.com/hieronimusbudi/simple-go-api/internal/adapter"
	"github.com/hieronimusbudi/simple-go-api/internal/config"
)

//	@title			Gathering App API
//	@version		v1.0.0
//	@description	# Introduction
//	@description	This is documentation for Gathering App API

func main() {
	config.SetConfig(".")
	r := adapter.Router()
	r.Run(fmt.Sprintf(":%s", config.Get().PORT))
}
