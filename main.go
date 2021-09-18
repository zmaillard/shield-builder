package main

import (
	"log"
	"net/http"
	"os"
	"sign-builder/core"
	"sign-builder/handlers"
)

func main() {
	core.Init()
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/saveshield", handlers.HandleShieldPostQuery)
	mux.HandleFunc("/api/getshield", handlers.HandleShieldQuery)

	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, mux))
}
