# Studtor_backend

Server for Studtor_frontend

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

mock_database: "true"
```

## Usage
Go to terminal and cd into \Studtor_backend\cmd
```bash
go run main.go
```

## API Endpoints
### API version 1

#### Authentication Service (Refer to Readme in Authentication Service for more details)

##### (POST) localhost:3000/v1/signup

##### (POST) localhost:3000/v1/verify

##### (POST) localhost:3000/v1/login

##### (POST) localhost:3000/v1/refresh

##### (GET) localhost:3000/v1/home/

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.