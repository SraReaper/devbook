package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
)

func main() {
	fmt.Printf("WebApp inicializado...")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":3003", r))
}
