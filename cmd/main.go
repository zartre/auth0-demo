package main

import "github.com/nzartre/auth0-demo/internal/router"

func main() {
	r := router.New()
	r.Run(":8080")
}
