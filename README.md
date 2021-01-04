## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites
1. Golang       => Backend
2. Docker(optional)
3. Make

## Development Stack
1. Golang
2. ReactJS
3. Mariadb(Bitnami galera cluster)

### Running & Initializing database
`make prepare`

## Running the tests

`make test`
`make test-coverage`

## Start/Stop Backend Application

`make run`
`make stop`

## Reset Backend Application

`make reset`