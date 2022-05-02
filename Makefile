proto:
	protoc pkg/todo/proto/*.proto --go_out=pkg/todo/pb --go-grpc_out=pkg/todo/pb
	protoc pkg/players/proto/*.proto --go_out=pkg/players/pb --go-grpc_out=pkg/players/pb

clean:
    rm pkg/todo/pb/*.go
	rm pkg/players/pb/*.go