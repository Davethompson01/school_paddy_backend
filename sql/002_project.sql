
CREATE TABLE paddyProject(
    category_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INTEGER NOT NULL,
    category TEXT NOT NULL,
    level VARCHAR(20) NOT NULL,
    topic TEXT NOT NULL,
    description TEXT NOT NULL,
    bidAmount DECIMAL(18,2) NOT NULL,
    deadline TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    update_at TIMESTAMP NOT NULL, 
    requirement TEXT NOT NULL
)

ALTER TABLE paddyproject DROP COLUMN user_id;
ALTER Table paddyproject add COLUMN user_id INTEGER REFERENCES students(user_id);

ALTER Table paddyproject add COLUMN discount_code VARCHAR(10);


ALTER TABLE students
ADD CONSTRAINT students_email_unique UNIQUE (email);

ALTER TABLE solution_expert
ADD CONSTRAINT solution_expert_email_unique UNIQUE (email);




SELECT EXISTS (
    SELECT 1 FROM students WHERE email = 'dthoadddndnnsddssd@example.com'
    UNION ALL
    SELECT 1 FROM solution_expert WHERE email = 'dthoadddndnnsddssd@example.com'
);