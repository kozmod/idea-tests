FROM golang:latest
LABEL stage=builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

FROM scratch
LABEL maintainer="Kozmo"
COPY --from=0 /usr/local/go/lib/time/zoneinfo.zip /
ENV TZ=Europe/Moscow
ENV ZONEINFO=/zoneinfo.zip
WORKDIR /app
COPY --from=0 /app/main .
EXPOSE 8080
CMD ["./main"]