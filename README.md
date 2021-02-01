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
export PATH=$PATH:$GOPATH/bin
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

### Generate the mocks

```
go generate ./...
```

### Run the tests

```
go test -cover ./...
```

Coverage should be, ideally, > 90%

### Run the command line to grab instances

Start the docker containers

```
docker run -d -p 9001:9001 enrico5b1b4/cloud1-api
docker run -d -p 9002:9002 enrico5b1b4/cloud1-api
```
Build and run the CLI

```
go install ./cmd/grabinstances
grabinstances http://localhost:9001 http://localhost:9002
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

### Using Go Plugin in IntelliJ

This behaves slightly bizarrely compared to Java/Ruby and others that have SDKs which can 
be added under the "Open Module Settings" dialog.  

Under Preferences -> Languages -> Go select GOPATH and add a project specific GOPATH
entry so that the IDE picks up any libraries loaded there

### Getting go packages

``` 
go get github.com/golang/mock/mockgen@v1.4.4
go get github.com/stretchr/testify/assert
```

### Using mockgen with `-self_package`

As per some of the suggestions for solving mocking cycles I tried using the `-self_package`
option in the `go:generate` lines running `mockgen`

``` 
//go:generate mockgen -package=grab -destination=../mocks/grab/mock_http_interfaces.go -self_package=github.com/fionahiklas/sky-cloud-reporter/grab . HttpResponse,HttpClient
```

Without the full path for self package it didn't do what was intended which was to strip 
out the import for the `grab` package.  The problem is that using the line above resulted in code that 
wouldn't work anyway since the HttpResponse is still defined in `grab` and is needed to compile



## References

### Golang

* [Go code organisation](https://golang.org/doc/code.html)
* [Print a variables type](https://golangcode.com/print-variable-type/) 
* [JSON Parsing](https://gobyexample.com/json)
* [Go JSON](https://blog.golang.org/json)
* [Initialising slices of structs](https://stackoverflow.com/questions/26159416/init-array-of-structs-in-go)
* [Appending slices](https://golang.org/pkg/builtin/#append)
* [Creating slices with length/capacity](https://blog.golang.org/slices-intro)
* [Convert between bytes and strings](https://yourbasic.org/golang/convert-string-to-byte-slice/)
* [Slice internals](https://blog.golang.org/slices)
* [for loops, range, and slices](https://gobyexample.com/range)
* [Prettty print JSON](https://golangbyexample.com/print-struct-variables-golang/)
* [Go http package](https://golang.org/pkg/net/http/)


### Testing

* [Gomock](https://github.com/golang/mock)
* [Gomock documentation](https://pkg.go.dev/github.com/golang/mock#readme-running-mockgen)
* [Go assert package](https://github.com/stretchr/testify)
* [Setup/teardown for tests](https://stackoverflow.com/questions/23729790/how-can-i-do-test-setup-using-the-testing-package-in-go)
* [Go testing with mocks](https://blog.codecentric.de/en/2017/08/gomock-tutorial/)
* [Can't compare functions](https://stackoverflow.com/questions/9643205/how-do-i-compare-two-functions-for-pointer-equality-in-the-latest-go-weekly)
* [Further issue with comparing functions](https://github.com/stretchr/testify/issues/182)

### Issues

#### Import cycles with mocks

* [Importing package for dependent interfaces](https://github.com/golang/mock/issues/352)
* [Using different packages to resolve cycles](https://stackoverflow.com/questions/50986170/how-to-avoid-import-cycles-in-mock-generation)
