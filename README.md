# go-rest

Go REST api with [go-chi](https://github.com/go-chi/chi)

## Usage

```sh
go run cmd/url-shortener/main.go --config=path/to/config.yaml
```

### Docker

For local usage see [docker-compose.local.yml](docker-compose.local.yml) and [Dockerfile.local](Dockerfile.local)
This is also deployed with [docker-compose.yml](docker-compose.yml) and Github Actions through SSH
Necessary repository secrets:
- `AUTH_PASS`: password for `HTTP_SERVER_PASSWORD`
- `DEPLOY_HOST`: ip or hostname of the server to deploy to
- `DEPLOY_USERNAME`: username of the server to deploy to
- `DEPLOY_SSH_KEY`: private key of the server to deploy to

## TODO

- [ ] implement PostgreSQL storage
- [x] add Dockerfile (+ compose?)
- [ ] add docs for endpoints
- [ ] add auth?
