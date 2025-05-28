# フィボナッチ数列API
指定されたインデックスに対応したフィボナッチ数を返却するREST APIです。

## 技術スタック
  -  Go(Gin)
  -  Docker
  -  AWS CDK

## エンドポイント
https://aa54leni16.execute-api.ap-northeast-1.amazonaws.com/dev/fib?n=[index]

indexに9223372036854775807までの0以上の整数を渡すことでフィボナッチ数が返却されます。

## ローカルでの実行方法
### ビルド
docker buildx build --build-arg ARCH="[ARCH]" --target local -t fibonacci_api .

CPUのアーキテクチャに合わせてARCHにはamd64かarm64を選択してください。

### 実行
docker run -p 8080:8080 fibonacci_api /functions/fibonacci

### curlによる確認
curl -X POST "http://localhost:8080/2015-03-31/functions/function/invocations" -d '{"path": "/fib?n=99"}'

## デプロイ
npx cdk deploy --context stage=[stage]

上記のコマンドによりstageを分離させてデプロイできます。
