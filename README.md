# Simple GO REST API with Gin and Kafka events

This is an API built with the Go programming language, the Gin web framework and Kafka to stream events. The purpose of this project is to learn the basic concepts of building a RESTful API using Go, Gin and Kafka Events in SAGA pattern.

## Table of Contents

- [Installation](#installation)
- [SAGA Pattern](#saga-pattern)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Testing the Application](#testing-the-application)
- [Running Kafka and Zookeeper with Docker](#running-kafka-and-zookeeper-with-docker)
- [Running Services](#running-services)
   - [Running Saga Execution Controller](#running-saga-execution-controller)
   - [Running Inventory Service](#running-inventory-service)
   - [Running Payment Service](#running-payment-service)
   - [Running Shipping Service](#running-shipping-service)
- [Building the API Gateway](#building-the-api-gateway)
- [Contributing](#contributing)
- [License](#license)

## Installation

To run this project, you need to have Go installed on your machine. You can download it from the [official Go website](https://golang.org/dl/).

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/MrRique15/go_rest_api.git
   ```
2. Change into the project directory:
   ```bash
   cd go_rest_api
   ```
3. Install the dependencies:
   ```bash
   go mod tidy
   ```

## SAGA Pattern

The SAGA pattern is a way to manage distributed transactions. It is a sequence of local transactions where each transaction updates the database and publishes a message or event to trigger the next transaction in the sequence. If one of the transactions fails, the SAGA pattern uses compensating transactions to undo the changes made by the preceding transactions.

You can check the SAGA pattern implementation in this project by checking the [inventory_service](inventory_service) and [main_api](main_api) folders.
Also, you can check the Diagram of the SAGA pattern implementation in the image below: 
![SAGA Pattern](saga_pattern_diagram.png)

## Usage

To start the server, run the following command:

```bash
cd main_api
go run main.go
```

The server will start on `http://localhost:8080`.

## API Endpoints

To check the API endpoints, you can access the [routes documentation](routes_doc/README.md).

## Testing the Application

To test the application, execute the following command:

```bash
cd main_api
go test -v
```

It will run all tests that has been set in files like *_test.go

## Running Kafka and Zookeeper with Docker

To run Kafka and Zookeeper with Docker, execute the following command:

```bash
cd kafka
docker-compose up
```
## Running Services
This section will show how to run the services that are part of the SAGA pattern implementation. All services need to be running to test the SAGA pattern correctly.
Obs*: In further versions of this project, it will be implemented a way to run all services with a single command using Docker Compose.

### Running Saga Execution Controller
The Saga Execution Controller will watch for new events by consuming the `order_sec` topic from kafka, triggered by the API Gateway and servicer. It will control the sequence of events to complete order processing or cancelation.

```bash
cd saga_execution_controller
go run main.go
```

### Running Inventory Service

The inventory service will watch for new events by consuming the `inventory` topic from kafka, triggered by Saga Execution Controller and then it will update items inventory and order status.
If the item is out of stock, it will send a message to the Saga Execution Controller informing the order cancelation, using the `order_sec` topic.
On the other hand, if the item is in stock, it will send a message to the Saga Execution Controller to proceed with the order, using the `order_sec` topic.

```bash
cd inventory_service
go run main.go
```

### Running Payment Service

The payment service will watch for new events by consuming the `payment` topic from kafka, triggered by Saga Execution Controller and then it will update order status when payment is completed.
If the order isn't reserved in inventory or the payment is not completed, it will send a message to the Saga Execution Controller informing the problem, using the `order_sec` topic.
On the other hand, if it's all checked, it will send a message to the Saga Execution Controller to proceed with the order, using the `order_sec` topic.

```bash
cd payment_service
go run main.go
```

### Running Shipping Service

The shipping service will watch for new events by consuming the `shipping` topic from kafka, triggered by Saga Execution Controller and then it will update order status when shipping is started.
If the order isn't paid, it will send a message to the Saga Execution Controller informing the problem, using the `order_sec` topic.
On the other hand, if it's all checked, it will send a message to the Saga Execution Controller to proceed with the order, using the `order_sec` topic.

```bash
cd shipping_service
go run main.go
```

## Building the API Gateway

To run build the API Gateway, execute the following command:

```bash
cd main_api
go build
```
And then run it by executing the right command for your OS:

### Windows
```bash
cd main_api
.\main_api.exe
```

### Linux
```bash
cd main_api
./main_api
```

This will start the API server on port 8080. You can use tools like [Postman](https://www.postman.com/) or `curl` to interact with the API.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you have any suggestions or improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
