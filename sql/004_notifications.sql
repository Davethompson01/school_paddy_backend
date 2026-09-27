

CREATE TABLE notification(
    notification_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    student_id INTEGER REFERENCES students(user_id) NOT NULL,
    solution_id INTEGER REFERENCES solution_expert(user_id) NOT NULL,
    message TEXT NOT NULL, 
    bid_id INTEGER REFERENCES bid(bid_id),
    created_at TIMESTAMP DEFAULT NOW() NOT NULL
);


Alter Table notification add COLUMN project_id INTEGER REFERENCES paddyproject(project_id)

ALTER TABLE