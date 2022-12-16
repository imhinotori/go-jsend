# go-jsend

A very simple implementation of JSend specification on Go

## Getting Started

### Installing

```shell
go get github.com/imhinotori/go-jsend
```

### Usage

Success:
```go
jsend.Success(data)
```

Fail:
```go
jsend.Fail(data)
```

Error (Message Only):
```go
jsend.Fail(message)
```
Error (Message & Code):
```go
jsend.ErrorWithCode(message, code)
```
Error (Message & Data):
```go
jsend.ErrorWithData(message, data)
```
Error (Message, Code & Data)
```go
jsend.ErrorFull(message, code, data)
```

## License

This project is licensed under the MIT License - see the LICENSE.md file for details

## Acknowledgments

* [Omniti-Labs JSend](https://github.com/omniti-labs/jsend)
* [CleverGo](https://github.com/clevergo/jsend)
* [Gamegos JSend](https://github.com/gamegos/jsend)