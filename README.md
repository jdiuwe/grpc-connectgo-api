# grpc-connectgo-api

basic gRPC API server which supports user registration, login and getting user profile.

### Components:
- [connect-go](https://connectrpc.com/) - implementation of Protobuf RPC to build gRPC-compatible HTTP API.
- [protobuf](https://protobuf.dev/getting-started/gotutorial/) - protocol buffer schema to define rpc methods and payload
- [buf](https://buf.build/) - tools to simplify dealing with protocol buffers
- [postgresql](https://www.postgresql.org/) - relational database to store users
- [sqlc](https://github.com/sqlc-dev/sqlc) - type-safe SQL compiler
- [migrate](https://github.com/golang-migrate/migrate) - db migrations library
- [JSON Web Tokens (JWT)](https://jwt.io/) - user claims
- [docker-compose](https://docs.docker.com/compose/) - tool to simplify running local Docker containers 


## RPC methods

### RegisterUser

request
```json
{
    "first_name": "john",
    "last_name": "smith",
    "email": "jsmith@pm.me",
    "password": "changeme"
}
```

response:
```json
{
    "message": "user account created",
    "status": "success",
    "uuid": "9f11f2c8-05ef-444e-ac69-7bcf1b792160"
}
```

LoginUser

request:
```json
{
    "email": "jsmith@pm.me",
    "password": "changeme"
}
```

response:
```json
{
    "message": "user logged in",
    "status": "success",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTY5NTA5MTEsInVzZXJfdXVpZCI6IjlmMTFmMmM4LTA1ZWYtNDQ0ZS1hYzY5LTdiY2YxYjc5MjE2MCJ9.I94GOOuHVlBuxE9P4PTpVumqzlq9GhNcQ2s-H_E0oLw"
}
```

GetUserAccount

request:
```json
{
    "uuid": "9f11f2c8-05ef-444e-ac69-7bcf1b792160"
}
```

response:
```json
{
    "user_account": {
        "uuid": "9f11f2c8-05ef-444e-ac69-7bcf1b792160",
        "email": "jsmith@pm.me",
        "first_name": "john",
        "last_name": "smith",
        "email_verified": false
    },
    "status": "success"
}
```

## Rest API
The API also supports HTTP requests:
- 127.0.0.1:8080/api.user.v1.UserService/RegisterUser
- 127.0.0.1:8080/api.user.v1.UserService/LoginUser
- 127.0.0.1:8080/api.user.v1.UserService/GetUserAccount