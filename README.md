# Money Transfers API

A simple API that simulates money transfers 
between different accounts. The project consists of two main models. The
first, is for managing accounts (CRUD) while the other is for
managing transactions between accounts. 
Next, I will explain a detailed description on folder structure and how
to run and test this project. The following diagram depicts the relationship
between account and transaction.

## Notes: the application is tested only on Linux Ubuntu.

![diagram](diagram.svg "Diagram")

## Table of Content
1. [Money Transfers API](#money-transfers-api)
1. [Folder Structure](#folder-structure)
1. [Helper Commands](#helper-commands)


## Folder Structure:

* `api/` Contains API methods separated by domain model.
* `cmd/` Here is the main package to start.
* `internal/`
  * `models/` struct for domain models
  * `multiplexer/` instead of importing a third party library I stuck to 
    only use standard libraries and create my own mux.
  * `services/` domain model services
* `Makefile` commands shortcuts
* `Transfer Money.postman_collection.json` exported Postman collection to
easing API test.


## Helper Commands:

* `make run` to run the API without the need to create an executable file.
  
* `make build` compiles the code and generates an executable file that will
be located in bin directory.
  
* `make test` runs all tests.

* `make curl-add-account` an example of using curl command to add a sample
new account.

