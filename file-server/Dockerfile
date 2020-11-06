FROM ubuntu:latest
LABEL maintainer="Kozmo"

ARG app_http_port
ARG app_https_port
ARG app_file_storege
ARG cert_path
ARG key_path
ARG ssl_config_path=./certs/san.conf
ARG root_cert_path=./certs/sec.syn.crt
ARG root_key_path=./certs/sec.syn.key
ARG crs_path=./certs/ott-ingress.csr

WORKDIR /app
COPY ./bin/main .
COPY etc/san.conf ./certs/san.conf

SHELL ["/bin/bash", "-c"]
RUN apt-get update && apt-get install -y curl &&  apt-get install -y vim
RUN openssl req -x509 -sha256 -nodes -days 3365 -newkey rsa:2048 -subj '/O=SynSec/CN=sec.syn' -keyout $root_key_path -out $root_cert_path
RUN openssl req -out $crs_path -newkey rsa:2048 -nodes -keyout $cert_path -config $ssl_config_path  -subj "/C=RU/ST=MSC/L=MOSCOW/O=SBT/CN=localhost:$app_https_port"
RUN openssl x509 -req -days 2365 -CA $root_cert_path -CAkey $root_key_path -set_serial 0 -in $crs_path -out $cert_path -extensions req_ext -extfile $ssl_config_path

ENV FILE_STORAGE=$app_file_storege
ENV HTTP_PORT=$app_http_port
ENV HTTPS_PORT=$app_https_port
ENV CERT_PATH=$cert_path
ENV KEY_PATH=$key_path
EXPOSE $app_http_port
EXPOSE $app_https_port
CMD ["./main"]