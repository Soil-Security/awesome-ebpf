# Unboxing Falco

## Install Falco and Falco Exporter

```
helm repo add falcosecurity https://falcosecurity.github.io/charts
helm repo update falcosecurity
```

```
helm install falco falcosecurity/falco \
  --namespace=falco \
  --create-namespace \
  --version=4.2.2 \
  --set="driver.kind=modern_ebpf" \
  --set="driver.modernEbpf.leastPrivileged=false" \
  --set="falco.grpc.enabled=true" \
  --set="falco.grpc_output.enabled=true" \
  --set="collectors.kubernetes.enabled=true"
```

```
helm install falco-exporter falcosecurity/falco-exporter \
  --namespace=falco
```

```
kubectl port-forward -n falco svc/falco-exporter 9376
```

<http://localhost:9376/metrics>

```
# HELP falco_events
# TYPE falco_events counter
falco_events{hostname="falco-5cjlj",k8s_ns_name="default",k8s_pod_name="nginx-85bfcd86d5-d4h45",priority="2",rule="Drop and execute new binary in container",source="syscall",tags=",PCI_DSS_11.5.1,TA0003,container,maturity_stable,mitre_persistence,process,"} 1
falco_events{hostname="falco-5cjlj",k8s_ns_name="default",k8s_pod_name="nginx-85bfcd86d5-d4h45",priority="5",rule="Terminal shell in container",source="syscall",tags=",T1059,container,maturity_stable,mitre_execution,shell,"} 1
falco_events{hostname="falco-5cjlj",k8s_ns_name="foo",k8s_pod_name="foo-569c65d497-sjf4d",priority="2",rule="Drop and execute new binary in container",source="syscall",tags=",PCI_DSS_11.5.1,TA0003,container,maturity_stable,mitre_persistence,process,"} 1
falco_events{hostname="falco-5cjlj",k8s_ns_name="foo",k8s_pod_name="foo-569c65d497-sjf4d",priority="5",rule="Terminal shell in container",source="syscall",tags=",T1059,container,maturity_stable,mitre_execution,shell,"} 1
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 2.4851e-05
go_gc_duration_seconds{quantile="0.25"} 3.244e-05
go_gc_duration_seconds{quantile="0.5"} 8.1696e-05
go_gc_duration_seconds{quantile="0.75"} 0.000140356
go_gc_duration_seconds{quantile="1"} 0.000284848
go_gc_duration_seconds_sum 0.00076428
go_gc_duration_seconds_count 7
[...]
```

```
helm uninstall -n falco falco-exporter
helm uninstall -n falco falco
```

## Trigger Suspect Actions that are Detected by Falco Rulesets

```
helm install event-generator falcosecurity/event-generator \
  --namespace event-generator \
  --create-namespace \
  --set="config.command=test" \
  --set="config.loop=false"
```

```
helm uninstall event-generator --namespace event-generator
```

## Extending Falco Using the gRPC API


[code/unboxing/falco/grpc-client/main.go](../code/unboxing/falco/grpc-client/main.go)


<https://github.com/falcosecurity/client-go>

## Resources

* <https://semaphoreci.com/blog/prometheus-grafana-kubernetes-helm>
