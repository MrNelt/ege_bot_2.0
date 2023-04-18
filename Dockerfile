FROM golang:1.20.3-alpine

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY . ./

RUN go build -o /bot

EXPOSE 8080

CMD [ "/bot" ]