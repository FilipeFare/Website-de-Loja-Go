package rotas

import (
	"Banco/controlador"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controlador.Index)  // Registra uma função para tratar requisições na rota especificada
	http.HandleFunc("/new", controlador.New) // Chama a outra página que no caso se chama "New"
	http.HandleFunc("/insert", controlador.Insert)
	http.HandleFunc("/delete", controlador.Delete)
	http.HandleFunc("/edit", controlador.Edit)
	http.HandleFunc("/update", controlador.Update)
}
