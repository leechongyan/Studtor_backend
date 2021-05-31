# Database Service

Database Service provides service for saving data.
Supported Connectors:
- [x] User
- [x] Tutor
- [x] Course
- [x] Availability
- [x] Booking
- [x] School

## Usage
### API version 1
#### User Connector
```go
// need to import client models as we will be interacting using client models with the database
import(
	userModel "github.com/leechongyan/Studtor_backend/database_service/client_models"
	userConnector "github.com/leechongyan/Studtor_backend/database_service/connector/user_connector"
)

// call init to get the connector 
connector := userConnector.Init()

// to retrieve user by email
user, err := userConnector.Init().SetUserEmail(email).GetUser()

// to retrieve userprofile by email
// profile has less fields exposed for security
userProfile, err := userConnector.Init().SetUserEmail(email).GetProfile()

// to retrieve user by id
user, err := userConnector.Init().SetUserId(id).GetUser()

// to retrieve userprofile by id
// profile has less fields exposed for security
userProfile, err := userConnector.Init().SetUserId(id).GetProfile()

// to save user (update or create will be determined by connector, which will check whether there is an existing entry)
// make sure that user object has email field, which is needed to check existing entry
err := userConnector.Init().SetUser(user).Add()

// to delete user by email
err := userConnector.Init().SetUserEmail(email).Delete()

// to delete user by id
err := userConnector.Init().SetUserId(id).Delete()
```

#### Tutor Connector
```go
import(
	tutorConnector "github.com/leechongyan/Studtor_backend/database_service/connector/tutor_connector"
)

// call init to get the connector 
connector := tutorConnector.Init()

// to register interest in teaching a course
err := tutorConnector.Init().SetTutorId(1).SetCourseId(2).Add()

// to deregister interest in teaching a course
err := tutorConnector.Init().SetTutorId(1).SetCourseId(2).Delete()

// to get all tutors teaching for a course
tutors, err := tutorConnector.Init().SetCourseId(2).SetSize(5).SetTutorId(1).GetAll()
```

#### Course Connector
Add and Delete Operations are currently not implemented
```go
import(
	courseConnector "github.com/leechongyan/Studtor_backend/database_service/connector/course_connector"
)

// call init to get the connector 
connector := courseConnector.Init()

// get a single course
course, err := courseConnector.Init().SetCourseId(1).GetSingle()

// get all courses
courses, err := courseConnector.Init().GetAll()

// get all courses taught by a tutor
courses, err := courseConnector.Init().SetTutorId(2).GetAll()
```

#### Availability Connector
```go
import(
	availabilityConnector "github.com/leechongyan/Studtor_backend/database_service/connector/availability_connector"
)

// call init to get the connector 
connector := availabilityConnector.Init()

// get all availability for a tutor
availabilities, err := availabilityConnector.Init().SetTutorId(1).SetFromTime("Set a from time").SetToTime("Set a to time").GetAll()

// add an availability for a tutor
err := availabilityConnector.Init().SetTutorId(1).SetFromTime("Set a from time").SetToTime("Set a to time").Add()

// delete an availability of a tutor
err := availabilityConnector.Init().SetTutorId(1).SetAvailabilityId(5).Delete()
```

#### Booking Connector
```go
import(
	bookingConnector "github.com/leechongyan/Studtor_backend/database_service/connector/booking_connector"
)

// call init to get the connector 
connector := bookingConnector.Init()

// get all booking for a user
bookings, err := bookingConnector.Init().SetUserId(1).SetFromTime("Set a from time").SetToTime("Set a to time").GetAll()

// add a booking for a student
err := bookingConnector.Init().SetUserId(1).SetCourseId(2).SetAvailabilityId(3).Add()

// delete a booking
err := bookingConnector.Init().SetUserId(1).SetBookingId(5).Delete()
```

#### School Connector
```go
import(
	schoolConnector "github.com/leechongyan/Studtor_backend/database_service/connector/school_connector"
)

// call init to get the connector 
connector := schoolConnector.Init()

// get all schools and their courses 
schools, err := schoolConnector.Init().GetAll()
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.