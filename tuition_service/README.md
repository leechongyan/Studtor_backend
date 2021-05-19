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

##### (GET) localhost:3000/v1/home/getallcourses/?from=0&size=2

Expected Returns:

```
[
    "CZ1001",
    "CZ2001"
]
```

#### Get all tutors from a specified tutor ID up to a specified size

##### (GET) localhost:3000/v1/home/getalltutors/?from=0&size=2

Expected Returns:

```
[
    "Chin",
    "Kangyu"
]
```

#### Get all tutors for a particular course from a specified tutor ID up to a specified size

##### (GET) localhost:3000/v1/home/getalltutorsforacourse/?course_code=CZ1003&from=0&size=2

Expected Returns:

```
[
    "Chin",
    "Kangyu"
]
```

#### Save an available time slot for a tutor

##### (POST) localhost:3000/v1/home/putavailabletimetutor

Request Body:

```
{
  "email": "clee051@e.ntu.edu.sg",
  "from": "2018-09-22T12:42:31Z",
  "to": "2018-09-25T12:50:31Z"
}
```

Expected Returns:

```
"Success"
```

#### Delete an available time slot for a tutor

##### (POST) localhost:3000/v1/home/deleteavailabletimetutor

Request Body:

```
{
  "email": "clee051@e.ntu.edu.sg",
  "from": "2018-09-22T12:42:31Z",
  "to": "2018-09-25T12:50:31Z"
}
```

Expected Returns:

```
"Success"
```

#### Get all the booked times for a tutor

##### (POST) localhost:3000/v1/home/getallbookedtimetutor

Request Body:

```
{
  "email": "clee051@e.ntu.edu.sg",
  "from": "2018-09-22T12:42:31Z",
  "to": "2018-09-25T12:50:31Z"
}
```

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

#### Get all available time slot for a tutor with a start date and end date

##### (POST) localhost:3000/v1/home/getavailabletimetutor

Request Body:

```
{
  "email": "clee051@e.ntu.edu.sg",
  "from": "2018-09-22T12:42:31Z",
  "to": "2018-09-25T12:50:31Z"
}
```

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

#### Book the time of a tutor

##### (POST) localhost:3000/v1/home/booktimetutor

Request Body:

```
{
  "email": "clee051@e.ntu.edu.sg",
  "from": "2018-09-22T12:42:31Z",
  "to": "2018-09-25T12:50:31Z"
}
```

Expected Returns:

```
"Success"
```

#### Unbook the time of a tutor

##### (POST) localhost:3000/v1/home/unbooktimetutor

Request Body:

```
{
  "email": "clee051@e.ntu.edu.sg",
  "from": "2018-09-22T12:42:31Z",
  "to": "2018-09-25T12:50:31Z"
}
```

Expected Returns:

```
"Success"
```

#### Get all the booked times of a student

##### (POST) localhost:3000/v1/home/getallbookedtimestudent

Request Body:

```
{
  "email": "clee051@e.ntu.edu.sg",
  "from": "2018-09-22T12:42:31Z",
  "to": "2018-09-25T12:50:31Z"
}
```

Expected Returns:

```
{
    "email": "clee051@e.ntu.edu.sg",
    "first_name": "Jeff",
    "last_name": "Lee",
    "time_slots": {
        "CZ1003": [
            "2021-05-20T00:17:09.8745778+08:00",
            "2021-05-20T00:17:09.8745778+08:00"
        ],
        "CZ1004": [
            "2021-05-20T00:17:09.8745778+08:00",
            "2021-05-20T00:17:09.8745778+08:00"
        ]
    }
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.