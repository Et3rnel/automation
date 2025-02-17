package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/pubsub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func provisionTopic(ctx *pulumi.Context) (*pubsub.Topic, error) {
	// Create a new Pub/Sub topic with the Pulumi resource name "gmailTopic" used internally to identify the resource
	topic, err := pubsub.NewTopic(ctx, "gmailTopic", &pubsub.TopicArgs{
		// Set the actual name of the topic using the constant
		Name: pulumi.String(TopicName),
	})

	if err != nil {
		return nil, err
	}

	// Export the topic name so that it can be referenced elsewhere in our Pulumi program
	ctx.Export("topicName", topic.Name)
	return topic, nil
}
