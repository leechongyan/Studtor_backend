# Studtor_backend

Server for Studtor_frontend 

## Set Up

Configure the config file:
* Please change the server email and password accordingly
* Set expiration time in hours
```yml
jwtKey: "9761278367815487"
accessExpirationTime: 1
refreshExpirationTime: 2
serverEmail: "studtorr@gmail.com"
serverEmailPW: "password"

database_config: "host=localhost user=postgres password=? dbname=? port=5432 sslmode=disable TimeZone=Asia/Shanghai"

google_bucket_name: "studtor"
mock_database: false
mock_storage: false
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

##### (POST) localhost:3000/v1/auth/refresh

##### (POST) localhost:3000/v1/user/logout

##### (GET) localhost:3000/v1/user

##### (GET) localhost:3000/v1/user/:user_id

#### Tuition Service (Refer to Readme in Tuition Service for more details)

##### (GET) localhost:3000/v1/schools

##### (GET) localhost:3000/v1/courses

##### (GET) localhost:3000/v1/courses/:course_id

##### (GET) localhost:3000/v1/courses/:course_id/tutors

##### (GET) localhost:3000/v1/courses/:course_id/tutors/:user_id

##### (GET) localhost:3000/v1/tutors/:tutor_id/courses

##### (POST) localhost:3000/v1/tutors/:tutor_id/courses/:course_id

##### (DELETE) localhost:3000/v1/tutors/:tutor_id/courses/:course_id

##### (POST) localhost:3000/v1/tutors/:tutor_id/availability

##### (DELETE) localhost:3000/v1/tutors/:tutor_id/availability/:availability_id

##### (GET) localhost:3000/v1/tutors/:tutor_id/availability

##### (POST) localhost:3000/v1/courses/:course_id/tutors/:tutor_id/availability/:availability_id

##### (DELETE) localhost:3000/v1/users/:user_id/bookings/:booking_id

##### (GET) localhost:3000/v1/user/bookings

##### (GET) localhost:3000/v1/tutors/1/bookings

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
