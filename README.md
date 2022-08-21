# News API

A microservice to store and retrieve news posts.

## Technologies

There are HTTP server written in Go and PostgreSQL storage.

## Getting Started

> Only `docker` and `docker-compose` are needed.

Run HTTP server and database with single command:

```shell
docker-compose up -d
```

Server will be accessible via port **9010** by default.

## Requests

GET /posts
`curl --location --request GET 'http://localhost:9010/posts`
or
`curl --location --request GET 'http://localhost:9010/posts?limit={limit}&ptoken={lastID}'`

Example
`curl --location --request GET 'http://localhost:9010/posts'`, `curl --location --request GET 'http://localhost:9010/posts?limit=100'`, `curl --location --request GET 'http://localhost:9010/posts?limit=2&ptoken=8da1644b-6fa9-414d-b11f-252fd994f461'`.

POST /posts
`curl --location --request POST 'http://localhost:9010/posts' --header 'Content-Type: application/json' --data-raw '[{"title":"{news_title}","content":"{news_content}"}]'`

Example
`curl --location --request POST 'http://localhost:9010/posts' --header 'Content-Type: application/json' --data-raw '[{"title":"1","content":"1"},{"title":"2","content":"2"}]'`

GET /posts/{uuid}
`curl --location --request GET 'http://localhost:9010/posts/{uuid}`

Example
`curl --location --request GET 'http://localhost:9010/posts/8da1644b-6fa9-414d-b11f-252fd994f460'`

PUT /posts/{uuid}
`curl --location --request PUT 'http://localhost:9010/posts/{uuid}' --header 'Content-Type: application/json' --data-raw '{"title":"news.title","content":"news.content"}'`

Example
`curl --location --request PUT 'http://localhost:9010/posts/db20f5fd-7ca4-4fde-b3d3-c7a2580a2eea' --header 'Content-Type: application/json' --data-raw '{"title":"666","content":"666"}'`

DELETE /posts/{uuid}
`curl --location --request DELETE 'http://localhost:9010/posts/{uuid}'`

Example
`curl --location --request DELETE 'http://localhost:9010/posts/db20f5fd-7ca4-4fde-b3d3-c7a2580a2eea'`


## Development

> Prerequisites: `go@1.18`, `make` must be installed.

All necessary configuration stored in [docker-compose.yml](docker-compose.yml).

### Binary

Build the binary `bin/news-api`:

```shell
make build
```

See also [Makefile](Makefile) for all available targets.

### Tests

Run unit tests:

```shell
make test
```

### Code style

Consistent code style enforced by `gofmt`, `EditorConfig` tools and `golangci-lint` linter.

Run linter:

```shell
make lint
```
