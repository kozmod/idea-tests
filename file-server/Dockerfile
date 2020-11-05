FROM ubuntu:latest
LABEL maintainer="Kozmo"
WORKDIR /app
COPY ./bin/main .
COPY ./certs/* ./certs/
ARG app_port
ARG app_file_storege
ARG cert_path
ARG key_path
ENV FILE_STORAGE=$app_file_storege
ENV PORT=$app_port
ENV CERT_PATH=$cert_path
ENV KEY_PATH=$key_path
EXPOSE $app_port
CMD ["./main"]