FROM golang:latest
LABEL stage=builder
ARG cert_location=/usr/local/share/ca-certificates
RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.com.crt
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/proxy.golang.org.crt
RUN update-ca-certificates
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