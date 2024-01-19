package main

import (
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var cwd string
	cwd, _ = os.Getwd()

    var app *fiber.App=fiber.New(fiber.Config{
        CaseSensitive:true,
        EnablePrintRoutes:false,
    })

    app.Static("/",filepath.Join(cwd,"../build"))

    app.Listen(":4200")
}