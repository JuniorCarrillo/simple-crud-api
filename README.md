# SimpleCrudApi v1

[![Go](https://github.com/JuniorCarrillo/simple-crud-api/workflows/Go/badge.svg)](https://github.com/JuniorCarrillo/simple-crud-api/actions?query=workflow:"Go")
[![GitHub tag](https://img.shields.io/github/tag/JuniorCarrillo/simple-crud-api?include_prereleases=&sort=semver&color=blue)](https://github.com/JuniorCarrillo/simple-crud-api/releases/)
[![Sourcegraph](https://sourcegraph.com/github.com/JuniorCarrillo/simple-crud-api/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/JuniorCarrillo/simple-crud-api?badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/JuniorCarrillo/simple-crud-api?style=flat-square)](https://goreportcard.com/report/github.com/JuniorCarrillo/simple-crud-api)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/JuniorCarrillo/simple-crud-api/master/LICENSE)
[![JuniorCarrillo - simple-crud-api](https://img.shields.io/static/v1?label=JuniorCarrillo&message=simple-crud-api&color=blue&logo=github)](https://github.com/JuniorCarrillo/simple-crud-api "Go to GitHub repo")
[![stars - simple-crud-api](https://img.shields.io/github/stars/JuniorCarrillo/simple-crud-api?style=social)](https://github.com/JuniorCarrillo/simple-crud-api)
[![forks - simple-crud-api](https://img.shields.io/github/forks/JuniorCarrillo/simple-crud-api?style=social)](https://github.com/JuniorCarrillo/simple-crud-api)
[![OS - macOS](https://img.shields.io/badge/OS-macOS-blue?logo=apple&logoColor=white)](https://www.apple.com/macos/ "Go to Apple homepage")
[![OS - Linux](https://img.shields.io/badge/OS-Linux-blue?logo=linux&logoColor=white)](https://www.linux.org/ "Go to Linux homepage")
[![Made with Docker](https://img.shields.io/badge/Made_with-Docker-blue?logo=docker&logoColor=white)](https://www.docker.com/ "Go to Docker homepage")
[![Made with Go](https://img.shields.io/badge/Go-1-blue?logo=go&logoColor=white)](https://golang.org "Go to Go homepage")
[![Made with MongoDB](https://img.shields.io/badge/MongoDB-3-blue?logo=mongodb&logoColor=white)](https://www.mongodb.com/ "Go to MongoDB homepage")

This is a Rest API for use as a Back End, which is developed in Go, using echo and can be connected to MongoDB. It is a simple platform, as an example for implementing CRUDs in Go with MongoDB.

This platform runs on port `:3000`, which can be modified in the environment variables of the `.env` file, for which an example is provided in `.env.example`.

### Structure of the system
Here is shown the organization and structural components, architecture, and specific organization of the repository.

#### Files in the repository
Within this repository, there are a total of 24 files (including `.env`) and 9 directories. These are shown in the following structural map:
```
.
├── Dockerfile
├── Makefile
├── README.md
├── api
│ ├── api.go
│ ├── dtos
│ │ ├── character.dto.go
│ │ ├── response.dto.go
│ │ └── team.dto.go
│ ├── handlers
│ │ ├── character.handler.go
│ │ ├── healthCheck.handler.go
│ │ └── team.handler.go
│ ├── middlewares
│ │ └── middlewares.go
│ ├── models
│ │ └── character.model.go
│ ├── router
│ │ └── router.go
│ └── tools
│     ├── calculatePages.go
│     ├── capitalize.tool.go
│     └── urlInCtx.tool.go
├── database
│ ├── characters.go
│ └── db.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
└── server.go
```

#### System libraries
This platform was developed in `go 1.19` with a `modules` configuration, making use of the following components or libraries:

```
require (
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/labstack/echo/v4 v4.10.2 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	go.mongodb.org/mongo-driver v1.11.3 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	golang.org/x/time v0.3.0 // indirect
)
```

## Install

1. Clone the repository
```
git clone https://github.com/JuniorCarrillo/simple-crud-api.git
```

2. Install dependencies

```
cd simple-crud-api && make install-example # Or 'make install' if you have environment variables ready or an .env file already created 
```

3. Run application

```
make run-example
```

4. Running on port 3000

```
curl http://localhost:3000/
```

## Install in docker container

1. Pull image
```
docker pull ghcr.io/juniorcarrillo/simple-crud-api:release
```

2. Run application
```
docker run -dp 3000:3000 \
--name simple-crud-api \
-e MONGO_USERNAME=root \
-e MONGO_PASSWORD=123456 \
-e MONGO_HOST=localhost \
-e MONGO_PORT=27017 \
-e MONGO_COLLECTION=simple-crud-api \
-e X_API_KEY=es123456 \
simple-crud-api
```

## Use

Here, the available endpoints or routes within the REST service are listed. It should be noted that to use it, you must have the `x-api-key` header (available in the environment variables or the `.env` file) for the routes under the pattern `/api/v1/`, which have a security middleware.

### POST /api/v1/team (Search in POST method)

Request example:
```
curl --location 'http://localhost:3000/api/v1/team' \
--header 'x-api-key: es123456' \
--header 'Content-Type: application/json' \
--data '{
    "name": "nombre",
    "pages": 1
}'
```

Response example:
```
{
    "info": {
        "count": 3,
        "pages": 1,
        "next": "http://localhost:3000/api/v1/character?pages=2&search=Nombre",
        "prev": ""
    },
    "result": [
        {
            "_id": "6420ccb6d60daa54c8dfe778",
            "name": "Nombre de ejemplo 1",
            "created_at": "2023-03-26T22:52:38.603Z",
            "updated_at": "2023-03-27T15:58:45.921Z"
        },
        {
            "_id": "6420cccd8023772a74e6594e",
            "name": "Nombre de ejemplo 5",
            "created_at": "2023-03-26T22:53:01.093Z"
        },
        {
            "_id": "6420d099ad315b1386be643d",
            "name": "Nombre de ejemplo 5",
            "created_at": "2023-03-26T23:09:13.986Z"
        }
    ]
}
```

### GET /api/v1/character (Search in GET method)

Request example:
```
curl --location 'http://localhost:3000/api/v1/character?pages=1&search=Nombre' \
--header 'x-api-key: es123456'
```

Response example:
```
{
    "info": {
        "count": 3,
        "pages": 1,
        "next": "http://localhost:3000/api/v1/character?pages=2&search=Nombre",
        "prev": ""
    },
    "result": [
        {
            "_id": "6420ccb6d60daa54c8dfe778",
            "name": "Nombre de ejemplo 1",
            "created_at": "2023-03-26T22:52:38.603Z",
            "updated_at": "2023-03-27T15:58:45.921Z"
        },
        {
            "_id": "6420cccd8023772a74e6594e",
            "name": "Nombre de ejemplo 5",
            "created_at": "2023-03-26T22:53:01.093Z"
        },
        {
            "_id": "6420d099ad315b1386be643d",
            "name": "Nombre de ejemplo 5",
            "created_at": "2023-03-26T23:09:13.986Z"
        }
    ]
}
```

### POST /api/v1/character (Create new character)

Request example:
```
curl --location 'http://localhost:3000/api/v1/character' \
--header 'x-api-key: es123456' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Joe",
    "genre": "Male"
}'
```

Response example:
```
{
    "message": "Item saved with id '6421e522bb8c68d830a53eab'.",
    "character": {
        "_id": "6421e522bb8c68d830a53eab",
        "name": "Joe",
        "created_at": "2023-03-27T13:49:06.472259-05:00"
    }
}
```

### GET /api/v1/character/:id (Get character by id)

Request example:
```
curl --location 'http://localhost:3000/api/v1/character/6421e522bb8c68d830a53eab' \
--header 'x-api-key: es123456'
```

Response example:
```
{
    "message": "Item found for id '6421e522bb8c68d830a53eab'.",
    "character": {
        "_id": "6421e522bb8c68d830a53eab",
        "name": "Joe",
        "created_at": "2023-03-27T18:49:06.472Z"
    }
}
```

### PUT /api/v1/character/:id (Update character by id)

Request example:
```
curl --location --request PUT 'http://localhost:3000/api/v1/character/6421e522bb8c68d830a53eab' \
--header 'x-api-key: es123456' \
--header 'Content-Type: application/json' \
--data '{
    "type": "Human"
}'
```

Response example:
```
{
    "message": "The element with the id '6421e522bb8c68d830a53eab' has been updated successfully."
}
```

### DELETE /api/v1/character/:id (Delete character by id)

Request example:
```
curl --location --request DELETE 'http://localhost:3000/api/v1/character/6421e522bb8c68d830a53eab' \
--header 'x-api-key: es123456'
```

Response example:
```
{
    "message": "The element with the id '6421e522bb8c68d830a53eab' has been deleted successfully."
}
```

## Contribution

If you want to contribute to this project, please follow these steps:

1. Fork the repository
2. Create a branch (`git checkout -b feature/new-feature`)
3. Make the necessary changes and commit (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/new-feature`)
5. Open a pull request

## License

[MIT License](https://github.com/JuniorCarrillo/simple-crud-api/blob/master/LICENSE) © Junior Carrillo

### Contact us
- [Email](mailto:soyjrcarrillo@gmail.com)
- [WhatsApp](https://wa.me/+573003328389)
- [Telegram](https://t.me/juniorcarrillo)
- [Facebook](https://fb.com/soyjrcarrillo)
- [Twitter](https://twitter.com/soyjrcarrillo)
- [Instagram](https://instagram.com/soyjrcarrillo)
