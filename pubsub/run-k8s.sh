#!/usr/bin/env zsh

envsubst < components_old/* | cat
envsubst < components_old/* | ko delete -f -
envsubst < components_old/* | ko apply -f -

sleep 10

kubectl wait \
    --for condition=Available \
    --for condition=Progressing \
    deployment \
    -l app=goapp-input-binding

stern \
    -l app=goapp-input-binding
