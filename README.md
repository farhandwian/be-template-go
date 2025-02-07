# CC BE

## Setup 
Create file `.env` based on `.env.example`



## How to run it using docker
```
$ docker network create my-network
$ docker run --name mariadb --network my-network -e MYSQL_ROOT_PASSWORD=12345 -e MYSQL_DATABASE=warehouse_db -p 3306:3306 -d mariadb:latest
$ docker build -t my-app .
$ docker run --name my-app --network my-network -p 8081:8081 my-app
```

## Planing
 - [x] Connect to DB via gorm
 - [x] Transaction
 - [x] Open API generator
 - [x] Sample output in open api
 - [x] Sample application template
 - [x] Separate StatusInternalServerError dan StatusBadRequest
 - [x] Gracefully shutdown
 - [ ] Send Email
 - [ ] Proper logging library
 - [ ] Whatsapp OTP
