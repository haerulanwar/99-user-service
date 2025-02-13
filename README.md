# 99-user-service
REST Api for User Service, this is part of 99.co assessment

## How to run

### Required

- sqlite
- go 1.23.0

### Conf

To set up your environment, use existing config at `.env` or customizing the configurations accordingly.


### Run
```
$ go mod vendor 
$ go run cmd/main.go 
```
Now server is running at port 8080 (default port config).

### User Service
The user service stores information about all the users on the system. Fields available in the user object:

- `id (int)`: User ID _(auto-generated)_
- `name (str)`: Full name of the user _(required)_
- `created_at (int)`: Created at timestamp. In microseconds _(auto-generated)_
- `updated_at (int)`: Updated at timestamp. In microseconds _(auto-generated)_

#### APIs
##### Get all users
Returns all the users available in the db (sorted in descending order of creation date).

```
URL: GET /users

Parameters:
page_num = int # Default = 1
page_size = int # Default = 10
```
```json
Response:
{
    "result": true,
    "users": [
        {
            "id": 1,
            "name": "Suresh Subramaniam",
            "created_at": 1475820997000000,
            "updated_at": 1475820997000000,
        }
    ]
}
```

##### Get specific user
Retrieve a user by ID
```
URL: GET /users/{id}
```
```json
Response:
{
    "result": true,
    "user": {
        "id": 1,
        "name": "Suresh Subramaniam",
        "created_at": 1475820997000000,
        "updated_at": 1475820997000000,
    }
}
```

##### Create user
```
URL: POST /users
Content-Type: application/x-www-form-urlencoded

Parameters: (All parameters are required)
name = str
```
```json
Response:
{
    "result": true,
    "user": {
        "id": 1,
        "name": "Suresh Subramaniam",
        "created_at": 1475820997000000,
        "updated_at": 1475820997000000,
    }
}
```