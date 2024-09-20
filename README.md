# 🎧 Spodify - Podcast Search API

Bem-vindo ao **Spodify**! Este projeto é uma API escrita em Go que interage com a API do Spotify para realizar buscas de podcasts.

## 📋 Especificações do Projeto

- **Linguagem**: Go
- **Framework Web**: Fiber
- **API Externa**: Spotify API
- **Funcionalidade Principal**: Busca de podcasts

## 🚀 Como Executar

### Pré-requisitos

- Go 1.16+
- Conta no Spotify Developer e credenciais da API

### Passos para Executar

1. **Clone o repositório**:

```sh
git clone git@github.com:pgsilva/go-spotify.git
cd go-spotify
``` 

2. **Configure suas credenciais do Spotify**:

Adicione o arquivo `.env` na raiz do projeto com as seguintes variáveis:

```sh
PORT=3000
SPOTIFY_CLIENT_ID={your_client_id}
SPOTIFY_CLIENT_SECRET={your_client_secret}
SPOTIFY_CONTENT_API_URL=https://api.spotify.com/v1/
SPOTIFY_ACCOUNT_API_URL=https://accounts.spotify.com/api/
````

3. **Instale as dependências**:

```sh
go mod tidy
```

ou

```sh
go mod download
```

4. **Execute a aplicação**:

```sh
go run cmd/main.go
```

A API estará disponível em `http://localhost:3000`.

## 📚 Endpoints

- ###  ✔️ Health Check

    - **Rota**: /api/v1/health

    - **Método**: GET

**Exemplo de Requisição**:

```sh
curl -X GET http://localhost:3000/api/v1/health
```


- ### 🎙️ Buscar Podcasts

    - **Rota**: /api/v1/spodify/player/search

    - **Método**: GET

    - **Parâmetros**: 
        - `q`: Termo de busca (query string)

**Exemplo de Requisição**:
    
```sh
curl -X GET http://localhost:3000/api/v1/spodify/player/search?q=nerdcast
```

# 🤝 Contribuições
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests.

---

👽 Mr Morales was here!
