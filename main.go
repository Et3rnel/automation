package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/pubsub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		topic, err := provisionTopic(ctx)
		if err != nil {
			return err
		}
		// TODO: dummy endpoint to change
		_, err = configureSubscription(ctx, topic, "https://localhost/push-endpoint")
		if err != nil {
			return err
		}
		return nil
	})
}

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

// configureSubscription creates a push subscription for the given topic.
// It configures the subscription to push messages to the specified endpoint.
//
// Parameters:
// - ctx: the Pulumi execution context.
// - topic: the Pub/Sub topic resource to associate the subscription with.
// - pushEndpoint: the URL endpoint to which the push notifications will be delivered.
func configureSubscription(ctx *pulumi.Context, topic *pubsub.Topic, pushEndpoint string) (*pubsub.Subscription, error) {
	// Create a new Pub/Sub subscription with push configuration that listens to the topic and pushes messages to the provided endpoint.
	subscription, err := pubsub.NewSubscription(ctx, "gmailSubscription", &pubsub.SubscriptionArgs{
		Topic: topic.Name,
		PushConfig: &pubsub.SubscriptionPushConfigArgs{
			PushEndpoint: pulumi.String(pushEndpoint),
		},
	})
	if err != nil {
		return nil, err
	}

	// Export the subscription name so that it can be easily referenced after deployment.
	ctx.Export("subscriptionName", subscription.Name)
	return subscription, nil
}
