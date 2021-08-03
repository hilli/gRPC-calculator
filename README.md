# gRPC-calculator

This is a test on making an gRPC server and client.

## Running it

### Server

```bash
go run calculator/calculator-server/server.go
```

### Client

```bash
go run calculator/calculator-client/client.go
```

## Developing Setup

- Make sure you are on go >= 1.16
- Have `protoc` installed (Ie `brew install protobuf`) to generate code from the `.proto` files
- Run `./devsetup.sh` to get the go specific generators installed to your `$GOPATH/bin`.
