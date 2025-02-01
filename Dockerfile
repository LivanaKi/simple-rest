FROM golang:alpine

WORKDIR /app
COPY . .

# install psql
RUN apk --update add postgresql-client

RUN go mod download
RUN go build -o simple-rest ./main.go
EXPOSE 8080
CMD ["./simple-rest"]