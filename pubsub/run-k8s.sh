#!/usr/bin/env zsh

envsubst < components/* | cat
envsubst < components/* | ko delete -f -
envsubst < components/* | ko apply -f -

sleep 10

kubectl wait \
    --for condition=Available \
    --for condition=Progressing \
    deployment \
    -l app=goapp-pubsub

stern \
    -l app=goapp-pubsub
