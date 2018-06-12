#!/usr/bin/env bash

CN=federation-v2-sig

function generate_certificates() {
#Refer to documentation (for example https://www.openssl.org/docs/manmaster/man5/x509v3_config.html or https://www.phildev.net/ssl/opensslconf.html )

ROOT_CA_KEY=ca.key
ROOT_CA_CERT=ca.crt

cat > ca.cfg<<EOF
[ req ]
default_bits       = 4096
default_md         = sha256
default_keyfile    = domain.com.key
prompt             = no
encrypt_key        = no
distinguished_name = req_distinguished_name
x509_extensions		= v3_ca
#req_extensions		= v3_req
# distinguished_name
[ req_distinguished_name ]
#countryName            = "FR"
#localityName           = "Sophia Antipolis"
#organizationName       = "Amadeus"
#organizationalUnitName = "TPE"
commonName             = ${CN}
#emailAddress           = "ilike@pizza.it"
[ v3_ca ]
#subjectKeyIdentifier=hash
#authorityKeyIdentifier=keyid:always,issuer:always
basicConstraints = critical, CA:true
keyUsage = critical, keyCertSign, digitalSignature, keyEncipherment
[ v3_req ]
# Lets at least make our requests PKIX complaint
subjectAltName=email:move
EOF

echo "Generate ${ROOT_CA_KEY} and ${ROOT_CA_CERT}"
openssl req -config ca.cfg -newkey rsa:2048 -nodes -keyout ${ROOT_CA_KEY} -x509 -days 36500 -out ${ROOT_CA_CERT}

openssl genrsa -out tls.key 2048

SERVICE=crd-webhook

cat > tls.cfg <<EOF
[req]
distinguished_name = req_distinguished_name
x509_extensions = v3_req
prompt = no

[req_distinguished_name]
CN = ${SERVICE}.default.svc

[v3_req]
keyUsage = critical, digitalSignature, keyEncipherment
extendedKeyUsage=serverAuth
#subjectAltName = @alt_names
#[alt_names]
#DNS.1 = ca.amadeus.io
EOF

SERVER_CRT=tls.crt
SERVER_KEY=tls.key

openssl req -new -key ${SERVER_KEY} -out tls.csr -config tls.cfg -batch -sha256

openssl x509 -req -days 36500 -in tls.csr -sha256 -CA ${ROOT_CA_CERT} -CAkey ${ROOT_CA_KEY} -CAcreateserial -out ${SERVER_CRT} -extensions v3_req -extfile tls.cfg

rm -rf *.cfg *.csr *.srl
}


function generate_mutating_admissionregistration() {
cat > artefacts/myresource-mutating-admissionregistration.yaml <<EOF
apiVersion: admissionregistration.k8s.io/v1beta1
kind: MutatingWebhookConfiguration
metadata:
  name: crd-mutating-wehbook
webhooks:
  - name: crd-mutating-wehbook.mygroup.amadeus.io
    clientConfig:
      service:
        name: crd-webhook
        namespace: default
        path: /mutate
      caBundle: ${CABUNDLE}
    failurePolicy: Fail
    rules:
      - apiGroups: ["mygroup.amadeus.io"]
        operations: [ "CREATE" ]
        apiVersions: ["v1alpha1"]
        resources: ["myresources"]
      - apiGroups: ["mygroup.amadeus.io"]
        operations: [ "UPDATE" ]
        apiVersions: ["v1alpha1"]
        resources: ["myresources"]
EOF
}

function generate_validating_admissionregistration() {
CABUNDLE=$(cat ca.crt | base64 | tr -d '\n')
cat > artefacts/myresource-validating-admissionregistration.yaml <<EOF
apiVersion: admissionregistration.k8s.io/v1beta1
kind: MutatingWebhookConfiguration
metadata:
  name: crd-validating-wehbook
webhooks:
  - name: crd-validating-wehbook.mygroup.amadeus.io
    clientConfig:
      service:
        name: crd-webhook
        namespace: default
        path: /validate
      caBundle: ${CABUNDLE}
    failurePolicy: Fail
    rules:
      - apiGroups: ["mygroup.amadeus.io"]
        operations: [ "CREATE" ]
        apiVersions: ["v1alpha1"]
        resources: ["myresources"]
      - apiGroups: ["mygroup.amadeus.io"]
        operations: [ "UPDATE" ]
        apiVersions: ["v1alpha1"]
        resources: ["myresources"]
EOF
}
