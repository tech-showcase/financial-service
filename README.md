## FINANCIAL SERVICE

### Description
This repo contains project that handle financial-related services. 
This service is part of a big system. 
The whole system will be used to present technology show case.

### Features
- Serve digital/crypto currency data

This service serve data that is mentioned above through GRPC.

### API
Please refer to all proto file [here](proto) for more detail about the provided API.
You can use GRPC by:
- Installing [protoc](http://google.github.io/proto-lens/installing-protoc.html)
- Installing [protoc-gen-go](https://grpc.io/docs/languages/go/quickstart/)
- Generating code by executing `protoc  --go_out=plugins=grpc:. proto/*/*.proto`


### How to run
#### Docker
- Install docker
- Build and run docker image as below
```shell script
$ docker build -t financial-service .
$ docker run -p 8082:8082 financial-service
```
