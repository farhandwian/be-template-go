# IAM

## Use Case
tempat bisnis proses utama yang mengorkestrasi model dan gateway

## Gateway
kode infrastruktur seperti pengaksesan database, kirim email, kirim whatsapp dan fungsi lainnya. Gateway harus mereturn AppError sebagai error

## Controller 
tempat mempublish use case menjadi API

## Model
tempatnya Entity dan Value object yang akan dipakai oleha use case

## Wiring
tempat menggabungkan usecase, gateway, middlware, dan controller

## Core
Framework kontrak utama yang dipakai oleh use case dan gateway

### API
```
Tag                          Access          Summary                             Method   URL
----------------------------------------------------------------------------------------------------------------------------------
IAM - User Management        ADMIN           Get all users                       GET      /users                                  
IAM - User Management        USER            Get current user detail             GET      /users/me                               
IAM - User Management        ADMIN           Get user detail by id               GET      /users/{id}                             
IAM - User Management        ADMIN           Update to user detail               PUT      /users/{id}                             
IAM - User Access            ADMIN           Get user access by id               GET      /users/{id}/access                      
IAM - User Access            ADMIN           Set access to user                  POST     /users/{id}/access                      
IAM - Account Management     ADMIN           Register user                       POST     /account/register                       
IAM - Account Management     ADMIN           Send email activation request       POST     /account/activate/initiate              
IAM - Account Management     ANONYMOUS       Submit email activation             POST     /account/activate/verify                
IAM - Authentication         ANONYMOUS       Initiate user login                 POST     /auth/login                             
IAM - Authentication         ANONYMOUS       Submit OTP for login                POST     /auth/login/otp                         
IAM - Authentication         ANONYMOUS       Get the new access token            POST     /auth/refresh-token                     
IAM - Authentication         USER            Logout session                      POST     /auth/logout                            
IAM - Pin Management         USER            Initiate change pin                 POST     /pin/change/initiate                    
IAM - Pin Management         USER            Submit pin changes                  POST     /pin/change/verify                      
IAM - Password Management    USER            Initiate change password            POST     /password/change/initiate               
IAM - Password Management    USER            Submit OTP for password changes     POST     /password/change/verify                 
IAM - Password Management    ADMIN           Initiate password reset             POST     /password/reset/initiate                
IAM - Password Management    ANONYMOUS       Submit password reset               POST     /password/reset/verify                  
CC                           OTHERS          Function one                        POST     /function/one                           
CC                           OTHERS          Function two                        GET      /function/two                           
CC                           OTHERS          Function three                      PUT      /function/three                         
CC                           OTHERS          Function four                       DELETE   /function/four                          

SWAGGER https://editor.swagger.io/?url=http://localhost:8081/openapi 
```

### Test
```
go test ./model -v
go test ./controller_test -v
```

### Test Individual
```
go test ./controller_test/shared.go ./controller_test/refresh_token_test.go -v
```

### Flow Test
```
go test ./flow_test/shared.go ./flow_test/01_test.go -v
go test ./flow_test/shared.go ./flow_test/02_test.go -v
```


### Docker
```
docker run --name mariadb-container -e MYSQL_ROOT_PASSWORD=12345 -e MYSQL_DATABASE=warehouse_db -p 3306:3306 -d mariadb:latest
```

