#!/usr/bin/env bash

docker build \
    -t nachtmaar/eventhubs-input-binding \
    --no-cache \
    .
docker push nachtmaar/eventhubs-input-binding
kubectl delete -f components/
kubectl apply -f components/
sleep 10

kubectl wait \
    --for condition=Available \
    --for condition=Progressing \
    deployment \
    -l app=goapp-input-binding

stern \
    -l app=goapp-input-binding
