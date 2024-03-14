package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Printf("WebApp inicializado na porta 3003...")
	log.Fatal(http.ListenAndServe(":3003", r))
}
