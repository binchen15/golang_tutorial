
```
python -m grpc_tools.protoc -I../protos --python_out=. --grpc_python_out=. ../protos/helloworld.proto
```

helloworld_pb2.py 
	contains our generated request and response classes
helloworld_pb2_grpc.py 
	contains our generated client and server classes.
