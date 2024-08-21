run: ## Run the application in docker
	docker compose -f ./src/docker-compose.yml up -d --build

run-watch: ## Run the docker watch mode
	docker compose -f ./src/docker-compose.yml up --watch

run-test: ## Run the test
	cd src/frontend && go test -v ./... && cd ../../src/worker && go test -v ./...

e2e-test: ## Run the e2e test
	"$(CURDIR)/src/frontend/scripts/e2e-testing.sh"
