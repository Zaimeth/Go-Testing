# OMDB API

## Tech-Stack
* Go, with modules
    * gin
    * gorm
    * godotenv
    * mapstructure
    * logrus
* PostgreSQL


## Prerequisites
* Go v1.14
* PostgreSQL

## Get Started
1. Setup your database
    * You can copy and run SQL script from ```app/migration/X_X_X.sql``` for create tables.
2. You must create ```.env``` file, you can copy from ```.env.example``` and then customize the value.

## Rest Quick Start
This project using Makefile for simplify your commands.

1. Run project
    ```
    $ make run-rest
    ```

2. Build project
    ```
    $ make build-rest
    ```

3. Run binary file
    ```
    $ make start-rest
    ```
   
## GRPC Quick Start
sorry for grpc service i can't finish it yet


