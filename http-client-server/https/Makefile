
CERT_LOCATION :=$(realpath ./etc/certs)

check.crt:
	openssl x509 -text -in ${CERT_LOCATION}/${CERT_NAME}.crt

gen:
	./gencert.sh

clear.crt:
	rm -rf ${CERT_LOCATION}
	mkdir -p ${CERT_LOCATION}