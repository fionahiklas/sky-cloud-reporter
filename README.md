## Overview

Coding exercise to read cloud instance data and present as a single, unified output

## Tools

### Docker

Example cloud APIs are provided as docker comntainers.  Installing Docker Desktop 
for Mac from [here](https://www.docker.com/products/docker-desktop)

### Go

Downloading the latest (1.15.7) from the [official site](https://golang.org/doc/install?download=go1.15.7.darwin-amd64.pkg)


## Getting Started

### Set GOPATH

```
export GOPATH=$HOME/wd/gobase
```

### Get GoMock 

```
go get github.com/golang/mock/mockgen@v1.4.4
```

### Get code

Get this code

```
go get github.com/fionahiklas/sky-cloud-reporter
```



## Notes

### Creating basic module setup

Ran the following command

```
go mod init github.com/fionahiklas/sky-cloud-reporter
```

This created the initial `go.mod` file

### Building the test HTTP client

```
go install cmd/testhttpclient/test_http_client.go
```

Will build and link the code and output an executable file called `test_http_client` and 
place it under `$GOPATH/bin`

Run using the following command

```
$GOPATH/bin/test_http_client http://localhost:9002/cloud/instances
```


## References

* [Go code organisation](https://golang.org/doc/code.html)
* [Print a variables type](https://golangcode.com/print-variable-type/) 
* [JSON Parsing](https://gobyexample.com/json)
* [Gomock](https://github.com/golang/mock)