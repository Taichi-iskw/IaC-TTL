#!/usr/bin/env node
import * as cdk from "aws-cdk-lib";
import { CdkStack } from "../lib/cdk-stack";

const projectName = "iac-ttl";

const app = new cdk.App();
new CdkStack(app, `${projectName}-stack`, {
	projectName,
	env: {
		account: process.env.CDK_DEFAULT_ACCOUNT,
		region: process.env.CDK_DEFAULT_REGION,
	},
});

cdk.Tags.of(app).add("Project", projectName);
