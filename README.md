# Studtor_backend


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
"Password": "password1",
"email": "clee051@e.ntu.edu.sg",
"user_type" : "USER"
}
```

Expected Returns:

```
{
"Success": "Successful Sign Up"
}
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
{
"Success": "Verified"
}
```
#### Login to get the access token
##### (POST) localhost:3000/v1/login

Request Body:
```
{
"first_name": "Jeff",
"last_name": "Lee",
"Password": "password1",
"email": "clee051@e.ntu.edu.sg",
"user_type" : "USER"
}
```
Expected Returns:
```
{
"first_name": "Jeff",
"last_name": "Lee",
"Password": "$2a$14$OBpe.WjBrAXut1UmBoZfSucSokfIaWZ3Y6Noxo4dcE/UkHma5i2AK",
"email": "clee051@e.ntu.edu.sg",
"token": "ACCESS_TOKEN",
"user_type": "USER",
"refresh_token": "REFRESH_TOKEN",
"verification_key": "838291",
"created_at": "2021-05-16T16:19:32+08:00",
"updated_at": "2021-05-16T16:19:32+08:00"
}
```
#### Access authorized pages
##### (GET) localhost:3000/v1/home/

Request Header:

Token: "Bearer: ACCESS_TOKEN"

Expected Returns:
```
{
"Success": "Successful Entry"
}
```
## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
