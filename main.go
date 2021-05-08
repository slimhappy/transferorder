package main

import (
	"transferorder/v1/router"
)

func main() {
	r := router.Router
	err := r.Run(":8000")
	if err != nil {
		panic(err)
	}
}