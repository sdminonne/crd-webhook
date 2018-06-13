#Webhook and CRD experimentation

This document contains some experimentations to better understand behaviors for CRD controller and wehbooks. Rationale: today 1.11 validation and defaulting for CRD is granted via OpenAPI schema, but Touring complete validation and complex defaulting (as far as I understand defaulting with cross checks: if fieldA is X then fieldB must be Y) needs webhooks.

In this experimentaton kubebuilder is used, at the time of writing (begin of June 2018) kubebuilder project an issue has been found: https://github.com/kubernetes-sigs/kubebuilder/issues/216 to ask for a WebHook Package to better implement a WebHook.

## Kicking the tires

```shell
$kubebuilder init --domain amadeus.io
```

the resource have been created via:

```shell
kubebuilder create resource --group mygroup --kind Myresource --version v1alpha1
```

## References

http://book.kubebuilder.io/

Good video from Stefan Schimanski (Red Hat) at KubeCon-EU 2018 https://www.youtube.com/watch?v=XsFH7OEIIvI

PR in istio pilot (the old one apprently) to implement webhook validation: https://github.com/istio/old_pilot_repo/pull/1158 now https://github.com/istio/istio

In the last hours I discovered a _library for writing admission webhooks_ https://github.com/openshift/generic-admission-server

and a full example:

https://github.com/GoogleCloudPlatform/agon/

# Setting up validating webhook


As reported here https://github.com/caesarxuchao/example-webhook-admission-controller
I followed Kubernetes e2e test for webhook setup https://github.com/kubernetes/kubernetes/blob/release-1.9/test/e2e/apimachinery/webhook.go and implementation
https://github.com/kubernetes/kubernetes/tree/release-1.9/test/images/webhook

First of all  `admissionregistration` should be enabled in your cluster.
You should not obtain an empty line. Instead you should obtain something like this:
```shell
$ kubectl api-versions | grep admissionregistration.k8s.io/v1beta1
admissionregistration.k8s.io/v1beta1
```




In the path will be supply:
`tls.key` and `tls.cert` which are the webhook cert/key pair and the `ca.crt` is the signing certificate
(to be supplied to the APIserver in caBundle during MutatingWebhook registration).


First of all let's start by
```shell
./hack/gen_certs.sh
```

and creating the Secret:

```shell
$ kubectl create secret tls myresource-validating-secret --cert tls.crt  --key tls.key  -o yaml
```


# Deployment

You must have a Kubernetes cluster configured.

First of all start the controller created with `kubebuilder`, locally it will register the CRD automatically, Running in cluster you must register the CRD.

Something like:

```shell
$ controller-manager --kubeconfig=${KUBECONFIG}
```

Now you can deploy:

```shell
$ kubectl create -f artefacts/myresource-validating-secret.yaml
```

```shell
$ kubectl create -f artefacts/myresource-validating-service.yaml
```

```shell
$ kubectl create -f artefacts/myresource-validating-pod.yaml
```

``shell
$ kubectl create -f artefacts/myresource-validating-admissionregistration.yaml
```
