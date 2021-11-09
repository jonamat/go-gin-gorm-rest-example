# go-gin-gorm-rest-example

Simple Go authenticated RESTful api server built with [Gin](https://github.com/gin-gonic/gin) and [Gorm](https://github.com/go-gorm/gorm).<br>
It uses Postgres as main DB and Redis to handle client sessions.

## Development

Start the server using `docker-compose up`

### Start devmode with docker-compose
- Start the required services with the command `docker-compose up redis postgres`
- Start Go in watch mode with [Gow](https://github.com/mitranim/gow) `make watch`

## Build binary
- Build dynamically linked binary with `make build`
- Build statically linked binary with `make build-static`

## Build Docker image
Build docker image with `make build-docker`


## Create the first user
Connect to Postgres with your favorite client and insert<br>
`INSERT INTO public.users (id, created_at, updated_at, deleted_at, email, "password", "level") VALUES(1, NULL, NULL, NULL, 'demo@demo.demo', 'demo', 'admin');`<br>
Now you can login to the server with 
```
{
  "email": "demo@demo.demo",
  "password": "demo",
  "level": "admin"
}
```