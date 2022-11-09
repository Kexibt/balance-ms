docker-build:
	docker build --pull --rm -f "database/Dockerfile" -t psql_ms:latest "database"
	docker build --pull --rm -f "Dockerfile" -t app_ms:latest "."
	docker compose -f "docker-compose.yml" up -d --build