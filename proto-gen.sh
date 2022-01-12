# protoc -I internal/proto-files/domain --go_out=internal/grpc internal/proto-files/domain/repository.proto
protoc -I . --go_out=paths=source_relative:. api/k8s/domain/v1beta1/repository.proto
# protoc --go_out=internal/grpc --proto_path=plugins=grpc:internal/proto-files/service/repository-service.proto
protoc -I . --go_out=paths=source_relative:. api/k8s/service/v1beta1/repository-service.proto