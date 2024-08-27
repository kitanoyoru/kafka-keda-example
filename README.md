do not use, from gpt btw :)

# Kafka KEDA Example

## Overview

**kafka-keda-example** is a sample project that demonstrates how to use [Kafka](https://kafka.apache.org/) as a messaging platform with [KEDA](https://keda.sh/) (Kubernetes Event-Driven Autoscaling) for scaling workloads based on external event sources. This project serves as a practical guide for developers looking to implement Kafka with KEDA in their Kubernetes environments.

## Features

- Integration with Kafka for message streaming.
- KEDA for automated scaling of consumers based on message queue length.
- Simple example demonstrating message production and consumption.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Docker](https://www.docker.com/) (for containerization)
- [Docker Compose](https://docs.docker.com/compose/) (for managing multi-container applications)
- [Kubernetes](https://kubernetes.io/) (for running the application in a cluster)
- [kubectl](https://kubernetes.io/docs/tasks/tools/) (Kubernetes command-line tool)
- [KEDA](https://keda.sh/docs/2.0/) installed in your Kubernetes cluster

## Project Structure

```
kafka-keda-example/
├── deployments/
│   ├── docker-compose.yaml
│   └── keda-deployment.yaml
├── schema/
│   ├── avro/
│   │   └── user.avsc
├── cmd/
│   └── main.go
├── Dockerfile
├── go.mod
└── README.md
```

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/kafka-keda-example.git
   cd kafka-keda-example
   ```

2. **Build the Docker image:**

   ```bash
   docker build -t kitanoyoru/kafka-keda-example:latest .
   ```

3. **Start Kafka and KEDA with Docker Compose:**

   ```bash
   docker compose -f ./deployments/docker-compose.yaml up -d
   ```

4. **Deploy the application to your Kubernetes cluster:**

   Apply the KEDA deployment configuration:

   ```bash
   kubectl apply -f ./deployments/keda-deployment.yaml
   ```

## Usage

1. **Produce messages to Kafka:**

   You can produce messages to the specified topic in Kafka. Adjust the producer settings (if necessary) in the code.

2. **Consume messages:**

   The consumer will automatically scale based on the message queue length, thanks to KEDA. Monitor the scaling behavior through Kubernetes commands:

   ```bash
   kubectl get deployments
   kubectl get hpa
   ```

## Configuration

You can customize the configuration by modifying the environment variables in the `docker-compose.yaml` file and `keda-deployment.yaml` file as necessary. Key configuration options include:

- `BOOTSTRAP_SERVERS`: Kafka bootstrap server address.
- `TOPIC`: Kafka topic to consume from.
- `CONSUMER_GROUP_ID`: Consumer group identifier.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you wish to contribute to this project.

## Author

[Your Name](https://yourwebsite.com) - [Your Email](mailto:youremail@domain.com)


