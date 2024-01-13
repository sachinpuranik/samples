# protoc --proto_path=api/proto/v1  --go_out=plugins=grpc:pkg/api/v1 todo-service.proto

# protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative example.proto

protoc --go_out=./ --go_opt=paths=import --go-grpc_out=./ --go-grpc_opt=paths=import ./api/proto/v1/todo-service.proto
