package main

import "gogolook/router"

func main() {
	r := router.SetupRouter()
	r.Run(":8000")
}
