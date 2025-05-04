import {
	CloudFormationClient,
	DeleteStackCommand,
} from "@aws-sdk/client-cloudformation";
import { Handler } from "aws-lambda";

interface Event {
	StackName: string;
}

export const handler: Handler<Event, void> = async (event) => {
	if (!event.StackName) throw new Error("StackName not provided");
	const cf = new CloudFormationClient({});
	await cf.send(new DeleteStackCommand({ StackName: event.StackName }));
	console.log(`DeleteStack triggered for ${event.StackName}`);
};
