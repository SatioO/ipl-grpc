proto:
	protoc pkg/todo/proto/*.proto --go_out=pkg/todo/pb --go-grpc_out=pkg/todo/pb

clean:
    rm pkg/todo/pb/*.go