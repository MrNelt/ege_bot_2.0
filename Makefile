BINARY_NAME=bot
PACKAGE_NAME=github.com/kappaprideonly/ege_bot_2.0

run:
	go run app.go

build-service:
	sudo docker-compose build

run-service:
	sudo docker-compose up -d $(BINARY_NAME)

stop-service:
	sudo docker-compose stop

destroy-service:
	sudo docker-compose down
