protoc --proto_path=./protos --go_out ./internal/pkg/api --go_opt paths=source_relative \
  --go-grpc_out ./internal/pkg/api --go-grpc_opt paths=source_relative \
  --grpc-gateway_out ./internal/pkg/api --grpc-gateway_opt paths=source_relative \
  --openapiv2_out ./internal/pkg/api \
  user.proto

protoc --proto_path=./protos --go_out ./internal/pkg/api --go_opt paths=source_relative \
  --go-grpc_out ./internal/pkg/api --go-grpc_opt paths=source_relative \
  --grpc-gateway_out ./internal/pkg/api --grpc-gateway_opt paths=source_relative \
  --openapiv2_out ./internal/pkg/api \
  product.proto

protoc --proto_path=./protos --go_out ./internal/pkg/api --go_opt paths=source_relative \
  --go-grpc_out ./internal/pkg/api --go-grpc_opt paths=source_relative \
  --grpc-gateway_out ./internal/pkg/api --grpc-gateway_opt paths=source_relative \
  --openapiv2_out ./internal/pkg/api \
  basket.proto

protoc --proto_path=./protos --go_out ./internal/pkg/api --go_opt paths=source_relative \
  --go-grpc_out ./internal/pkg/api --go-grpc_opt paths=source_relative \
  --grpc-gateway_out ./internal/pkg/api --grpc-gateway_opt paths=source_relative \
  --openapiv2_out ./internal/pkg/api \
  order.proto

protoc --proto_path=./protos --go_out ./internal/pkg/api --go_opt paths=source_relative \
  --go-grpc_out ./internal/pkg/api --go-grpc_opt paths=source_relative \
  --grpc-gateway_out ./internal/pkg/api --grpc-gateway_opt paths=source_relative \
  --openapiv2_out ./internal/pkg/api \
  payment.proto
