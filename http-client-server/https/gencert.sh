OUTPUT=./etc/certs

CERT_NAME=${1:-test}

SSL_CONFIG_PATH=./cert.conf
ROOT_CERT_PATH=${OUTPUT}/${CERT_NAME}_root.crt
ROOT_KEY_PATH=${OUTPUT}/${CERT_NAME}_root.key
KEY_PATH=${OUTPUT}/${CERT_NAME}.key
CSR_PATH=${OUTPUT}/${CERT_NAME}.csr
CERT_PATH=${OUTPUT}/${CERT_NAME}.crt

mkdir -p ${OUTPUT}

openssl req -newkey rsa:2048 -nodes \
  -x509 -sha256 \
  -days 3650  \
  -subj "/O=SynSec/CN=${CERT_NAME}.syn" \
  -keyout "${ROOT_KEY_PATH}" \
  -out "${ROOT_CERT_PATH}"
openssl req -newkey rsa:2048 -nodes \
  -subj "/C=RU/ST=MSC/L=MOSCOW/O=SBT/CN=localhost" \
  -config ${SSL_CONFIG_PATH} \
  -keyout "${KEY_PATH}" \
  -out "${CSR_PATH}"
openssl x509 -req -days 3649 \
  -CA "${ROOT_CERT_PATH}" -CAkey "$ROOT_KEY_PATH" \
  -set_serial 0 \
  -in "${CSR_PATH}" \
  -out "${CERT_PATH}" \
  -extensions req_ext \
  -extfile ${SSL_CONFIG_PATH}

rm "${CSR_PATH}"