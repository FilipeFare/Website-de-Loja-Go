package modelos

import "Banco/bancodedados"

type Produto struct {
	Id        int
	Nome      string
	Descricao string
	Preco     float64
	Estoque   int
}

func AdicionaProdutos() []Produto {
	db := bancodedados.BancoDeDados()

	selectProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}          // Cria um novo struct produto vazio.
	produtos := []Produto{} // Cria uma lista vazia para armazenar produtos.

	for selectProdutos.Next() {
		var id, estoque int
		var nome, descricao string
		var preco float64

		err := selectProdutos.Scan(&id, &nome, &descricao, &preco, &estoque) // Lê os dados da linha atual e armazena nas variáveis.
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Estoque = estoque

		produtos = append(produtos, p) // Adiciona o struct produto à lista de produtos.
	}
	defer db.Close()
	return produtos
}

func CriaProdutoNovo(nome, descricao string, preco float64, estoque int) {
	db := bancodedados.BancoDeDados()

	insereDados, err := db.Prepare("INSERT INTO produtos (nome, descricao, preco, estoque) VALUES ($1, $2, $3, $4)") // Prepara a instrução SQL para inserção
	if err != nil {
		panic(err)
	}

	// Os parâmetros $1, $2, $3 e $4 serão substituídos pelos valores de nome, descricao, preco e estoque, respectivamente.
	_, err = insereDados.Exec(nome, descricao, preco, estoque) // Executa a instrução SQL
	if err != nil {
		panic(err)
	}
	defer insereDados.Close()

}

func DeletaProduto(id string) {
	db := bancodedados.BancoDeDados()

	deletarProduto, err := db.Prepare("DELETE FROM produtos WHERE id=$1")

	if err != nil {
		panic(err.Error())
	}

	_, err = deletarProduto.Exec(id)
	if err != nil {
		panic(err)
	}
	defer db.Close()

}

func EditaProduto(id string) Produto {
	db := bancodedados.BancoDeDados()

	produtoBanco, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	pAtualizar := Produto{}

	for produtoBanco.Next() {
		var id, estoque int
		var nome, descricao string
		var preco float64

		err := produtoBanco.Scan(&id, &nome, &descricao, &preco, &estoque)
		if err != nil {
			panic(err.Error())
		}
		pAtualizar.Id = id
		pAtualizar.Nome = nome
		pAtualizar.Descricao = descricao
		pAtualizar.Preco = preco
		pAtualizar.Estoque = estoque
	}
	db.Close()
	return pAtualizar
}

func AtualizaProduto(id int, nome, descricao string, preco float64, estoque int) {
	db := bancodedados.BancoDeDados()

	atualizaDados, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, estoque=$4 where id=$5")
	if err != nil {
		panic(err)
	}

	_, err = atualizaDados.Exec(nome, descricao, preco, estoque, id)
	if err != nil {
		panic(err)
	}
	defer atualizaDados.Close()
}
