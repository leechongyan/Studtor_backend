# Authentication Service

Authentication Service provides authentication service for Studtor backend with the following requirements:
- [x] Only edu email is accepted
- [x] Verification code sent to edu email to validate account
- [x] Access Token and Refresh Token will be generated after login
- [x] Access Token will be used to access authorized pages
- [x] Refresh Token will be used to get a new Access Token if existing Access Token has expired
- [x] Refresh Token will be deleted after logout

![Alt text](https://github.com/leechongyan/Studtor_backend/blob/database_interface/images/workflow.JPG "Authentication Flow")

## Set Up

Configure the config file:
* These fields are required
* Please change the server email and password accordingly
* Set expiration time in hours
```yml
port: ":3000"

jwtKey: "9761278367815487"
accessExpirationTime: "1"
refreshExpirationTime: "2"
serverEmail: "studtorr@gmail.com"
serverEmailPW: "password"

mock_database: "true"
```


## API Endpoints
### API version 1

#### Sign up for an account

##### (POST) localhost:3000/v1/auth/signup

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
##### (POST) localhost:3000/v1/auth/verify

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
##### (POST) localhost:3000/v1/auth/login

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

#### Logout 
##### (POST) localhost:3000/v1/auth/logout

Request Body:

```
"clee051@e.ntu.edu.sg"
```

Expected Returns:

```
"Success"
```

#### Refresh access token
##### (POST) localhost:3000/v1/auth/refresh

Request Body:

```
"clee051@e.ntu.edu.sg"
```

Expected Returns:

```
"Access Token"
```

#### Access authorized pages
This is for testing authority
##### (GET) localhost:3000/v1

Request Header:

Token: "Bearer: ACCESS_TOKEN"

Expected Returns:

```
"Success"
```

#### Get User Information
If no user id is specified, then the current user will be returned
##### (GET) localhost:3000/v1/user/*user

Request Header:

Token: "Bearer: ACCESS_TOKEN"

Expected Returns:

```
{
    "first_name": "Jeff",
    "id": "10",
    "last_name": "Lee"
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.