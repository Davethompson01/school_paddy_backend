-- CREATE STUDENTS
CREATE TABLE solution_expert(
    user_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email TEXT NOT NULL,
    Phone_number VARCHAR(20) NOT NULL,
    role VARCHAR(10),
    Password TEXT NOT NULL,
    auth_method VARCHAR(30),
    update_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP  NOT NULL DEFAULT NOW()
)


ALTER TABLE solution_expert ADD COLUMN email_verified BOOLEAN DEFAULT false;

