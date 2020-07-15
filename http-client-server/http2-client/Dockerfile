FROM golang:latest
LABEL stage=builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build.client BIN_OUT=main GO_ENV='GO111MODULE="on" CGO_ENABLED=0 GOOS=linux'
#RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main/main.go

FROM alpine:latest
LABEL maintainer="Kozmo"
RUN apk add --no-cache bash
WORKDIR /app
COPY --from=0 /app/main .
#EXPOSE 8080
CMD ["./main"]