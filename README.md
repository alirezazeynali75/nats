# NATS Project

This project demonstrates a simple implementation of a NATS-based messaging system with JetStream support. It includes a publisher and a subscriber service, both of which interact with a NATS server.



## Features

- **NATS Server** with JetStream enabled.
- **Publisher** service to send messages to a specific subject.
- **Subscriber** service to listen for messages on a specific subject.
- Dockerized setup for easy deployment.

## Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/) (if running locally)

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/alirezazeynali75/nats.git
cd nats
```

### 2. Build and Run with Docker Compose

```bash
docker-compose up --build
```

This will:

- Start the NATS server with JetStream enabled.
- Build and run the publisher and subscriber services.

### 3.  Publisher and Subscriber

- The Publisher sends messages to a subject specified via the --subject flag.
- The Subscriber listens for messages on the same subject.

#### Example Commands

- Publish a message:

```sh
docker-compose exec publisher ./publisher --subject "example.subject" --data "Hello, NATS!"
```

- Subscriber logs the received message:

```sh
docker-compose logs subscriber
```
