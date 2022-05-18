* Build Container: docker build -t sixbell/rios:v1 .
* Run standalone container: docker run -d -p 9080:9080 --name RiOS  sixbell/rios:v1
* Run redis + api: docker-compose up

for Kubernetes

docker build -t sixbell/rios:v1 .
