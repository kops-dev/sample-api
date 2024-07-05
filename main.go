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

	app.GET("/error", func(c *gofr.Context) (interface{}, error) {
		return nil, fmt.Errorf("sample error")
	})

	app.GET("/warn", func(c *gofr.Context) (interface{}, error) {
		c.Logger.Warn("sample warn log")

		return nil, nil
	})

	app.GET("/fatal", func(c *gofr.Context) (interface{}, error) {
		c.Logger.Fatal("sample warn log")

		return nil, nil
	})

	app.GET("/debug", func(c *gofr.Context) (interface{}, error) {
		c.Logger.Debug("sample debug log")

		return nil, nil
	})

	app.GET("/info", func(c *gofr.Context) (interface{}, error) {
		c.Logger.Info("sample info log")

		return nil, nil
	})

	app.Run()
}
