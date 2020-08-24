$ go get github.com/golang/protobuf
$ go get github.com/golang/protobuf/proto
$ brew install protobuf

```
	$ protoc --go_out=. *.proto
```
go run main.go person.pb.go 
