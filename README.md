# My Project

Este projeto é uma aplicação fullstack que combina um frontend em React e um backend em Golang. Ele oferece uma interface de usuário para encurtar URLs e um servidor backend que lida com a geração, armazenamento e redirecionamento de URLs encurtadas.

## Índice

1. [Sobre o Projeto](#sobre-o-projeto)
2. [Tecnologias Utilizadas](#tecnologias-utilizadas)
3. [Instalação e Configuração](#instalação-e-configuração)
   - [Pré-requisitos](#pré-requisitos)
   - [Configuração do Backend](#configuração-do-backend)
   - [Configuração do Frontend](#configuração-do-frontend)
4. [Como Usar](#como-usar)
5. [Estrutura de Pastas](#estrutura-de-pastas)
6. [Contribuindo](#contribuindo)
7. [Licença](#licença)
8. [Contato](#contato)

## Sobre o Projeto

O **My Project** é uma aplicação de encurtamento de URLs que permite aos usuários encurtar URLs longas e personalizar o slug da URL encurtada. O frontend oferece uma interface intuitiva para os usuários interagirem com o serviço, enquanto o backend em Golang gerencia a lógica de encurtamento e redirecionamento de URLs.

## Tecnologias Utilizadas

- **Frontend:**
  - [React](https://reactjs.org/)
  - [HTML5](https://developer.mozilla.org/en-US/docs/Web/Guide/HTML/HTML5)
  - [CSS3](https://developer.mozilla.org/en-US/docs/Web/CSS)
  - [JavaScript (ES6+)](https://developer.mozilla.org/en-US/docs/Web/JavaScript)

- **Backend:**
  - [Golang](https://golang.org/)
  - [MongoDB](https://www.mongodb.com/)
  - [Gorilla Mux](https://github.com/gorilla/mux)
  - [Go Modules](https://blog.golang.org/using-go-modules)

## Instalação e Configuração

### Pré-requisitos

- **Node.js** e **npm** (para o frontend)
- **Golang** (para o backend)
- **MongoDB** (para armazenamento de URLs)

### Configuração do Backend

1. **Clone o repositório e navegue até o diretório do backend:**

   ```bash
   git clone https://github.com/username/my-project.git
   cd my-project/back

2. **Instale as dependências:**
   ```bash
   go mod tidy

3. **Inicie o servidor:**
   
   Certifique-se de que o MongoDB esteja em execução e, em seguida, execute:
   ```bash
   go run main.go

O backend estará rodando em http://localhost:8080.


### Configuração do Frontend

1. **Navegue até o diretório do frontend:**

   ```bash
   cd ../front

2. **Instale as dependências:**
   ```bash
   npm install

3. **Inicie o servidor de desenvolvimento:**
   ```bash
   npm start
   
O frontend estará acessível em http://localhost:3000.

## How to Use
1.  Acesse a interface do usuário em http://localhost:3000.
2.  Insira a URL que você deseja encurtar.
3.  Opcionalmente, insira um slug personalizado para a URL.
4.  Clique em "Shorten URL" para gerar a URL encurtada.
5.  Clique no link encurtado para ser redirecionado à URL original.

## Estrutura de Pastas

## Contribuindo

## License


## Contact

