FROM golang:latest
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /redis-get-test
EXPOSE 8081
CMD ["/redis-get-test"]