package main

import (
	"github.com/hohmannr/go-redoc"
	echoredoc "github.com/hohmannr/go-redoc/echo"
	"github.com/labstack/echo/v4"
)

func main() {
	doc := redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./openapi.json",
		SpecPath:    "/openapi.json",
		DocsPath:    "/docs",
	}

	r := echo.New()
	r.Use(echoredoc.New(doc))

	println("Documentation served at http://127.0.0.1:8000/docs")
	panic(r.Start(":8000"))
}
