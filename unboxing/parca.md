# Parca

## Install

1. Create a namespace for Parca components to run in:
   ```
   kubectl create namespace parca
   ```
2. Provision Parca to use the API and UI:
   ```
   kubectl apply -f https://github.com/parca-dev/parca/releases/download/v0.24.0/kubernetes-manifest.yaml
   ```
3. To view the Parca UI and access the API, port-forward using the default port `7070`:
   ```
   kubectl -n parca port-forward service/parca 7070:7070
   ```
   Once the Parca is running, and you set up the port-forwarding you can navigate to the web interface by
   visiting http://localhost:7070.

   At this stage you won't see any data because you haven't configured any data sorouce. Thus, Parca UI hasn't digested any data to display.
4. Provision the Parca Agent as a `DaemonSet`:
   ```
   kubectl apply -f https://github.com/parca-dev/parca-agent/releases/download/v0.39.1/kubernetes-manifest.yaml
   ```
5. Set up port-forward using the default port `7071`:
   ```
   kubectl -n parca port-forward $(kubectl -n parca get pod -lapp.kubernetes.io/name=parca-agent -ojsonpath="{.items[0].metadata.name}") 7071:7071
   ```
   View active profiles by visiting http://localhost:7071.

## Uninstall

```
kubectl delete -f https://github.com/parca-dev/parca-agent/releases/download/v0.39.1/kubernetes-manifest.yaml
kubectl delete -f https://github.com/parca-dev/parca/releases/download/v0.24.0/kubernetes-manifest.yaml
```

```
kubectl delete namespace parca
```

* https://www.parca.dev/docs/kubernetes/
