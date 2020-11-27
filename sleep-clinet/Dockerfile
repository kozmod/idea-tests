FROM golang:latest
LABEL stage=builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN  GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -o main .

FROM ubuntu:latest
LABEL maintainer="Kozmo"
RUN apt-get update && apt-get -y install curl
RUN apt-get update && apt-get -y install telnet
WORKDIR /app
COPY --from=0 /app/main .
EXPOSE 8080
CMD ["./main"]