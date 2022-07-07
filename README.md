# k8s 

appはGoのTodo API

localhost →ポードフォワード→ [Goコンテナ(Pod):8080] → [MySQLコンテナ(Pod):3306]

## k8s適用手順(ローカル DockerDesktop)

goのDockerfileをビルドしておく

```
docker build ./app -t goapi:1.0
```


### 1. DB
1. DB用NameSpace作成
```
kubectl apply -f namespace/mysql-namespace.yaml
```


2. MySQLデプロイ
```
kubectl apply -f mysql/mysql-deployment.yaml -n database
```

3. サービス作成
```
kubectl apply -f mysql/mysql-service.yaml -n database
```

### 2. API

1. API用NameSpace作成
```
kubectl apply -f namespace/api-namespace.yaml 
```

2. ConfigMap, Secret作成. APIデプロイ
```
kubectl apply -f api-deployment-local.yaml,api-configmap-local.yaml,api/api-secret.yaml -n api
```

3. サービス作成
```
kubectl apply -f api/api-service.yaml -n api
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
kubectl delete -f api-deployment-local.yaml,api-configmap-local.yaml,api/api-secret.yaml,api/api-service.yaml -n api
```

2. DB削除
```
kubectl delete -f mysql/mysql-deployment.yaml,mysql/mysql-service.yaml -n database
```

3. NameSpace削除
```
kubectl delete -f namespace/api-namespace.yaml,namespace/mysql-namespace.yaml
```

## CI
Gibhub Actionでプルリク時に新しいAPIイメージの生成。マニフェスト変更のPR作成


## CD - ArgoCD -

### 1. ArgoCDインストール
ネームスペース作成 インストール
```
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

初期ログインパスワード取得
```
kubectl get secret argocd-initial-admin-secret -n argocd -o jsonpath='{.data.password}' | base64 --decode
```

ArgoCDポートフォワード
```
kubectl -n argocd port-forward service/argocd-server 30080:80
```

ブラウザでログイン
user: admin
pass: secretから取得した初期パスワード

### 2. applicationデプロイ
AppProject作成
```
kubectl apply -f argocd/argocd-appproject-test.yaml
```

Application適用
```
kubectl apply -f argocd/argocd-application-mysql.yaml
kubectl apply -f argocd/argocd-application-sample-app.yaml
```

APIをポートフォワード
```
kubectl  port-forward service/api 8080:80 
```

### 3. 削除
```
kubectl delete -f argocd/argocd-application-sample-app.yaml
kubectl delete -f argocd/argocd-application-mysql.yaml
kubectl delete -f argocd/argocd-appproject-test.yaml
kubectl delete -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
kubectl delete namespace argocd
```