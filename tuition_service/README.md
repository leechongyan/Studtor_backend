# Tuition Service

Tuition Service provides service for querying tuition matters:
- [x] Access Token must be present
- [x] Get all courses
- [x] Get all tutors
- [x] Get all tutors for a course
- [x] Save available time slots for a tutor
- [x] Delete available time slots for a tutor
- [x] Get all booked time slots for a tutor
- [x] Get all available time slots for a tutor
- [x] Book a time slot for a tutor
- [x] Unbook a time slot for a tutor
- [x] Get all booked time slots for a student

## API Endpoints
### API version 1

#### Get all the schools for filtering of courses

##### (GET) localhost:3000/v1/schools

Expected Returns:

```
[
    {
        "ID": 1,
        "SchoolCode": "NTU",
        "SchoolName": "Nanyang Technological University",
        "FacultyCodes": [
            "SCSE",
            "NBS",
            "NIE"
        ],
        "FacultyNames": [
            "Computer Science",
            "Business School",
            "Education"
        ]
    },
    {
        "ID": 2,
        "SchoolCode": "NUS",
        "SchoolName": "National University Singapore",
        "FacultyCodes": [
            "CS",
            "BIZ",
            "EDU"
        ],
        "FacultyNames": [
            "Computer Science",
            "Business School",
            "Education"
        ]
    }
]
```

#### Get all the courses

##### (GET) localhost:3000/v1/courses

Expected Returns:

```
[
    {
        "ID": 1,
        "CourseCode": "CZ1003",
        "CourseName": "Computational Thinking",
        "TutorSize": 10,
        "StudentSize": 15
    },
    {
        "ID": 2,
        "CourseCode": "CZ1004",
        "CourseName": "Computational System",
        "TutorSize": 10,
        "StudentSize": 15
    }
]
```

#### Get a courses

##### (GET) localhost:3000/v1/courses/:course_id

Expected Returns:

```
{
    "ID": 1,
    "CourseCode": "CZ1003",
    "CourseName": "Computational Thinking",
    "TutorSize": 10,
    "StudentSize": 15
}
```

#### Get the tutors for a course 

##### (GET) localhost:3000/v1/courses/:course_id/tutors

##### Optional Arguments: ?from_id=0&size=2

Expected Returns:

List of Tutors

#### Get a tutor for a course 

##### (GET) localhost:3000/v1/courses/:course_id/tutors/:user_id

Expected Returns:

A Tutor

#### Get all the courses taught by a tutor

##### (GET) localhost:3000/v1/tutors/:tutor_id/courses

Expected Returns:

```
[
    {
        "ID": 1,
        "CourseCode": "CZ1003",
        "CourseName": "Computational Thinking",
        "TutorSize": 10,
        "StudentSize": 15
    },
    {
        "ID": 2,
        "CourseCode": "CZ1004",
        "CourseName": "Computational System",
        "TutorSize": 10,
        "StudentSize": 15
    }
]
```

#### Register to teach a course

##### (POST) localhost:3000/v1/tutors/:tutor_id/courses/:course_id

Expected Returns:

```
"Success"
```

#### Deregister to teach a course

##### (DELETE) localhost:3000/v1/tutors/:tutor_id/courses/:course_id

Expected Returns:

```
"Success"
```

#### Put an availability

##### (POST) localhost:3000/v1/tutors/:tutor_id/availability

Request Body:

```
{
  "from": "2018-09-22T12:42:31Z",
  "to": "2018-09-25T12:50:31Z"
}
```

Expected Returns:

```
"Success"
```

#### Delete an availability

##### (DELETE) localhost:3000/v1/tutors/:tutor_id/availability

Request Body:

```
{
  "availability_id": 1
}
```

Expected Returns:

```
"Success"
```

#### Get the availability of a tutor

##### (GET) localhost:3000/v1/tutors/:tutor_id/availability

##### Optional Arguments: ?from=2018-09-22T12:42:31Z&to=2018-09-25T12:50:31Z

Expected Returns:

```
[
    {
        "ID": 1,
        "TutorID": 2,
        "AvailableFrom": 2018-09-22T12:42:31Z,
        "AvailableTo": 2018-09-25T12:50:31Z
    },
    {
        "ID": 2,
        "TutorID": 2,
        "AvailableFrom": 2018-09-22T12:42:31Z,
        "AvailableTo": 2018-09-25T12:50:31Z
    }
]
```

#### Book the availability of a tutor

##### (POST) localhost:3000/v1/courses/:course_id/tutors/:tutor_id/availability/:availability_id

Expected Returns:

```
"Success"
```

#### Delete a booking

##### (DELETE) localhost:3000/v1/users/:user_id/bookings/:booking_id

Expected Returns:

```
"Success"
```

#### Get all bookings

##### (GET) localhost:3000/v1/users/:user_id/bookings

Expected Returns:

```
[
    {
        "ID": 1,
        "CourseCode": "CZ1003",
        "CourseName": "Introduction to Computational Thinking",
        "TutorID": 1,
        "StudentID": 2,
        "StudentName": "Jeff",
        "FromTime": 2018-09-25T12:50:31Z,
        "ToTime": 2018-09-25T12:50:31Z
    },
    {
        "ID": 2,
        "CourseCode": "CZ1004",
        "CourseName": "Introduction to Computational System",
        "TutorID": 4,
        "StudentID": 2,
        "StudentName": "Jeff",
        "FromTime": 2018-09-25T12:50:31Z,
        "ToTime": 2018-09-25T12:50:31Z
    }
]
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.