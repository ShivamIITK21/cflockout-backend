## CFLockout Backend
 
This is the backend for CFLockout

# How to Run?

1). Create an .env file with the following

```
HOST=localhost
USER=<postgres username>
PASSWORD=<postgres pass>
DB_NAME=<Name of DB>
DB_PORT=<Port of DB>
SSL_MODE=disable
SERVER_PORT=<port of server>
SECRET_KEY=<key for jwt auth>
```

2). Spin up the server

```
go run main.go
```


3). The following routes have been implemented

```
POST /auth/login
POST /auth/signup
POST /lockout/create
GET /lockout?session_id=
GET /problems?user=
GET /problems/refresh
```