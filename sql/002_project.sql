
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


ALTER Table paddyproject add COLUMN student_id INTEGER REFERENCES students(user_id) NOT NULL;

ALTER TABLE paddyproject 
ALTER COLUMN student_id SET NOT NULL;

ALTER Table paddyproject add COLUMN discount_code VARCHAR(10);







ALTER TABLE paddyProject ADD COLUMN accepted_a_expert_already BOOLEAN default false;