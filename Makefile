proto:
	protoc pkg/todo/proto/*.proto --go_out=pkg/todo/pb --go-grpc_out=pkg/todo/pb
	protoc pkg/players/proto/*.proto --go_out=pkg/players/pb --go-grpc_out=pkg/players/pb
	protoc pkg/teams/proto/*.proto --go_out=pkg/teams/pb --go-grpc_out=pkg/teams/pb

clean:
    rm pkg/todo/pb/*.go
	rm pkg/players/pb/*.go
	rm pkg/teams/pb/*.go