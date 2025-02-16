package main

import (
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
