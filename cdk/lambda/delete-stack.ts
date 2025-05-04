import {
	CloudFormationClient,
	DeleteStackCommand,
} from "@aws-sdk/client-cloudformation";
import {
	SchedulerClient,
	DeleteScheduleCommand,
} from "@aws-sdk/client-scheduler";
import { Handler } from "aws-lambda";

interface Event {
	StackName: string;
}

export const handler: Handler<Event, void> = async (event) => {
	if (!event.StackName) throw new Error("StackName not provided");
	const cf = new CloudFormationClient({});
	await cf.send(new DeleteStackCommand({ StackName: event.StackName }));
	console.log(`DeleteStack triggered for ${event.StackName}`);

	await deleteScheduler(event.StackName);
};

async function deleteScheduler(stackName: string) {
	const client = new SchedulerClient({});
	const scheduleName = `ttl-delete-${stackName}`;
	try {
		await client.send(
			new DeleteScheduleCommand({
				Name: scheduleName,
				GroupName: "iac-ttl",
			})
		);
		console.log(`Deleted schedule: ${scheduleName}`);
	} catch (err) {
		console.error("Failed to delete schedule:", err);
	}
}
