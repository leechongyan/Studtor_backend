-- Students table
CREATE TABLE students (
    student_id INTEGER PRIMARY KEY,
    student_name TEXT
);

-- Tutors table
CREATE TABLE tutors (
    tutor_id INTEGER PRIMARY KEY,
    tutor_name TEXT
);

-- Courses table
CREATE TABLE courses (
    course_id  INTEGER PRIMARY KEY,
    course_code TEXT,
    course_name TEXT
);