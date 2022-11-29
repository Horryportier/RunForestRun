package main

import (
	app "code/forrest/app"
	"log"
)


func main() {
        err := app.Start()
        if err != nil {
                log.Fatal(err)
        }
}

