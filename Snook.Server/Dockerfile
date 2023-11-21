FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify 

COPY *.go ./

EXPOSE 8080

RUN CGO_ENABLED=0 GOOS=linux go build -o /snook-server-app


CMD ["/snook-server-app"]