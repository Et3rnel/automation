package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// Config holds all your Pulumi configuration values.
type Config struct {
	GCPProjectId string
}

// NewConfig loads the Pulumi configuration from the current stack
// and returns a pointer to a Config structure.
func NewConfig(ctx *pulumi.Context) (*Config, error) {
	gcpCfg := config.New(ctx, "gcp")

	// Retrieve the required GCP project ID from the configuration.
	projectId := gcpCfg.Require("project")
	return &Config{
		GCPProjectId: projectId,
	}, nil
}
