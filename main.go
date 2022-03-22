package main

import (
	"fmt"
	"github.com/tegarsubkhan236/go-fiber-project/bootstrap"
	"github.com/tegarsubkhan236/go-fiber-project/pkg/env"
	"log"
)

func main() {
	app := bootstrap.NewApplicationApp()
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", env.GetEnv("APP_HOST", "localhost"), env.GetEnv("APP_PORT", "4000"))))
}
