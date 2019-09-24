#!/bin/bash

kubectl create ns kubeflow
kubectl create ns kubeflow-anonymous
kubectl apply -f /tmp/mla-deploy.yaml -n kubeflow
