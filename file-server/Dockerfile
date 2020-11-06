FROM ubuntu:latest
LABEL maintainer="Kozmo"
SHELL ["/bin/bash", "-c"]
ARG app_port
ARG app_file_storege
#ARG root_cert_path
#ARG root_key_path
ARG cert_path
ARG key_path
WORKDIR /app
COPY ./bin/main .
COPY ./certs/san.conf ./certs/san.conf
RUN apt-get update && apt-get install -y curl &&  apt-get install -y vim
RUN openssl req -x509 -sha256 -nodes -days 3365 -newkey rsa:2048 -subj '/O=SynSec/CN=sec.syn' -keyout ./certs/sec.syn.key -out ./certs/sec.syn.crt
RUN openssl req -out ./certs/ott-ingress.csr -newkey rsa:2048 -nodes -keyout $cert_path -config ./certs/san.conf  -subj "/C=RU/ST=MSC/L=MOSCOW/O=SBT/CN=localhost:$app_port"
RUN openssl x509 -req -days 2365 -CA ./certs/sec.syn.crt -CAkey ./certs/sec.syn.key -set_serial 0 -in ./certs/ott-ingress.csr -out $cert_path -extensions req_ext -extfile ./certs/san.conf
ENV FILE_STORAGE=$app_file_storege
ENV PORT=$app_port
ENV CERT_PATH=$cert_path
ENV KEY_PATH=$key_path
EXPOSE $app_port
CMD ["./main"]