# Go Monitoring Agent

## Purpose
DevOps training repository focused on **Go, JSON, YAML**.

## Application
Tiny Go HTTP health service. Build with `go build ./...`; run `go run .`; check `/health`. Dockerfile shows a static binary in a `scratch` runtime.

## Suggested student exercises
1. Inspect the source and identify each language/configuration type.
2. Build and run the app locally.
3. Add or improve automated tests and CI validation.
4. Build a container image and tag it with the Git commit SHA.
5. Add dependency/image scanning to the pipeline.
6. Deploy to a Kubernetes development namespace where appropriate.
7. Record build metadata and define a rollback procedure.

## Repository note
This is a teaching starter project. Review the workflow and build configuration for your environment before using it in production.
