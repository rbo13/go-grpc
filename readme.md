# GO-GRPC

##### COMPILING
```
$ cd go-grpc
$ protoc --go_out=plugins=grpc:. proto/blockchain.proto
```

##### RUNNING

Run server first:
```
$ cd go-grpc
$ go run server/main.go
```

Open new terminal, run the client as a CLI:
```
$ go run client/main.go
```

Flags supported for the CLI:
```
--list, - Lists the blockchain;
--add,  - Adds new data to blockchain;
```

Sample:
```
$ go run client/main.go --add
$ go run client/main.go --list
```