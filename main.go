package main

import (
	app "RunForestRun/code/app"
	"log"
)


func main() {
        err := app.Start()
        if err != nil {
                log.Fatal(err)
        }
}

