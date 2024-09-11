package main

import "github.com/xws117/bluebell/routers"

type Product struct {
	Price int
	Name  string
}

func main() {
	server := routers.SetupRouter()
	server.Run(":8080")

}
