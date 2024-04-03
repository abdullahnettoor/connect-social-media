GOCMD=go

run: ## Start application
	$(GOCMD) run cmd/main.go

deps: ## Install dependencies
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	$(GOCMD) get -u -t -d -v ./...
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor

swag: ## Generate swagger docs
	swag init -g internal/infrastructure/api/server.go   
# 	swag init -g pkg/api/handler/admin.go -o ./cmd/api/docs # -o is to define the output location of swagger docs folder

nodemon:
	nodemon --exec $(GOCMD) run cmd/main.go --signal SIGTERM

consumer:
	$(GOCMD) run internal/infrastructure/kafka/consumer/consumer.go $(id)