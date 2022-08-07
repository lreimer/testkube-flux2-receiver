# Testkube Flux2 Receiver

A notification receiver for Flux2 events to trigger Testkube test executions.

## Usage

```bash
# deploy and configure the Flux2 system
kubectl -n flux-system create secret generic testkube-flux2-url --from-literal=address=http://testkube-flux2-receiver.testkube.svc.cluster.local:8080/event
kubectl apply -k k8s/flux2/

# deploy the Testkube Flux2 receiver
kubectl -k k8s/
````

## Maintainer

M.-Leander Reimer (@lreimer), <mario-leander.reimer@qaware.de>

## License

This software is provided under the Apache v2.0 open source license, read the `LICENSE`
file for details.
