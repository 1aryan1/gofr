package main

import "gofr.dev/pkg/gofr"

func main() {
    
    app := gofr.New()

    
    app.GET("/homePage", func(ctx *gofr.Context) (interface{}, error) {

        return "Hello from the other side ", nil
    })

    app.Start()
}