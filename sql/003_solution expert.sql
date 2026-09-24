-- CREATE STUDENTS
CREATE TABLE solution_expert (
    user_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    name VARCHAR(100) NOT NULL,

    email TEXT NOT NULL,

    phone_number VARCHAR(20) NOT NULL,

    role VARCHAR(30),

    password TEXT NOT NULL,

    auth_method VARCHAR(30),

    update_at TIMESTAMP NOT NULL DEFAULT NOW(),

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);


ALTER TABLE solution_expert
ALTER COLUMN auth_method TYPE VARCHAR(30);

ALTER TABLE solution_expert ADD COLUMN email_verified BOOLEAN DEFAULT false;

ALTER TABLE solution_expert
ADD CONSTRAINT solution_expert_email_unique UNIQUE (email);

DROP TABLE solution_expert;