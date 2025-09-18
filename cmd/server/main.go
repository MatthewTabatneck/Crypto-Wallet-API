package main

import (
	"log"
	"net/http"

	"github.com/MatthewTabatneck/Crypto-Wallet-API/internal/handlers"
)

func main() {
	http.HandleFunc("/ping", handlers.PingHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
