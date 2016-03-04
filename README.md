# Hands on Kubernetes

Materials for my _Hands on Kubernetes_ session at UBS Coding Harbor meetup.

## What to bring?

To make the session run smoothly, please bring a laptop with:
- working Internet connection
- Git
- Docker
  - and run [pull-docker-images.sh](pull-docker-images.sh) to save on downloads during the session
- your favorite text editor
- (OPTIONAL) your favorite Docker-ready tech-stack for building services with some REST API

## The plan

- Intro
  - Docker
  - Kubernetes (k8s)
- Run Kubernetes locally via Docker
- Pods
- Replication Controllers
- Labels
- Scaling
- Services
  - Service discovery
  - Load balancing
  - Virtual IPs and service proxies
- Releasing
  - Rolling updates
  - Canary releases

## Cheat sheet

Install k8s:
```zsh
cd k8s-on-docker
docker-compose up -d
./tunnel-apiserver.sh  # For Docker on VM
# Install kubernetes-cli (kubectl). For Mac:
brew install kubernetes-cli  # brewi if you use Prezto
./enable-dns.sh
```

Play with k8s:
```zsh
kubectl get [node|pod|rc|svc]
kubectl describe type/name
docker build -t k8s-hands-on/products:v1 .
kubectl create -f def.yml
kubectl replace -f def.yml
kubectl delete type/name
docker exec -it httpie... bash  # Use kubectl exec on real cluster
kubectl scale rc/products --replicas=10
kubectl rolling-update products --image=k8s-hands-on/products:v2
kubectl rolling-update products --image=k8s-hands-on/products:v2 --update-period=1s
kubectl rolling-update products --image=k8s-hands-on/products:v2 --rollback
kubectl label pod -lsvc=products track=stable
```
