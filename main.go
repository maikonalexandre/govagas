package main

import (
	"fmt"
	"github.com/maikonalexandre/govagas/config"
	"github.com/maikonalexandre/govagas/router"
)

func main() {

	err := config.InitDB()

	if err != nil {
		fmt.Printf("opa")
		return
	}

	router.Initialize()
}
