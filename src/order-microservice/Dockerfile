FROM golang:1.20

RUN mkdir /app

WORKDIR /app

ADD src/order-microservice ./src/order-microservice
ADD src/order-microservice/go.mod ./src/order-microservice/.
ADD src/order-microservice/go.sum ./src/order-microservice/.

ADD src/shared ./src/shared

WORKDIR /app/src/order-microservice

RUN go mod download
RUN go build cmd/main.go

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

EXPOSE 8000

ENTRYPOINT /go/bin/CompileDaemon --build="go build cmd/main.go" --command=./main --color=true --polling=true 