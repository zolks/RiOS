#!/bin/bash
export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')

while true
do curl -X PUT $(minikube ip):$INGRESS_PORT/sigas/establish -H "Content-Type: application/json" -H "Accept: application/json" -d "{\"ani\": \"21981024950\", \"dnis\": \"0800994455\"}"
sleep .5
done