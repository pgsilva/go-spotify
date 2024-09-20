# 🎧 Spodify - Podcast Search API

Bem-vindo ao **Spodify**! Este projeto é uma API escrita em Go que interage com a API do Spotify para realizar buscas de podcasts. 🚀

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
POSTGRES_USER=usr
POSTGRES_PASSWORD=pass
POSTGRES_DB=maindb
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
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

###  ✔️ Health Check

```sh
curl -X GET http://localhost:3000/api/v1/health
```


### 🎙️ Get Search Podcasts

- **Rota**: /v1/api/spodify/podcast/search

    
```sh
curl -X GET http://localhost:3000/v1/api/spodify/podcast/search?q=nerdcast
```

### 🎙️Get Podcast

- **Rota**: /v1/api/spodify/podcast/:id



```sh
curl -X GET http://localhost:3000/v1/api/spodify/podcast/22Wgt4ASeaw8mmoqAWNUn1
```

### 🎙️Get Episódios

- **Rota**: /v1/api/spodify/podcast/:id/episode 


```sh
curl -X GET http://localhost:3000/v1/api/spodify/podcast/22Wgt4ASeaw8mmoqAWNUn1/episode
```

### 🎙️Get Episódio 

- **Rota**: /v1/api/spodify/podcast/:id/episode/:episode_id
    
```sh
curl -X GET http://localhost:3000/v1/api/spodify/podcast/22Wgt4ASeaw8mmoqAWNUn1/episode/0k6Z2J0J4BjK2Q3CJ4n8kP
```

### ▶️ Put Play

- **Rota**: /v1/api/spodify/podcast/:id/episode/:episode_id/play


```sh
curl -X PUT http://localhost:3000/v1/api/spodify/podcast/22Wgt4ASeaw8mmoqAWNUn1/episode/0k6Z2J0J4BjK2Q3CJ4n8kP/play
```

### 🛑 Put Pause

- **Rota**: /v1/api/spodify/podcast/:id/episode/:episode_id/pause

```sh
curl -X PUT http://localhost:3000/v1/api/spodify/podcast/22Wgt4ASeaw8mmoqAWNUn1/episode/0k6Z2J0J4BjK2Q3CJ4n8kP/pause
```

### 📱 Get Devices

- **Rota**: /v1/api/spodify/devices

```sh
curl -X GET http://localhost:3000/v1/api/spodify/devices
```

### ⭐ Get Favorite Devices

- **Rota**: /v1/api/spodify/db/device

```sh
curl -X GET http://localhost:3000/v1/api/spodify/db/device
```

### ⭐ Post Favorite Device

- **Rota**: /v1/api/spodify/db/device

```sh
curl -X POST http://localhost:3000/v1/api/spodify/db/device -d '{"id":"1","name":"device1", "type":"type1"}'
```

### ⭐ Delete Favorite Device

- **Rota**: /v1/api/spodify/db/device/:id

```sh
curl -X DELETE http://localhost:3000/v1/api/spodify/db/device/1
```

### ⭐ Get Favorite Podcasts

- **Rota**: /v1/api/spodify/db/podcast

```sh
curl -X GET http://localhost:3000/v1/api/spodify/db/podcast
```

### ⭐ Post Favorite Podcast

- **Rota**: /v1/api/spodify/db/podcast

```sh

curl -X POST http://localhost:3000/v1/api/spodify/db/podcast -d '{"id":"1","name":"podcast1", "type":"type1", "uri":"uri1", "poster":"poster1"}'
```

### ⭐ Delete Favorite Podcast

- **Rota**: /v1/api/spodify/db/podcast/:id

```sh
curl -X DELETE http://localhost:3000/v1/api/spodify/db/podcast/1
```



# 🤝 Contribuições
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests.

---

👽 Mr Morales was here!