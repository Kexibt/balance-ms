docker-build:
	docker build --pull --rm -f "database/Dockerfile" -t psql_ms:latest "database"
	docker build --pull --rm -f "Dockerfile" -t app_ms:latest "."
	docker compose -f "docker-compose.yml" up -d --build

run-no-docker:
	ls go.mod || go mod init github.com/Kexibt/balance-microservice
	go mod tidy
	go mod download
	echo "\\033[1;32m Если приложение не стартануло, то попробуйте следуюшее: \\033[0;39m"
	echo "\\033[1;32m Развернуть бд PostgreSQL, файл инициализации лежит в database/initfiles. \\033[0;39m"
	echo "\\033[1;32m После развертывания бд, запишите данные для подключения в cfg.yml(не пинайте за строку подключения, времени немного было c:) \\033[0;39m"
	echo "\\033[1;94m Не удаляйте /micro_balance из строки подключения!! \\033[0;39m"
	echo "\\033[1;32m Снова запустите эту команду \\033[0;39m"
	go run main.go