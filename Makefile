BINARY_NAME=bot
PACKAGE_NAME=github.com/kappaprideonly/ege_bot_2.0

run:
	go run app.go

clean:
	rm -f ${BINARY_NAME}

build-docker:
	sudo docker build -t $(BINARY_NAME) .

run-docker-release:
	sudo docker run --name $(BINARY_NAME) -d $(BINARY_NAME)

run-docker:
	sudo docker run --name $(BINARY_NAME)

clean-docker:
	sudo docker remove $(BINARY_NAME)

stop-docker:
	sudo docker stop $(BINARY_NAME)

tidy:
	go mod tidy
