# k8s 

TODO

- [ ] CI - Github Actions 

- [ ] CD - ArgoCD


appはGoのTodo API

localhost →ポードフォワード→ [Goコンテナ(Pod):8080] → [MySQLコンテナ(Pod):3306]

## k8s適用手順

goのDockerfileをビルドしておく

```
docker build ./api -t goapi:1.0
```


### 1. DB
1. DB用NameSpace作成
```
kubectl apply -f mysql-namespace.yaml
```


2. MySQLデプロイ
```
kubectl apply -f mysql-deployment.yaml -n database
```

3. サービス作成
```
kubectl apply -f mysql-service.yaml -n database
```

### 2. API

1. API用NameSpace作成
```
kubectl apply -f api-namespace.yaml 
```

2. ConfigMap, Secret作成. APIデプロイ
```
kubectl apply -f api-deployment.yaml,api-configmap.yaml,api-secret.yaml -n api
```

3. サービス作成
```
kubectl apply -f api-service.yaml -n api
```

### 3. 接続
1. ポートフォワード
```
kubectl port-forward svc/api 8080:80 -n api
```

2. ブラウザでアクセスまたはcurl
```
$ curl http://localhost:8080
```

### 4. リソース削除
1. API削除
```
kubectl delete -f api-deployment.yaml,api-configmap.yaml,api-secret.yaml,api-service.yaml -n api
```

2. DB削除
```
kubectl delete -f mysql-deployment.yaml,mysql-service.yaml -n database
```

3. NameSpace削除
```
kubectl delete -f api-namespace.yaml,mysql-namespace.yaml
```