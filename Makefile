BINARY_NAME=bot
PACKAGE_NAME=github.com/kappaprideonly/ege_bot_2.0

build:
	go build -v -o $(BINARY_NAME)

run:
	/$(BINARY_NAME)

clean:
	rm -f ${BINARY_NAME}

build-docker:
	sudo docker build -t $(BINARY_NAME) .

run-docker-release:
	sudo docker run --name $(BINARY_NAME) --rm -d --env-file config/./.env $(BINARY_NAME)

run-docker:
	sudo docker run --name $(BINARY_NAME) --rm --env-file config/./.env $(BINARY_NAME)

stop-docker:
	sudo docker stop $(BINARY_NAME)

clean-docker:
	sudo docker rm $(BINARY_NAME)

init:
	go mod init $(PACKAGE_NAME) 

tidy:
	go mod tidy
