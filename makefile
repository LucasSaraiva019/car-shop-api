init:
	swagger init spec --title "A Shop application" \
	  --description "Shop application" \
	  --version 1.0.0 \
	  --scheme http \
	  --consumes application/json \
	  --produces application/json
generate:
	swagger generate server --name shop-car-api --spec ./swagger/swagger.yml --server-package=gen/restapi  --model-package=gen/models --api-package=operations
run:
	go run ./cmd/shop-car-api-server/main.go --port 8000
