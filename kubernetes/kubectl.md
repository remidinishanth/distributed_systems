```shell
kubectl < get | delete | create >   < resource_type>
```

Example:

* Getting all the pods `kubectl get pods`
* Getting specific pod `kubectl get pod my-pod`
* Creation, you can also use apply command `kubectl < create | apply >  -f ./my-pod-config.yml`
* Port-forwarding `kubectl port-forward my-pod 8080:80`
