run: ## Run the application in docker
	docker compose -f ./src/docker-compose.yml up -d --build

run-watch: ## Run the docker watch mode
	docker compose -f ./src/docker-compose.yml up --watch

e2e-test: ## Run the e2e test
	"$(CURDIR)/src/frontend/scripts/e2e-testing.sh"
