CREATE DATABASE student_book_course;

\c student_book_course

CREATE TABLE student (
    student_id BIGSERIAL NOT NULL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(200)
);

CREATE TABLE book (
    book_id BIGSERIAL NOT NULL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    author VARCHAR(100) NOT NULL
);

CREATE TABLE student_book (
    student_id BIGSERIAL REFERENCES student(student_id),
    book_id BIGSERIAL REFERENCES book(book_id)
);

ALTER TABLE student_book ADD CONSTRAINT unique_borrow UNIQUE(student_id, book_id);

CREATE TABLE course (
    course_id BIGSERIAL NOT NULL PRIMARY KEY,
    instructor VARCHAR(200) NOT NULL,
    title VARCHAR(200) NOT NULL
);

CREATE TABLE student_course (
    student_id BIGSERIAL REFERENCES student(student_id),
    course_id BIGSERIAL REFERENCES course(course_id)
);

ALTER TABLE student_course ADD CONSTRAINT unique_enroll UNIQUE(student_id, course_id);