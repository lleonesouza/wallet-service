## Get Started 
 
1. Put environments on ```.env``` using the ```.env.example``` as guide
2. Generate ORM needed files ```go run github.com/prisma/prisma-client-go generate```
3. Run the containers ```docker-compose up```
4. Add ```127.0.0.1 db```  to ```/etc/hosts``` file (step required to up migrations on ṕostgres container using container name)
5. Up Migrations ```go run github.com/prisma/prisma-client-go migrate dev```


### Business Logic

- Every user has ONE Wallet
- Every Wallet start with *WALLET_BALANCE_INIT* (set environment) coins
- User can send/receive coins
- Shopkeeper can only receive coins

### Auth

Users can be created calling a POST request to ```/user``` with a json containing *email*, *password*, *cpf*, *name* and *lastname*. 

Example:

```json
{
  "email": "john@doe.com",
  "password": "12345678",
  "name": "john",
  "lastname": "doe",
  "cpf": "123456789"
}
```

After created, login can be done by calling a POST request to ```/user/login``` with a json containing the *email* and *password*

Example:

```json
{
  "email": "john@doe.com",
    "password": "12345678"
}
```

Response:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImxsZW9uZXNvdXphNDMxMkBsaXZlLmNvbSIsImlkIjoiY2Q2NTU2YjktZjk1ZC00MmM0LWJhYzgtOTQwNjk5NDc0MTY0IiwiZXhwIjoxNjc1Mzg4ODIxfQ.z29cwjAuaHlE8ee1mVYf5lyO6owmPYTiouqxojs6cF4"
}
```

Now the *token* can be used as *Bearer Token* to make request to protected routes, the same logic works for Shopkeeper.

# Swagger

To generate swagger docs files, use

```
go run github.com/swaggo/swag/cmd/swag init -g main.go --output docs
```

then use /swagger/ to access

# API Routes

## User:
| TYPE  | URL | DESCRIPTION | 
| - | - | - |
| POST | /user  | Creates User | 
| POST | /user/login | Authenticate User with JWT | 
| GET | /user | Gets user info with balance | 
| PUT | /user | Updates 'name' and 'lastname' User field | 


## Shopkeeper:
| TYPE  | URL | DESCRIPTION | 
| - | - | - |
| POST | /shopkeeper  | Creates Shopkeeper | 
| POST | /shopkeeper/login | Authenticate Shopkeeper with JWT | 
| GET | /shopkeeper | Get Shopkeeper info with balance | 
| PUT | /shopkeeper | Update 'name' and 'lastname' Shopkeeper field |

## Transaction:
| TYPE  | URL | DESCRIPTION | 
| - | - | - |
| POST | /transaction  | Creates a transaction | 
| GET | /transaction | List trasactions from Wallet | 

## Swagger:
| TYPE  | URL | DESCRIPTION | 
| - | - | - |
| POST | /swagger/  | Render swagger API docs | 


## Folders Structure
### Prisma
Contains *schema.prisma* where is used to generate and migrate the table on the database, also contains the migration folder that has the migrations history

### Handlers
Contains Echo HTTP handlers functions and DTOs (Data Transfer Objects), responsible for the communication layer.

### Services
Contains the bussiness logic functions that are injected in Handlers

### Errors
Contains errors struct

### Config
Contains env and configuration structs 

# Contribute

Feel free to create a Pull Request, it will be welcome
