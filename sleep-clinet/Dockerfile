FROM golang:latest
LABEL stage=builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN  GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
LABEL maintainer="Kozmo"
RUN apk update && apk add mc tmux curl busybox-extras openssh-client apache2-utils vim
WORKDIR /app
COPY --from=0 /app/main .
EXPOSE 8080
CMD ["./main"]