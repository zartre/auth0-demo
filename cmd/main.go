package main

import (
	"github.com/joho/godotenv"
	"github.com/nzartre/auth0-demo/internal/authen"
	"github.com/nzartre/auth0-demo/internal/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	auth, err := authen.New()
	if err != nil {
		panic(err)
	}

	r := router.New(auth)
	r.Run(":8080")
}
