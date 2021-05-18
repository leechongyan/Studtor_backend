-- Students table
CREATE TABLE users (
    student_id INTEGER PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    password TEXT,
    email TEXT,
    token TEXT,
    user_type TEXT,
    refresh_token TEXT,
    v_key TEXT,
    verified INTEGER,
    created_at TEXT,
    updated_at TEXT
);

-- Courses table
CREATE TABLE courses (
    course_id  INTEGER PRIMARY KEY,
    course_code TEXT,
    course_name TEXT
);

-- Tutors table
CREATE TABLE tutors (
    tutor_id  INTEGER PRIMARY KEY,
    tutor_name TEXT
);