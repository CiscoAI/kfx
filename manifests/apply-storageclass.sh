#!/bin/bash

kubectl apply -f /tmp/local-path-storage.yaml
kubectl patch sc local-path -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'
