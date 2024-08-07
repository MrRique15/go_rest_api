# Simple GO REST API with Gin

This is a simple API built with the Go programming language and the Gin web framework. The purpose of this project is to learn the basic concepts of building a RESTful API using Go and Gin.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Running the Application](#running-the-application)
- [Building the Application](#building-the-application)
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

## Usage

To start the server, run the following command:

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

## API Endpoints

### POST /users/register

- **Description**: Register a new user in the system.
- **Request Body**: 
  - Example:
    ```json
    {
	    "name": "string",
	    "email": "string",
	    "password": "string",
	    "confirm_password": "string"
    }
    ```
- **Response**: 
  - Status: `201 Created`
  - Body:
    ```json
    {
	    "status": 201,
	    "message": "success",
	    "data": {
	    	"data": {
	    		"InsertedID": "ObjectID"
	    	}
	    }
    }
    ```

### POST /users/login

- **Description**: Login a user in the system.
- **Request Body**: 
  - Example:
    ```json
    {
	    "email": "string",
	    "password": "string"
    }
    ```
- **Response**: 
  - Status: `200 OK`
  - Body:
    ```json
    {
	    "status": 200,
	    "message": "sucess",
	    "data": {
	    	"data": {
	    		"id": "ObjectID",
	    		"name": "string",
	    		"email": "string"
	    	}
	    }
    }
    ```

### PUT /users/edit/:id

- **Description**: Edit an existing user.
- **Request Body**: 
  - Example:
    ```json
    {
	    "id": "ObjectID",
	    "name": "string",
	    "email": "string",
	    "old_password": "string",
	    "new_password": "string"
    }
    ```
- **Response**: 
  - Status: `200 OK`
  - Body:
    ```json
    {
	    "status": 200,
	    "message": "success",
	    "data": {
	    	"data": {
	    		"id": "ObjectID",
	    		"name": "string",
	    		"email": "string"
	    	}
	    }
    }
    ```

### GET /users/:id

- **Description**: Get an user by ID.
- **Response**: 
  - Status: `200 OK`
  - Body:
    ```json
    {
	    "status": 200,
	    "message": "success",
	    "data": {
	    	"data": {
	    		"id": "ObjectID",
	    		"name": "string",
	    		"email": "string"
	    	}
	    }
    }
    ```

## Running the Application

To run the application, execute the following command:

```bash
go run main.go
```

## Testing the Application

To test the application, execute the following command:

```bash
go test -v
```

It will run all tests that has been set in files like *_test.go

## Building the Application

To run build application, execute the following command:

```bash
go build
```
And then run it by executing the right command for your OS:

### Windows
```bash
.\go_rest_api.exe
```

### Linux
```bash
./go_rest_api
```

This will start the API server on port 8080. You can use tools like [Postman](https://www.postman.com/) or `curl` to interact with the API.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you have any suggestions or improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
