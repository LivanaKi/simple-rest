FROM golang:alpine
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o simple-rest ./main.go
EXPOSE 8080
CMD ["./simple-rest"]