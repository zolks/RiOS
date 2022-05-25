* Build Container: docker build -t sixbell/rios:v1 .
* Run standalone container: docker run -d -p 9080:9080 --name RiOS  sixbell/rios:v1
* Run redis + api: docker-compose up

--- for Kubernetes
kubectl create namespace ri
kubectl config set-context $(kubectl config current-context) --namespace=ri

kubectl apply -f kubernetes/redis-master.yml 
or
kubectl apply -f <(istioctl kube-inject -f kubernetes/redis-deployment.yml) -n ri
kubectl create -f kubernetes/redis-service.yml -n ri


eval $(minikube docker-env)
docker build -t sixbell/rios:v1 .

kubectl apply -f <(istioctl kube-inject -f kubernetes/Deployment.yml) -n ri
kubectl create -f kubernetes/Service.yml -n ri


kubectl get pods -w -n ri


kubectl apply -f kubernetes/Gateway.yml -n ri
or
kubectl create -f kubernetes/Gateway.yml -n ri

update -> kubectl rollout restart deployment --selector=app=rios

------- TEST 
export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')
curl -X PUT $(minikube ip):$INGRESS_PORT/rios/establish -d '{"ani": "21981024950", "dnis": "0800994455"}'
curl -X PUT $(minikube ip):$INGRESS_PORT/rios/registerFlow -d '{"id":1001,"service_number":"0800994455","name":"FluxoTeste","start_node":100,"nodes":{"100":{"id":100,"type":"Start","name":"Início","welcome_message":"Vocêligouparaacentralvivo","next_node_id":101},"101":{"id":101,"type":"Callcenter","name":"CC_01","active":true,"cc_number":"0800-998787","default_node_id":103,"error_node_id":102},"102":{"id":102,"type":"End","name":"EndError","end_cause":"504"},"103":{"id":103,"type":"End","name":"EndOK","end_cause":"200"}}}'


---- DELETE
delete all -> kubectl delete all --all -n ri
delete pod + replica set -> kubectl delete rs -l app=sigas,version=v1
                            kubectl delete deployment -l app=sigas,version=v1   

---- DARK Launch
kubectl create -f kubernetes/destination-rule-v1-v2.yml -n ri
kubectl create -f kubernetes/virtual-service-v1-mirror-v2.yml -n ri
kubectl logs rios-v2-7fd74d8f94-sws85 -c rios -n ri -f