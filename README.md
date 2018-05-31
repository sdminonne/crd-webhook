#Webhook and CRD experimentation

This document contains some experimentations to better understand behaviors for CRD controller and wehbooks. Rationale: today 1.11 validation and defaulting for CRD is granted via OpenAPI schema, but Turing complete validation and complex defaulting (as far as I understand defaulting with cross checks: if fieldA is X then fieldB must be Y) needs webhooks.

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

PR in istio pilot (the old one apprently) to implment webhook validation: https://github.com/istio/old_pilot_repo/pull/1158 now https://github.com/istio/istio
