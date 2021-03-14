# GCP 시작 

```
gcloud init
```
- 계정 선택 및 프로젝트, 클러스터 선택 

# GCP login 

```
gcloud auth login
```
- 단순 login만 했을 때 container registry에 push 할 때 문제가 생김

# GCP project 지정

```
gcloud config set project [project-id]
```

# kube config 항복 생성 

```
gcloud container cluster get-credentials [cluster name] --zone [zone 이름]
```




