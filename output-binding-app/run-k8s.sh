#!/usr/bin/env bash

docker build \
    -t nachtmaar/eventhubs-output-binding \
    --no-cache \
    .
docker push nachtmaar/eventhubs-output-binding
kubectl delete -f components/
kubectl apply -f components/

sleep 10

kubectl wait \
    --for condition=Available \
    --for condition=Progressing \
    deployment \
    -l app=goapp-output-binding

kubectl logs \
    -l app=goapp-output-binding \
    -c goapp-output-binding \
    -f

