# services command
start-gateway:
	clear && go run src/gateway-service/cmd/web/main.go

start-role:
	clear && go run src/role-service/cmd/web/main.go

start-product:
	clear && go run src/product-service/cmd/web/main.go

start-order:
	clear && go run src/order-service/cmd/web/main.go
# docker command
start-docker:
	clear && docker compose -f ./docker/docker-compose.yml up -d 

stop-docker:
	clear && docker compose -f ./docker/docker-compose.yml down --remove-orphans

clean-docker:
	clear && docker system prune && docker volume prune && docker image prune -a -f && docker container prune

generate-proto:
	clear && protoc --proto_path=grpc/proto grpc/proto/*.proto --go_out=grpc --go-grpc_out=grpc
