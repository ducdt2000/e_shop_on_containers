FROM golang:1.20

RUN mkdir /app

WORKDIR /app

ADD src/order-microservice/go.mod .
# ADD src/order-microservice/go.sum .

RUN go mod download
ADD src/order-microservice .

RUN go build cmd/main.go

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

EXPOSE 8000

ENTRYPOINT /go/bin/CompileDaemon --build="go build cmd/main.go" --command=./main