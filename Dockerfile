FROM golang:1.18-alpine

RUN apk add git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o goservice

EXPOSE 8080

CMD ["/app/goservice"]