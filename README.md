# Construction

## For init project

```shell
go mod download
```

## For build project

```shell
go build crud.go
```

## For init linux

```shell
./crud
```

## For mod development

air

## Configure Environments

you have to add file .env and paste or configure environments

```dotenv
PORT=:8000
DB_HOST=localhost 
DB_USER=postgres 
DB_PASSWORD=root 
DB_NAME=test_go 
DB_PORT=5432 
DB_SSMODE=disable
```

# Docker

### Construct  for docker

```shell
docker build . -t e-commerse-go
```

### Run docker

```shell
docker run -dp 8080:8080 --name e-commerse-go e-commerse-go
```

### Remove image and  instance

```shell
docker rm e-commerse-go 
docker rmi e-commerse-go
```

### stop container

```shell
docker stop e-commerse-go
```

### start container

```shell
docker start e-commerse-go
```




