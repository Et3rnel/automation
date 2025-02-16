package main

import (
	"log"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Load configuration into our custom Config structure
		config, err := NewConfig(ctx)
		if err != nil {
			return err
		}

		topic, err := provisionTopic(ctx)
		if err != nil {
			return err
		}
		// TODO: dummy endpoint to change
		_, err = configureSubscription(ctx, topic, "https://localhost/push-endpoint")
		if err != nil {
			return err
		}

		// Call Gmail Watch API with the configuration
		if err := callGmailWatch(config); err != nil {
			log.Fatalf("Error calling Watch API: %v", err)
		}

		return nil
	})
}
