CREATE TABLE IF NOT EXISTS TASK(
    id serial PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    due timestamp,
    done boolean NOT NULL,
    createdat timestamp NOT NULL,
    updatedat timestamp
);

insert into task (title, description, done, createdat)
values ('read go documentation', 'read go docs to figure how to test containers', false, current_timestamp);