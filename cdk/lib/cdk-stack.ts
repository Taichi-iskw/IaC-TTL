import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import * as lambda from "aws-cdk-lib/aws-lambda";
import * as iam from "aws-cdk-lib/aws-iam";
import * as lambdaNode from "aws-cdk-lib/aws-lambda-nodejs";
import * as scheduler from "aws-cdk-lib/aws-scheduler";
import * as ssm from "aws-cdk-lib/aws-ssm";

export class CdkStack extends cdk.Stack {
	constructor(scope: Construct, id: string, props?: cdk.StackProps) {
		super(scope, id, props);

		// delete lambda
		const destroyFn = new lambdaNode.NodejsFunction(
			this,
			createName("DestroyExecFn"),
			{
				runtime: lambda.Runtime.NODEJS_22_X,
				entry: "lambda/delete-stack.ts",
				handler: "handler",
				timeout: cdk.Duration.minutes(5),
				bundling: { format: lambdaNode.OutputFormat.ESM },
			}
		);

		destroyFn.addToRolePolicy(
			new iam.PolicyStatement({
				actions: ["cloudformation:DeleteStack"],
				resources: ["*"],
			})
		);
		destroyFn.addToRolePolicy(
			new iam.PolicyStatement({
				actions: ["scheduler:DeleteSchedule"],
				resources: [
					`arn:aws:scheduler:${this.region}:${this.account}:schedule/iac-ttl/*`,
				],
			})
		);

		// scheduler invoke role
		const schedulerInvokeRole = new iam.Role(
			this,
			createName("SchedulerInvokeRole"),
			{
				assumedBy: new iam.ServicePrincipal("scheduler.amazonaws.com"),
				description:
					"Allows EventBridge Scheduler to invoke DestroyExecFn",
			}
		);

		schedulerInvokeRole.addToPolicy(
			new iam.PolicyStatement({
				actions: ["lambda:InvokeFunction"],
				resources: [destroyFn.functionArn],
			})
		);

		// scheduler group
		new scheduler.CfnScheduleGroup(this, createName("SchedulerGroup"), {
			name: "iac-ttl",
		});

		// ssm parameter
		new ssm.StringParameter(this, createName("DestroyFunctionArnParam"), {
			parameterName: "/iac-ttl/destroy-fn-arn",
			stringValue: destroyFn.functionArn,
		});

		new ssm.StringParameter(this, createName("SchedulerRoleArnParam"), {
			parameterName: "/iac-ttl/scheduler-role-arn",
			stringValue: schedulerInvokeRole.roleArn,
		});
	}
}

function createName(name: string) {
	const projectName = "iac-ttl";
	return `${projectName}-${name}`;
}
