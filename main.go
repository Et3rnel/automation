package main

import auth "automation/internal"

func main() {
	auth.Run()

	// pulumi.Run(func(ctx *pulumi.Context) error {
	// 	/*// Load configuration into our custom Config structure
	// 	config, err := NewConfig(ctx)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	topic, err := provisionTopic(ctx)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// TODO: dummy endpoint to change
	// 	_, err = configureSubscription(ctx, topic, "https://localhost/push-endpoint")
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// Call Gmail Watch API with the configuration
	// 	if err := callGmailWatch(config); err != nil {
	// 		log.Fatalf("Error calling Watch API: %v", err)
	// 	}*/
	// 	gcpProvider, err := gcp.NewProvider(ctx, "gcp", &gcp.ProviderArgs{})
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// Create a service account
	// 	serviceAccount, err := serviceaccount.NewAccount(ctx, "serviceAccount", &serviceaccount.AccountArgs{
	// 		AccountId:   pulumi.String("my-service-account"),
	// 		DisplayName: pulumi.String("My Service Account"),
	// 	}, pulumi.Provider(gcpProvider))
	// 	if err != nil {
	// 		return err
	// 	}

	// 	fmt.Println(serviceAccount.Name)

	// 	return nil
	// })
}
