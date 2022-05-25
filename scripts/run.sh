#!/bin/bash
export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')

url=$1
if [ -z "$url" ]
then
    #url="$(minikube ip):$INGRESS_PORT/customer"
    url="curl -X PUT $(minikube ip):$INGRESS_PORT/rios/establish -d '{"ani": "21981024950", "dnis": "0800994455"}'"
fi

while true
do curl $url
sleep .5
done