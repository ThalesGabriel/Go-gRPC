# Go-gRPC

## O que é e o que faz esse app?

- Esta é uma pequena aplicação GoLang integrada com gRPC;
- Não possui um banco de dados;
- Aqui temos a implementação do cliente e servidor gRPC;
- O cliente envia uma requisição, o servidor faz alguma alteração simulando um DB e retorna ao cliente

# pequena introdução ao gRPC 

- gRPC é uma forma de conexão cliente servidor ideal para microserviços mas não substitui o REST
- gRPC é mais rápida já que possui suporte ao HTTP 2.0, ou seja, as requisições e respostas podem ser feitas na mesma conexão
- gRPC utiliza protocol buffers em vez do json utilizado pelo rest;
- Os protocol buffers são bem mais eficientes visto que são dados binários em vez do plaintext do json;

## Métodos

- Esse app é centrado em registros aqui chamados `Users`
- Temos a função `AddUser` em que o cliente envia um dado para o servidor e o servidor retorna alterando o `ID`
- Temos a função `AddUserVerbose` em que o stream de dados é mostrado. O cliente envia uma requisição de adição de usuário e o servidor retorna o `status` da adição a cada passo, ou seja, `Inited` ... `Completed`
- Temos a função `AddUsers` em que mais uma vez utilizamos o stream porém dessa vez apenas do lado do cliente. Criamos uma lista mocada e com um laço de repetição a cada `3seg` enviamos um usuário para o servidor. Quando os dados param de ser recebidos no lado do servidor, o servidor retorna a lista de usuários.
- Temos a função `AddUserStreamBoth` que é auto explicativa. O cliente utiliza a mesma lista da função anterior e a cada cliente enviado já recebe a resposta do usuário adicionado na lista

## Para rodar o App

- primeiramente certifique-se de ter o `Go` instalado na sua máquina ou utilize `containers`
- Como mencionado anteriormente a aplicação se divide em cliente servidor então você precisará de duas instâncias abertas do seu terminal
- Vá na pasta `cmd/client/main.go`
- Escolha a função que você quer visualizar na função `main`
- Primeiramente execute o servidor `go run cmd/server/main.go`
- Em seguida execute o cliente `go run cmd/client/main.go`
