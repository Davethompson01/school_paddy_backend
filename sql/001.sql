-- CREATE STUDENTS

CREATE TABLE students(
    user_id INTEGER GENERATED ALWAYS AS PRIMARY KEY,
    student_uuid int NOT NULL,
    name VARCHAR(100) NOT NULL,
    Phone_number VARCHAR(20) NOT NULL,
    role VARCHAR(10)
    Password TEXT NOT NULL,
    auth_method VARCHAR(30),
    update_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL
)


-- insert into sutdent
-- INSERT INTO students(use)