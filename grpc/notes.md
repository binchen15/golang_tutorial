0. protocol buffer documentation

	https://developers.google.com/protocol-buffers/docs/overview

1. protocol buffer

.proto

define the service and message

```
// The greeter service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```

	compiler
		protoc

2.	gRPC

	server ---- stub

3.	core concepts	
	https://grpc.io/docs/what-is-grpc/core-concepts/

	IDL
		interface defintion language

	Unary vs Streaming RPC

	Meta data
		send between server and client

	channel
		host, port
		state: connected, idle

4.	tutorial using nodejs

	stub constructor on client side
	service descriptor

create server

