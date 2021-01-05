#!/bin/sh

docker build -t perevozov/arch:3 .

kubectl delete -f kube/deployment.yaml -n myapp

sleep 5

kubectl create -f kube/deployment.yaml -n myapp
