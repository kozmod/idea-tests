FROM ubuntu:latest
LABEL maintainer="Kozmo"

ARG app_http_port
ARG app_https_port
ARG app_file_storege
ARG cert_path
ARG key_path

WORKDIR /app
ARG ssl_config_path=./certs/san.conf
ARG root_cert_path=./certs/sec.syn.crt
ARG root_key_path=./certs/sec.syn.key
ARG csr_path=./certs/ott-ingress.csr

COPY ./bin/main .
COPY etc/san.conf ./certs/san.conf
COPY files/test.txt ./files/test.txt

SHELL ["/bin/bash", "-c"]
#WORKDIR /app/certs
RUN apt-get update && apt-get install -y curl &&  apt-get install -y vim
#RUN openssl req -x509 -sha256 -nodes -days 3365 -newkey rsa:2048 -subj '/O=SynSec/CN=sec.syn' -keyout sec.syn.key -out sec.syn.crt
#RUN openssl req -out ott-ingress.csr -newkey rsa:2048 -nodes -keyout ott-ingress.key -config san.conf  -subj "/C=RU/ST=MSC/L=MOSCOW/O=SBT/CN=localhost:$app_https_port"
#RUN openssl x509 -req -days 2365 -CA sec.syn.crt -CAkey sec.syn.key -set_serial 0 -in ott-ingress.csr -out ott-ingress.crt -extensions req_ext -extfile san.conf
RUN openssl req -x509 -sha256 -nodes -days 3365 -newkey rsa:2048 -subj '/O=SynSec/CN=sec.syn' -keyout $root_key_path -out $root_cert_path
RUN openssl req -out $csr_path -newkey rsa:2048 -nodes -keyout $key_path -config $ssl_config_path  -subj "/C=RU/ST=MSC/L=MOSCOW/O=SBT/CN=localhost"
RUN openssl x509 -req -days 2365 -CA $root_cert_path -CAkey $root_key_path -set_serial 0 -in $csr_path -out $cert_path -extensions req_ext -extfile $ssl_config_path

#WORKDIR /app
ENV FILE_STORAGE=$app_file_storege
ENV HTTP_PORT=$app_http_port
ENV HTTPS_PORT=$app_https_port
ENV CERT_PATH=$cert_path
ENV KEY_PATH=$key_path
EXPOSE $app_http_port
EXPOSE $app_https_port
CMD ["./main"]