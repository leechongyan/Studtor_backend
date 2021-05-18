# Studtor_backend

![Alt text](https://github.com/leechongyan/Studtor_backend/tree/database_interface/images/workflow.jpg?raw=true "Authentication Flow")

## Set Up

Configure the config file:
* Please change the server email and password accordingly
* Set expiration time in hours
```yml
port: ":3000"

jwtKey: "9761278367815487"
expirationTime: "10"
serverEmail: "studtorr@gmail.com"
serverEmailPW: "password"
```

## Usage
Go to terminal and cd into \Studtor_backend\cmd
```bash
go run main.go
```

## API Endpoints
### API version 1

#### Sign up for an account

##### (POST) localhost:3000/v1/signup

Request Body:

```
{
"first_name": "Jeff",
"last_name": "Lee",
"password": "password1",
"email": "clee051@e.ntu.edu.sg",
"user_type" : "USER"
}
```

Expected Returns:

```
"Success"
```

#### Verify with verification code
##### (POST) localhost:3000/v1/verify

Request Body:

```
{
"email": "clee051@e.ntu.edu.sg",
"verification_key": "838291"
}
```

Expected Returns:

```
"Success"
```

#### Login to get the access token
##### (POST) localhost:3000/v1/login

Request Body:

```
{
"email": "clee051@e.ntu.edu.sg",
"password": "password1"
}
```

Expected Returns:

```
"Access Token"
```

#### Refresh access token
##### (POST) localhost:3000/v1/refresh

Request Header:

Token: "Bearer: REFRESH_TOKEN"

Request Body:

```
"clee051@e.ntu.edu.sg"
```

Expected Returns:

```
"Access Token"
```

#### Access authorized pages
##### (GET) localhost:3000/v1/home/

Request Header:

Token: "Bearer: ACCESS_TOKEN"

Expected Returns:

```
"Success"
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.