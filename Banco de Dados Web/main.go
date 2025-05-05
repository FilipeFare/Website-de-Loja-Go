package main

import (
	"net/http"
	"Banco/rotas"
)

func main() {
	rotas.CarregaRotas()
	http.ListenAndServe(":7000", nil) // Inicia o servidor HTTP na porta 8000
}
