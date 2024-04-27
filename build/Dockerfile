FROM golang:1.19-alpine3.15
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN GO111MODULE="" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/auth-service cmd/main.go

ENV URL="0.0.0.0"
ENV PORT="80"

EXPOSE 80

RUN chmod +x /app/auth-service

CMD ["/app/auth-service"]