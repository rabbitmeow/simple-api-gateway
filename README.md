# Simple API Gateway

Written in [Go](https://golang.org). This gateway using [jwt](https://jwt.io) for authenticating the request. Make sure the `jwt_server_key` in `config.json` is same as jwt server key in your token generate process.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

**!! WARNING !!** This project is using Go [Modules](https://blog.golang.org/using-go-modules) which is the minimum version of installed golang is 1.11. 

### Installing

A step by step series of examples that tell you how to get a development env running
1. Clone this project `git clone https://github.com/vinbyte/simple-api-gateway.git` outside $GOPATH (for version >= 1.11) or inside $GOPATH/src (for version < 1.11)
2. **(for version under 1.11)** After cloning it into your GOPATH, you need to run this to install all dependencies :
`go mod tidy`
3. Setting up your port, your [jwt](https://jwt.io) server key, or any service in `config.json`
4. If the `mode` is not `production`, it will print all request log including your error. Set the `mode` to `production` to production purpose
5. Run `go run main.go`

## Built With

* [Gin](https://github.com/gin-gonic/gin)
* [Viper](https://github.com/spf13/viper)