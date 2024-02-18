package main

import (
	"go-on-store/app/router"
)

func main() {

	r := router.SetupRouter()

	r.Run(":8080")
}
