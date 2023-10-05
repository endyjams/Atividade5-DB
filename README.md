## API Pública (Hospedada em uma instância EC2 na AWS)

- http://3.212.166.89:8000
- Swagger: http://3.212.166.89:8000/docs/index.html

## Como testar?

- go build
- ./{go_generated_file}.exe || go run ./main.go
- projeto vai rodar na porta 8000
- acesse http://localhost:8000

## Entidades e Endpoints

- Fruta (Model)
- "nome": string
- "preco": float

- Endpoint GET: http://3.212.166.89:8000/obter/fruta/{nome}
- Endpoint POST: http://3.212.166.89:8000/criar/fruta
- Endpoint PUT: http://3.212.166.89:8000/atualizar/fruta
- Endpoint DELETE: http://3.212.166.89:8000/deletar/fruta
----------------------------------------------------------------------
- Fornecedor (Model)
- "nome": string
- "telefone": string

- Endpoint GET: http://3.212.166.89:8000/obter/fornecedor/{nome}
- Endpoint POST: http://3.212.166.89:8000/criar/fornecedor
- Endpoint PUT: http://3.212.166.89:8000/atualizar/fornecedor
- Endpoint DELETE: http://3.212.166.89:8000/deletar/fornecedor
----------------------------------------------------------------------
- Estoque (Model)
- "nomeFruta": string
- "nomeFornecedor": string
- "quantidade": int

- Endpoint GET: http://3.212.166.89:8000/obter/estoque/{nomeFornecedor}/{nomeFruta}
- Endpoint POST: http://3.212.166.89:8000/criar/estoque
- Endpoint PUT: http://3.212.166.89:8000/atualizar/estoque
- Endpoint DELETE: http://3.212.166.89:8000/deletar/estoque
- Endpoint POST(Transaction on Database): http://3.212.166.89:8000/preencher/estoque

## Gerando e acessando swagger

- swag init
- http://localhost:8000/docs/index.html
