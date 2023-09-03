<h1>Go: crie uma aplicação web</h1> 

<p align="center">
  
  <img height="45" align="center" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" />
  <img height="45" align="center" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/postgresql/postgresql-original-wordmark.svg" />

       
</p>

> Status do Projeto: :warning: em desenvolvimento

### Tópicos 

:small_blue_diamond: [Descrição do projeto](#descrição-do-projeto)

<!-- :small_blue_diamond: [Ultimas atualizações](#ultimas-atualizações)

:small_blue_diamond: [Funcionalidades](#funcionalidades)  -->

:small_blue_diamond: [Pré Requisitos](#pré-requisitos)

:small_blue_diamond: [Como configurar a aplicação](#como-configurar-a-aplicação-arrow_forward)


... 

## Descrição do projeto 

<p align="justify">
  O Projeto foi criado como projeto do curso de mesmo nome oferecido pela <a href="https://cursos.alura.com.br/course/go-lang-web" > alura </a>. Este projeto tem como finalidade o desenvolvimento de uma loja online com adição e exclusão de items em um banco de dados postgres
</p>

<!-- ## Ultimas atualizações :new:
<p align="justify">
  O projeto encontra-se na versão 1.0 e atualmente verifica o site da <a href="http://www.alura.com.br"> alura </a>, e outros 4 endereços, sendo 3 com status fixos e um ultimo com status variavel. A aplicação retorna se o site está ativo ou não. O código está modularizado e executa comandos de leitura, manipulação e criação de arquivos txt
</p>

## Funcionalidades

:heavy_check_mark: Monitoramento de status
- Leitura de arquivo texto com sites os quais devem ser monitorados

:heavy_check_mark: Log de status
- Criação e manipulação de arquivo texto com log de status -->

## Pré-requisitos

:warning: [GO](https://medium.com/xp-inc/primeiros-passos-com-golang-1abdc60bba50)

:warning: [POSTGRES](https://www.postgresql.org/download/)

:warning: [NODEJS](https://nodejs.org/en/download)


## Como configurar a aplicação :arrow_forward:

Após instalado e configurado o Postgres crie um novo database chamado alura_loja:

```
CREATE DATABASE alura_loja
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Portuguese_Brazil.1252'
    LC_CTYPE = 'Portuguese_Brazil.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;
```

Após a criação de banco de dados execute o seguinte script dentro dele:
```
Create table produtos(
    id serial primarykey,
    nome varchar,
    descricao varchar,
    preco decimal,
    quantidade int
)   
```

Agora no terminal, clone o projeto dentro da sua pasta "src" - no mesmo diretório o qual o Go foi instalado.: 

```
git clone https://github.com/GabrielP0rt0/Go-a-linguagem-do-Google
```

Abra o projeto e execute o seguinte comando no terminal:

```
npm install
```

> Esse comando será o responsável por instalar o bootstrap, necessário para execução de nossas páginas html

## Como rodar os testes

Abra o terminal dentro da pasta a qual se encontra o arquivo main.go e digite:

```
go run main.go
```


Copyright :copyright: 2023 - Go-a-linguagem-do-Google
