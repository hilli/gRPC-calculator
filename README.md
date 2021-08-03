# gRPC-calculator

This is a test on making an gRPC server and client.

## Running it from repo


#### Server


```bash
go run calculator/calculator-server/server.go
```

#### Client


```bash
go run calculator/calculator-client/client.go
```

or

```bash
go run calculator/calculator-client/client.go 609 57
```

## Install the binaries

If you are _really_ fond of adding 2 numbers in a client server manner, you can install the binaries directly:

```bash
go install github.com/hilli/gRPC-calculator/calculator/calculator-server
go install github.com/hilli/gRPC-calculator/calculator/calculator-client
```

And run it like

```bash
$ calculator-server &
[1] 46937
2021/08/03 14:53:53 server listening at 127.0.0.1:50051
$ calculator-client 5 9
2021/08/03 14:54:00 Starting Sum request
2021/08/03 14:54:00 Sum result: 14
```

## Developing Setup

- Make sure you are on go >= 1.16
- Have `protoc` installed (Ie `brew install protobuf`) to generate code from the `.proto` files
- Run `./devsetup.sh` to get the go specific generators installed to your `$GOPATH/bin`.
