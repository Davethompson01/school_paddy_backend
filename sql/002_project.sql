
CREATE TABLE paddyProject(
    project_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
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

ALTER TABLE paddyproject
RENAME COLUMN category_id TO project_id;

ALTER Table paddyproject add COLUMN student_id INTEGER REFERENCES students(user_id) NOT NULL;

ALTER TABLE paddyproject 
ALTER COLUMN student_id SET NOT NULL;

ALTER Table paddyproject add COLUMN discount_code VARCHAR(10);

ALTER TABLE paddyproject ADD COLUMN status TEXT NOT NULL;




ALTER TABLE paddyProject ADD COLUMN accepted_a_expert_already BOOLEAN default false;


ALTER TABLE applied_projects RENAME TO bid;
ALTER TABLE bid RENAME COLUMN applied_projects_id to bid_id;
CREATE TABLE bid(
    bid_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    student_id INTEGER REFERENCES students(user_id) NOT NULL,
    solution_expert_id INTEGER REFERENCES solution_expert(user_id) NOT NULL,
    project_id INTEGER REFERENCES paddyProject(project_id) NOT NULL,
    accepted BOOLEAN NOT NULL,
    isCompleted BOOLEAN NOT NULL,
    status VARCHAR(30) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL
)

ALTER Table bid add COLUMN Accepted_a_expert_already BOOLEAN default false;