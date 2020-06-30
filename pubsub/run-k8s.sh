#!/usr/bin/env zsh

docker build \
    -t nachtmaar/az-eh-pubsub \
    --no-cache \
    .
docker push nachtmaar/az-eh-pubsub

envsubst < components_old/* | cat
envsubst < components_old/* | kubectl delete -f -
envsubst < components_old/* | kubectl apply -f -

sleep 10

kubectl wait \
    --for condition=Available \
    --for condition=Progressing \
    deployment \
    -l app=goapp-input-binding

stern \
    -l app=goapp-input-binding
