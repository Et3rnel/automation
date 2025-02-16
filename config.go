package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Config holds all your Pulumi configuration values.
type Config struct {
	GCPProjectId string
}

// NewConfig loads the Pulumi configuration from the current stack
// and returns a pointer to a Config structure.
func NewConfig(ctx *pulumi.Context) (*Config, error) {
	cfg := pulumi.NewConfig()
	// Retrieve the required GCP project ID from the configuration.
	projectId := cfg.Require("gcp:project")
	return &Config{
		GCPProjectId: projectId,
	}, nil
}
