#!/bin/bash

set -x

source ./hack/common.sh

#./hack/gen_certs.sh
generate_certificates

LOCALKUBECONFIG=${KUBECONFIG:-/var/run/kubernetes/admin.kubeconfig}
kubectl create secret tls myresource-validating-secret --cert tls.crt  --key tls.key

LOCALKUBECONFIG=${KUBECONFIG:-/var/run/kubernetes/admin.kubeconfig}
./cmd/controller-manager/controller-manager  --kubeconfig=${LOCALKUBECONFIG} -v=4 &


#####################
# Creation of service
#####################
kubectl create -f artefacts/crd-webhook-service.yaml

############################################################
# Creation of naked (don't try this in prod) pod for service
############################################################
kubectl create -f artefacts/crd-webhook-pod.yaml


generate_validating_admissionregistration
kubectl create -f artefacts/myresource-validating-admissionregistration.yaml
#rm -fr artefacts/myresource-validating-admissionregistration.yaml


generate_mutating_admissionregistration
kubectl create -f artefacts/myresource-mutating-admissionregistration.yaml
#rm -fr artefacts/myresource-mutating-admissionregistration.yaml

rm -f tls.crt tls.key
rm -f ca.crt ca.key


echo "Now you can build your pod"
echo "kubectl create -f ./hack/sample/myresource.yaml"
