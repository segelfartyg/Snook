FROM golang:1.21

WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod verify 

COPY . ./

EXPOSE 8080

RUN CGO_ENABLED=0 GOOS=linux go build -o /webserver

CMD ["/webserver"]