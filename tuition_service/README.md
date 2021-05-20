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

#### Get all courses from a specified course code up to a specified size

##### (GET) localhost:3000/v1/courses

##### Optional Arguments: ?from=0&size=2

Expected Returns:

```
[
    "CZ1001",
    "CZ2001"
]
```

#### Save an available time slot for a tutor

##### (POST) localhost:3000/v1/home/putavailabletime

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

#### Delete an available time slot for a tutor

##### (POST) localhost:3000/v1/home/deleteavailabletime

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

#### Get all tutors for a particular course from a specified tutor ID up to a specified size

##### (GET) localhost:3000/v1/tutors/*course

##### Optional Arguments: ?from_id=0&size=2

Expected Returns:

```
[
    {
        "code": "CZ1003",
        "students": 10,
        "title": "Computational Thinking",
        "tutors": 15
    },
    {
        "code": "CZ3003",
        "students": 4,
        "title": "Object Thinking",
        "tutors": 20
    }
]
```

#### Get all available time slot for a tutor with a start date and end date

##### (GET) localhost:3000/v1/availabletime/:tutor

##### Optional Arguments: ?from=2018-09-25T12:50:31Z&to=2018-09-25T12:50:31Z

Expected Returns:

```
{
    "email": "clee051@e.ntu.edu.sg",
    "first_name": "Jeff",
    "last_name": "Lee",
    "time_slots": [
        [
            "2021-05-19T16:39:05.9712695+08:00",
            "2021-05-19T16:39:05.9712695+08:00"
        ],
        [
            "2021-05-19T16:39:05.9712695+08:00",
            "2021-05-19T16:39:05.9712695+08:00"
        ]
    ]
}
```

#### Get all the booked times for a user

##### (GET) localhost:3000/v1/bookedtime/:user

##### Optional Arguments: ?from=2018-09-25T12:50:31Z&to=2018-09-25T12:50:31Z

Expected Returns:

```
{
    "email": "clee051@e.ntu.edu.sg",
    "first_name": "Jeff",
    "last_name": "Lee",
    "time_slots": {
        "CZ1003": [
            "2021-05-20T00:16:03.2733615+08:00",
            "2021-05-20T00:16:03.2733615+08:00"
        ],
        "CZ1004": [
            "2021-05-20T00:16:03.2733615+08:00",
            "2021-05-20T00:16:03.2733615+08:00"
        ]
    }
}
```

#### Book the time of a tutor

##### (POST) localhost:3000/v1/book

Request Body:

```
{
  "course": "CZ1003",
  "from": "2018-09-22T12:42:31Z",
  "to": "2018-09-25T12:50:31Z",
  "tutor": "Jeff"
}
```

Expected Returns:

```
"Success"
```

#### Unbook the time of a tutor

##### (POST) localhost:3000/v1/unbook

Request Body:

```
{
  "course": "CZ1003",
  "from": "2018-09-22T12:42:31Z",
  "to": "2018-09-25T12:50:31Z",
  "tutor": "Jeff"
}
```

Expected Returns:

```
"Success"
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.