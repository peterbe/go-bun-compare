package main

import (
	"log"
	"os"

	"github.com/valyala/fasthttp"
)

var jsonData []byte

func main() {

	// Define request handler
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		var err error
		jsonData, err = os.ReadFile("data.json")
		if err != nil {
			log.Fatalf("Failed to read JSON file: %v", err)
		}

		// Set content type
		ctx.Response.Header.Set("Content-Type", "application/json")

		// Write JSON data
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.Write(jsonData)
	}

	// Start server
	addr := ":3000"
	log.Printf("Server starting on %s", addr)
	if err := fasthttp.ListenAndServe(addr, requestHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}
