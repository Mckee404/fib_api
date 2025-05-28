# デプロイ
npx cdk deploy --context stage=[stage]

上記のコマンドを実行することでstageごとに分離した環境を作成できます。

## スタック
  -  API Gateway
  -  Lambda
  -  ECR

デプロイの際にbackendディレクトリ内のDockerfileがビルドされたイメージが自動で作成されたECRのリポジトリにプッシュされ、API GatewayによってLambda関数が実行される際にはECRにプッシュされたコンテナイメージが実行されます。
