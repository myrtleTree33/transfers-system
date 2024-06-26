# Transfers System

Transfers System is a template for a web application that allows users to transfer money between accounts. This template is built using the TransfersSystem API and is designed to be a starting point for developers to build their own applications.

## Features

The TransfersSystem Template Starter includes the following features:

- Nix for virtualised local development
- Docker for containerisation
- Gorm for database management
- Mockery for mocking
- GORM for database ORM
- Goose for database migrations
- Air for live reloading
- Postman for API testing
- Swagger for API documentation
- CORS for cross-origin resource sharing

## Installation

To install the TransfersSystem Template Starter, you will need to have a working installation of Docker, Git and Devbox (Virtual environment based off Nix-Shell).  

Clone the repository
```bash
git clone <https://github.com/myrtleTree33/transfers-system>
```

Install Git
```bash
sudo apt-get install git
```

Install Docker
```bash
sudo apt-get install docker
```

Install Devbox: See https://www.jetify.com/devbox/docs/installing_devbox/

## Entering the dev environment

To enter the dev environment, run the following command in the root directory of the project:
```bash
make
make launch
```

## Running the application

To run the application, run the following command in the root directory of the project:

```bash
cd ./backend
air
```

## Database migrations

To create a new migration, run the following command in the root directory of the project:

```bash
cd backend/
go run ./cmd/goosecli create migrations/add_idempotencies_table sql
```

To run the migrations, run the following command in the root directory of the project:

```bash
cd backend/
go run ./cmd/goosecli up
```

## Edge cases handled

The following edge cases are handled in the application:

### Duplicate transfers / idempotency

Create and update operations are idempotent. If a transfer with the same ID is created or updated multiple times, the operation will only be performed once.

This is enforced by the middleware that checks for the existence of a transfer with the same ID in the database before creating or updating the transfer.  A HTTP header "idempotency-key" is used to identify the transfer.

The process works as follows:

1. Store the request in an idempotency table with the idempotency key + hash of the request data as the primary key.
2. Check if the request has already been processed by looking up the idempotency key in the idempotency table.
3. If the request has already been processed, return the response from the previous request.
4. If the request has not been processed, process the request and store the response in the idempotency table.
5. Return the response to the client.

#### Example of enforcing idempotency

For future requests, include the middleware in the request to check for idempotency: https://github.com/myrtleTree33/transfers-system/blob/747f692d3eb0ea77476f0a475eceedae1736697f/backend/main.go#L112-L115


Next, reply with the wrapped function which stores the response in the database: https://github.com/myrtleTree33/transfers-system/blob/747f692d3eb0ea77476f0a475eceedae1736697f/backend/internal/controllers/accounts_routes.go#L54

### Insufficient funds

If a transfer is created with an amount that exceeds the balance of the source account, the transfer will not be created and an error message will be returned.

### Negative transfer amounts

If a transfer is created with a negative amount, the transfer will not be created and an error message will be returned.

### Invalid account IDs

If a transfer is created with an invalid source or destination account ID, the transfer will not be created and an error message will be returned.

## Postman collection

To test the API, you can use the Postman collection provided in the `postman` directory. This collection contains a set of requests that can be used to interact with the API.

The collection can be imported into Postman by clicking on the `Import` button in the top left corner of the Postman window and selecting the `transfers-system.postman_collection.json` file.

https://api.postman.com/collections/772335-cbcd4ab3-71d9-4cd5-9fff-f8a89d48d907?access_key=PMAT-01HYREPC21C8QMY0QCDGED8MSY

## Authors

* **Joel Tong** - *Initial work* - [Joel Tong](https://github.com/myrtletree33)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

