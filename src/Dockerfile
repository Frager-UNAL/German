FROM golang:1.12.0-alpine3.9

RUN apk add git

RUN mkdir /app
ADD . /app
WORKDIR /app

# Copy and download dependency

RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/mux
RUN go get github.com/joho/godotenv

# Build the application
RUN go build main.go

EXPOSE 19940
EXPOSE 3306

# Command to run when starting the container
CMD ["./main"]
