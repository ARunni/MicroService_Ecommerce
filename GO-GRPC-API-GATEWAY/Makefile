run:
	go run cmd/main.go

proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/admin/admin.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/user/user.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/product/product.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/cart/cart.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/order/order.proto