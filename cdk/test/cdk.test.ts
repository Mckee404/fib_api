import * as cdk from "aws-cdk-lib";
import { Template } from "aws-cdk-lib/assertions";
import { FibApiStack } from "../lib/cdk-stack";

test("Lambda Function Created", () => {
	const app = new cdk.App();
	const stack = new FibApiStack(app, "TestStack", { stage: "dev" });
	const template = Template.fromStack(stack);

	template.hasResourceProperties("AWS::Lambda::Function", {
		FunctionName: "fib_api_dev",
		MemorySize: 128,
		Timeout: 30,
		Architectures: ["arm64"],
	});
});

test("API Gateway /fib GET exists", () => {
	const app = new cdk.App();
	const stack = new FibApiStack(app, "TestStack", { stage: "dev" });
	const template = Template.fromStack(stack);

	template.hasResourceProperties("AWS::ApiGateway::Method", {
		HttpMethod: "GET",
	});

	template.hasResourceProperties("AWS::ApiGateway::Resource", {
		PathPart: "fib",
	});
});
