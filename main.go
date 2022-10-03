package main

import "rest-api-gin/routers"

func main() {
	var PORT = ":8081"

	routers.StartServer().Run(PORT)
}
