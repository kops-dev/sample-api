package main

import (
	"fmt"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.GET("/hello", func(c *gofr.Context) (interface{}, error) {
		return fmt.Sprintf("Hello! Welcome to %s", c.GetAppName()), nil
	})

	app.Run()
}
