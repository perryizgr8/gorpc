Commands:

```
> python -m grpc_tools.protoc -Iprocedures --python_out=. --pyi_out=. --grpc_python_out=. procedures/procedures.proto

> protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative procedures/procedures.proto
```