package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthHandler(http.ResponseWriter, *http.Request) error {
	mux := app.mount()
	fmt.Println("Server setup OK ✅")
	fmt.Println("Mux:", mux) // just to show that mount() worked
	return nil
}
