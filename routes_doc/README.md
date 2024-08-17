## CUSTOMER Endpoints

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
	    		"_id": "ObjectID",
          "name": "string",
          "email": "string",
          "password": "string"
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
	    		"_id": "ObjectID",
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
	    "_id": "ObjectID",
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
	    		"_id": "ObjectID",
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
	    		"_id": "ObjectID",
	    		"name": "string",
	    		"email": "string"
	    	}
	    }
    }
    ```

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

## ORDER Endpoints

### POST /orders/new

- **Description**: Register a new order in the system and stream an event to Kafka.
- **Request Body**: 
  - Example:
    ```json
    {
      "customer_id": "ObjectID",
      "price": "float",
      "items": [
        {
          "product_id": "ObjectID",
          "quantity": "int"
        }
      ],
      "status": "string"
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
          "_id": "ObjectID",
          "customer_id": "ObjectID",
          "price": "float",
          "items": [
            {
              "product_id": "ObjectID",
              "quantity": "int"
            }
          ],
          "status": "string"
	    	}
	    }
    }
    ```

### PUT /orders/update

- **Description**: Update an existing order in the system and stream an event to Kafka.
- **Request Body**: 
  - Example:
    ```json
    {
      "_id": "ObjectID",
      "customer_id": "ObjectID",
      "price": "float",
      "items": [
        {
          "product_id": "ObjectID",
          "quantity": "int"
        }
      ],
      "status": "string"
    }
    ```
- **Response**: 
  - Status: `200 OK`
  - Body:
    ```json
    {
	    "status": 201,
	    "message": "success",
	    "data": {
	    	"data": {
          "_id": "ObjectID",
          "customer_id": "ObjectID",
          "price": "float",
          "items": [
            {
              "product_id": "ObjectID",
              "quantity": "int"
            }
          ],
          "status": "string"
	    	}
	    }
    }
    ```

## PRODUCTS Endpoints

### POST /products/new

- **Description**: Register a new product in the system.
- **Request Body**: 
  - Example:
    ```json
    {
      "name": "string",
      "stock": "int"
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
          "_id": "ObjectID",
          "name": "string",
          "stock": "int"
      	}
      }
    }
    ```

### PUT /products/update

- **Description**: Update an existing product in the system.
- **Request Body**: 
  - Example:
    ```json
    {
      "_id": "ObjectID",
      "name": "string",
      "stock": "int"
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
          "_id": "ObjectID",
          "name": "string",
          "stock": "int"
      	}
      }
    }
    ```
