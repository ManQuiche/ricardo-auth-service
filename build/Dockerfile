FROM golang:1.16.8-bullseye
WORKDIR /app

COPY ../go.mod ./
COPY ../go.sum ./
RUN go mod download

COPY ./* ./
RUN go build -o /app/auth-service

ENV URL="locahost"
ENV PORT="80"

EXPOSE 80

RUN chmod +x /app/auth-service

CMD ["/app/auth-service"]