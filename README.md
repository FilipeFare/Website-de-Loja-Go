# Banco de Dados Web

Este é um projeto de aplicação web desenvolvido com Go (Golang) que implementa um sistema completo de gerenciamento de produtos (CRUD). A interface é construída com HTML, Bootstrap e JavaScript, enquanto o backend interage com um banco de dados PostgreSQL usando o pacote `database/sql`.

## ✨ Funcionalidades

- Visualização da lista de produtos
- Cadastro de novos produtos
- Edição dos dados dos produtos
- Exclusão de produtos
- Interface web responsiva com Bootstrap
- Interações dinâmicas com JavaScript

## 🔧 Tecnologias Utilizadas

- Go (Golang) – Backend e rotas
- PostgreSQL – Banco de dados relacional
- HTML5 – Estrutura das páginas
- Bootstrap – Estilização e layout responsivo
- JavaScript – Funcionalidades dinâmicas na interface

## 📂 Estrutura do Projeto

Banco de Dados Web  
├── apresentacao/      – Templates HTML com Bootstrap e JavaScript  
├── bancodedados/      – Conexão e operações com o PostgreSQL  
├── controlador/       – Controladores que lidam com a lógica de negócio  
├── modelos/           – Definição das estruturas de dados utilizadas pela aplicação  
├── rotas/            – Configuração das rotas da aplicação  
├── main.go           – Arquivo principal que inicia o servidor  
├── go.mod            – Gerenciador de dependências do Go  
└── go.sum            – Verificação de integridade das dependências  

## ▶️ Como Executar o Projeto

Siga os passos abaixo para rodar a aplicação localmente:

1. Configure o banco de dados PostgreSQL:  
   Crie um banco de dados no PostgreSQL (ex: produtosdb).  
   Crie a tabela necessária usando o seguinte SQL:  
   `CREATE TABLE produtos ( id SERIAL PRIMARY KEY, nome TEXT, descricao TEXT, preco NUMERIC, quantidade INT );`

2. Atualize a conexão com o banco de dados:  
   No arquivo `bancodedados/db.go`, atualize os dados de conexão conforme sua configuração:  
   `const stringDeConexao = "user=SEU_USUARIO dbname=produtosdb password=SUA_SENHA host=localhost sslmode=disable"`

3. Execute a aplicação:  
   `go run main.go`

4. Acesse no navegador:  
   `http://localhost:7000`

---

**Projeto desenvolvido por [FilipeFare].**
