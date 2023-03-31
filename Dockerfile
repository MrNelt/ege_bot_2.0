FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY .env ./

COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /bot

EXPOSE 8080

CMD [ "/bot" ]