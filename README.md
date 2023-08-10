# Foundation

[![Go Build](https://github.com/ri-nat/foundation/actions/workflows/go.yml/badge.svg)](https://github.com/ri-nat/foundation/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ri-nat/foundation)](https://goreportcard.com/report/github.com/ri-nat/foundation)
[![License](https://img.shields.io/github/license/ri-nat/foundation)](https://opensource.org/licenses/MIT)

> **Early Development Notice:** Foundation is currently in an early development stage. While you're welcome to explore and experiment, it's not yet ready for production use.

## Overview

Foundation is a Go-based microservices framework aimed to help developers write scalable, resilient, and highly available applications with ease. By providing a cohesive set of well-chosen tools and features, Foundation aims to minimize boilerplate and allow developers to focus on writing business logic.

Foundation is built upon several proven technologies including:

- **gRPC**: A high-performance, open-source RPC framework.
- **gRPC Gateway**: A gRPC to JSON reverse proxy.
- **Protocol Buffers**: A language-neutral extensible mechanism for serializing structured data, used for gRPC and Kafka message serialization.
- **Kafka**: A powerful distributed streaming platform.
- **OAuth 2.0**: An industry-standard authorization framework.
- **PostgreSQL**: A robust open-source relational database system.
- **WebSockets**: Enabling real-time, bi-directional, and full-duplex communication channels over TCP connections.

## Key Features

- [x] **Running Modes**: Choose the mode that best fits your needs: `gateway`, `grpc`, `http`, `worker`, `events_worker`, or `job`.
- [x] **Transactional Outbox**: Implement the transactional outbox pattern for transactional message publishing to Kafka.
- [x] **Unified Logging**: Conveniently log with colors during development and structured logging in production using `logrus`.
- [ ] **Tracing**: Trace and log your requests in a structured format with OpenTracing.
- [x] **Metrics**: Collect and expose service metrics to Prometheus.
- [x] **Health Check**: Provide Kubernetes with health status of your service.
- [x] **(m)TLS**: TLS authentication for Kafka and mTLS for gRPC.
- [x] **Graceful Shutdown**: Ensure clean shutdown on `SIGTERM` signal reception.
- [x] **Helpers**: A variety of helpers for common tasks.
- [ ] **CLI**: A CLI tool to help you get started and manage your project.

## Integrations

Foundation comes with built-in support for:

- [x] **PostgreSQL**: Easily connect to a PostgreSQL database.
- [x] **Dotenv**: Load environment variables from .env files.
- [x] **ORY Hydra**: Authenticate users on a gateway with ORY Hydra.
- [x] **gRPC Gateway**: Expose gRPC services as JSON endpoints.
- [x] **Kafka**: Produce and consume messages with Kafka (via `kafka-go`).
- [x] **Sentry**: Report errors to Sentry.

## Getting Started

Currently, the best way to get started is by exploring the codebase. We're working on providing sample implementations and thorough documentation to make the onboarding process even smoother. Stay tuned!

## CLI Tool

To install the CLI tool, run:

```bash
go install github.com/ri-nat/foundation/cmd/foundation@main
```

There are several commands available:

```bash
foundation completion # Generate shell completion scripts (prints to stdout)
foundation db:migrate # Run database migrations
foundation db:rollback # Rollback database migrations
foundation start # Start the service (you will be prompted to choose a service to start)
foundation test # Run tests
```

You can also run `foundation` without any arguments to see a list of available commands, or run `foundation <command> --help` to see the available options for a specific command.

## Contributing

We're always looking for contributions from the community! If you've found a bug, have a suggestion, or want to add a new feature, feel free to open an issue or submit a pull request.

## License

Foundation is released under the [MIT License](./LICENSE).
