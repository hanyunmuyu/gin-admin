package main

import (
	"gin-admin/router"
	"gin-admin/seeds"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		seeds.Run(os.Args)
	} else {
		router.Run()

	}
}
