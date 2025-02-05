# Parca

```
kubectl create namespace parca
```

```
kubectl apply -f https://github.com/parca-dev/parca/releases/download/v0.22.0/kubernetes-manifest.yaml
kubectl apply -f https://github.com/parca-dev/parca-agent/releases/download/v0.36.0/kubernetes-manifest.yaml
```

```
kubectl -n parca port-forward service/parca 7070
```
```
kubectl -n parca port-forward $(kubectl -n parca get pod -lapp.kubernetes.io/name=parca-agent -ojsonpath="{.items[0].metadata.name}") 7071
```

http://localhost:7070
http://localhost:7071

```
kubectl delete -f https://github.com/parca-dev/parca-agent/releases/download/v0.36.0/kubernetes-manifest.yaml
kubectl delete -f https://github.com/parca-dev/parca/releases/download/v0.22.0/kubernetes-manifest.yaml
```

```
kubectl delete namespace parca
```
