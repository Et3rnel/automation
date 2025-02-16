package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/pubsub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func provisionTopic(ctx *pulumi.Context) (*pubsub.Topic, error) {
	// Create a new Pub/Sub topic with the Pulumi resource name "gmailTopic" used internally to identify the resource
	topic, err := pubsub.NewTopic(ctx, "gmailTopic", &pubsub.TopicArgs{
		// Set the actual name of the topic to "gmail-notifications"
		Name: pulumi.String("gmail-notifications"),
	})

	if err != nil {
		return nil, err
	}

	// Export the topic name so that it can be referenced elsewhere in our Pulumi program
	ctx.Export("topicName", topic.Name)
	return topic, nil
}
