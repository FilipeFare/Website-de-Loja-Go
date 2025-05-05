package controlador

import (
	"Banco/modelos"
	"net/http"
	"strconv"
	"text/template"
)

var apresent = template.Must(template.ParseGlob("apresentacao/*.html")) // Encapsula o template e verifica se ocorreu algum erro ao carregá-lo

func Index(w http.ResponseWriter, r *http.Request) { // Responde às requisições na página inicial e mostra o template index

	produtos := modelos.AdicionaProdutos()
	println("Página inicial acessada")
	erro := apresent.ExecuteTemplate(w, "Index", produtos)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
}

func New(w http.ResponseWriter, r *http.Request) {
	apresent.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {     // O método POST é usado para enviar dados ao servidor, como ao enviar um formulário.
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		estoque := r.FormValue("estoque")

		precoConvertido, err := strconv.ParseFloat(preco, 64) //converte a string que vem do html
		if err != nil {
			http.Error(w, "Erro na conversão do preço", http.StatusBadRequest)
			return
		}

		estoqueConvertido, err := strconv.Atoi(estoque) //converte a string que vem do html
		if err != nil {
			http.Error(w, "Erro na conversão da estoque", http.StatusBadRequest)
			return
		}

		modelos.CriaProdutoNovo(nome, descricao, precoConvertido, estoqueConvertido)

		http.Redirect(w, r, "/", http.StatusMovedPermanently) // Volta para o ínicio
	} else {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")     // Pega o ID do produto a partir da requisição na URL
	modelos.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Edit(w http.ResponseWriter, r *http.Request) { // Formulário para preencher as informações nos campos a serem editados.
	id := r.URL.Query().Get("id")
	produto := modelos.EditaProduto(id)
	apresent.ExecuteTemplate(w, "edit", produto)

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {     // O método POST é usado para enviar dados ao servidor, como ao enviar um formulário.
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		estoque := r.FormValue("estoque")
	

	idConvertido, err := strconv.Atoi(id) //converte a string que vem do html
		if err != nil {
			http.Error(w, "Erro na conversão do preço", http.StatusBadRequest)
			return
		}

	precoConvertido, err := strconv.ParseFloat(preco, 64) 
		if err != nil {
			http.Error(w, "Erro na conversão do preço", http.StatusBadRequest)
			return
		}

		estoqueConvertido, err := strconv.Atoi(estoque) 
		if err != nil {
			http.Error(w, "Erro na conversão da estoque", http.StatusBadRequest)
			return
		}

		modelos.AtualizaProduto(idConvertido, nome, descricao, precoConvertido, estoqueConvertido)

		http.Redirect(w, r, "/", http.StatusMovedPermanently) // Volta para o ínicio
	} else {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}