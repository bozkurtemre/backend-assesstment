run: ## Run the app locally
	docker compose -f ./src/docker-compose.yml up -d --build

run-watch: ## Run the docker watch mode
	docker compose -f ./src/docker-compose.yml up --watch
