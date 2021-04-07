# Zoo service

This is a service provided for the accompanying presentation. You can build your own frontend to display, update or add animals.

## build

In order to build this service, you need to install some dependencies:

- go
- swag
- make (optional)
- docker (optional)

### how to build (linux)

The easy way:

```sh
make all # this should take care of everything that needs to be done.
```

If you don't want to install make or docker:

```sh
go mod download
go install github.com/swaggo/swag/cmd/swag
swag init -g ./main.go
go build -o app ./main.go

```

## testing

I have not written any tests yet 😅

## how to run

This is assuming, you have used the `make all` command.

Once the project is build, you just need to run the docker container:

```sh
docker run --rm -it -p 8080:8080 docker.io/library/zooservice
```

- The documentation is available under [http://localhost:8080/docs/index.html](http://localhost:8080/docs/index.html)
- The API is available under [http://localhost:8080/api/v0](http://localhost:8080/api/v0)
