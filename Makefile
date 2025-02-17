.PHONY: preview

# Show a preview of updates to a stack’s resources
preview:
	pulumi preview

# Ensures hthe go.mod and go.sum files accurately reflect the project's dependencies by adding any missing modules and removing those no longer needed
tidy:
	go mod tidy

run:
	go run main.go