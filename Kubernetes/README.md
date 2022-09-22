#K8s Imerpative Commends 

## Pod

```sh
kubectl run nginx --image=nginx 
```

```sh
kubectl create nginx --image=nginx -o yaml
```

## Deployment

```sh
kubectl create deploy nginx --image=nginx -o yml
```

### with replicas

```sh
kubectl create deploy nginx --image=nginx --replicas=4 
```

### editing replicas

```sh
kubectl scale deploy nginx --replicas=5
```

### create deploy and drop to yml

```sh
kubectl create deployment nginx --image=nginx --dry-run=client -o yaml > nginx-deployment.yaml
```

- --dry-run=client 는 반영 하지 않고 file만 드랍

## Service

```sh
kubectl expose pod redis --port=6379 --name redis-service --dry-run=client -o yaml
```

### with expose 

```sh
kubectl expose pod nginx --type=NodePort --port=80 --name=nginx-service --dry-run=client -o yaml
```



# GCP 

## GCP 시작 

```
gcloud init
```
- 계정 선택 및 프로젝트, 클러스터 선택 

## GCP login 

```
gcloud auth login
```
- 단순 login만 했을 때 container registry에 push 할 때 문제가 생김 

## GCP project 지정

```
gcloud config set project [project-id]
```

## kube config 항복 생성 

```
gcloud container clusters get-credentials [cluster name] --zone [zone 이름]
```

# Kubernetes Dashboard 

## Dashboard 설치 
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-beta1/aio/deploy/recommended.yaml 
```

## kubectl proxy로 kubernetes dashboard 실행 
```
kubectl proxy --port=8001 &
```

### kubernetes login 인증키 생성 
```
cat <<EOF | kubectl create -f -
 apiVersion: v1
 kind: ServiceAccount
 metadata:
   name: admin-user
   namespace: kube-system
EOF
```

### Cluster Role binding 
```
cat <<EOF | kubectl create -f -
 apiVersion: rbac.authorization.k8s.io/v1
 kind: ClusterRoleBinding
 metadata:
   name: admin-user
 roleRef:
   apiGroup: rbac.authorization.k8s.io
   kind: ClusterRole
   name: cluster-admin
 subjects:
 - kind: ServiceAccount
   name: admin-user
   namespace: kube-system
EOF
```

### Token 확인 
```
kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep admin-user | awk '{print $1}') 
```
- http://127.0.0.1:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login
- 위 링크로 접속 후 Token 복사 후 dashboard 로그인! 기무리

# Model Serving 

## model output structure 확인

```
infer = loaded.signatures["serving_default"]
print(infer.structured_outputs)
```
- 출력문에 output 형태를 확인 할 수 있음. 




```
