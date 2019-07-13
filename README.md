# Go Gin Boilerplate
> A starter project with Golang, Gin and DynamoDB

[![Build Status][travis-image]][travis-url]
[![codebeat badge](https://codebeat.co/badges/ed248580-942c-4ffc-919f-d3681d28a799)](https://codebeat.co/projects/github-com-vsouza-go-gin-boilerplate)
[![Go Version][go-image]][go-url]
[![License][license-image]][license-url]
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fvsouza%2Fgo-gin-boilerplate.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fvsouza%2Fgo-gin-boilerplate?ref=badge_shield)


Golang Gin boilerplate with DynamoDB resource. Supports multiple configuration environments.

![](header.jpg)

This project use a [DynamoDB base docker image](https://github.com/vsouza/docker-dynamoDB-local).

Setup DynamoDB dependency:  `docker run -d -p 8080:8080 vsouza/dynamo-local --port 8080`

### Boilerplate structure

```
.
├── Makefile
├── Procfile
├── README.md
├── config
│   ├── config.go
│   ├── development.yaml
│   ├── production.yaml
│   └── test.yaml
├── controllers
│   └── user.go
├── db
│   └── db.go
├── forms
│   └── user.go
├── header.jpg
├── main.go
├── middlewares
│   └── auth.go
├── models
│   └── user.go
└── server
    ├── router.go
    └── server.go
```

## Installation

__Install Godeps__

`go get github.com/tools/godep`

then run:

```sh
make deps
```

## Usage example

`curl http://localhost:8888/health`

## Development setup

Running DynamoDB on Docker Image:

check this project: [vsouza/docker-dynamoDB-local](https://github.com/vsouza/docker-dynamoDB-local)

Setup Glide:

[Glide Package Management](https://glide.sh)
