# Simple GO REST API with Gin and Kafka events

This is an API built with the Go programming language, the Gin web framework and Kafka to stream events. The purpose of this project is to learn the basic concepts of building a RESTful API using Go, Gin and Kafka Events in SAGA pattern.

## Table of Contents

- [Installation](#installation)
- [SAGA Pattern](#saga-pattern)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Testing the Application](#testing-the-application)
- [Contributing](#contributing)
- [License](#license)

## Installation

To run this project, you need to have GO installed on your machine. You can download it from the [official Go website](https://golang.org/dl/).
Also, you need to have Docker installed to run Kafka, Zookeeper and the services that are part of the SAGA pattern implementation.

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/MrRique15/go_rest_api.git
   ```
2. Change into the project directory:
   ```bash
   cd go_rest_api
   ```
3. Change MongoDB URI and MongoDB Database env in docker-compose.yml

4. build and Start Docker containers:
   ```bash
   docker-compose up --build
   ```

## SAGA Pattern

The SAGA pattern is a way to manage distributed transactions. It is a sequence of local transactions where each transaction updates the database and publishes a message or event to trigger the next transaction in the sequence. If one of the transactions fails, the SAGA pattern uses compensating transactions to undo the changes made by the preceding transactions.

You can check the SAGA pattern implementation in this project by checking the [inventory_service](inventory_service) and [main_api](main_api) folders.
Also, you can check the Diagram of the SAGA pattern implementation in the image below: 
![SAGA Pattern](saga_pattern_diagram.png)

## Usage

After running Docker containers, you can access the API in the following URL: `http://localhost:8080`.
All the requests made to this URL will be handled by the API Gateway, which will publish events to Kafka topics when needed, to trigger the services that are part of the SAGA pattern implementation.

## API Endpoints

To check the API endpoints, you can access the [routes documentation](routes_doc/README.md).

## Testing the Application

To test the API Gateway application, execute the following command:

```bash
cd main_api
go test -v
```

It will run all tests that has been set in files like *_test.go

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you have any suggestions or improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
