package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/pubsub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
