package main

import (
	"hades/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":3077")
}
