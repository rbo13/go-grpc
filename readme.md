# GO-GRPC

##### RUNNING

Run server first:
```
$ cd grpc
$ go run server/main.go
```

Open new terminal, run the client as a CLI:
```
$ go run client/main.go
```

Flags supported for the CLI:
--list, - Lists the blockchain;
--add,  - Adds new data to blockchain;

Sample:
```
$ go run client/main.go --add
$ go run client/main.go --list
```