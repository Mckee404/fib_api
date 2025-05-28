import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import * as apigateway from "aws-cdk-lib/aws-apigateway";

export interface Props extends cdk.StackProps {
	stage: string;
}

export class FibApiStack extends cdk.Stack {
	public readonly fibApi: cdk.aws_lambda.Function;

	constructor(scope: Construct, id: string, props: Props) {
		super(scope, id, props);

		let name: string;
		if (props.stage === "prod") {
			name = "fib_api";
		} else {
			name = `fib_api_${props.stage}`;
		}

		this.fibApi = new cdk.aws_lambda.DockerImageFunction(this, name, {
			functionName: name,
			reservedConcurrentExecutions: 3,
			memorySize: 128,
			timeout: cdk.Duration.seconds(30),
			code: cdk.aws_lambda.DockerImageCode.fromImageAsset("../backend", {
				file: "Dockerfile",
				buildArgs: {
					ARCH: "arm64",
				},
			}),
			architecture: cdk.aws_lambda.Architecture.ARM_64,
		});

		// API Gatewayの作成
		const api = new apigateway.LambdaRestApi(this, "FibApiGateway", {
			handler: this.fibApi,
			proxy: false,
			restApiName: `${name}-gateway`,
			description: "Fibonacci API Gateway",
			deployOptions: {
				stageName: props.stage, // ステージ名をprops.stageで指定
			},
		});

		// /fib パスを追加
		const fib = api.root.addResource("fib");
		fib.addMethod("GET");
	}
}
