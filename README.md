# Studtor_backend

Server for Studtor_frontend 

## Set Up

Configure the config file:
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

## Usage
Go to terminal and cd into \Studtor_backend\cmd
```bash
go run main.go
```

## API Endpoints
### API version 1

#### Authentication Service (Refer to Readme in Authentication Service for more details)

##### (POST) localhost:3000/v1/auth/signup

##### (POST) localhost:3000/v1/auth/verify

##### (POST) localhost:3000/v1/auth/login

##### (POST) localhost:3000/v1/auth/logout

##### (POST) localhost:3000/v1/auth/refresh

##### (GET) localhost:3000/v1

##### (GET) localhost:3000/v1/:user

#### Tuition Service (Refer to Readme in Tuition Service for more details)

##### (GET) localhost:3000/v1/courses

##### (POST) localhost:3000/v1/putavailabletime

##### (POST) localhost:3000/v1/deleteavailabletime

##### (GET) localhost:3000/v1/tutors/*course

##### (GET) localhost:3000/v1/availabletime/:tutor

##### (GET) localhost:3000/v1/bookedtime/:user

##### (POST) localhost:3000/v1/book

##### (POST) localhost:3000/v1/unbook


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.