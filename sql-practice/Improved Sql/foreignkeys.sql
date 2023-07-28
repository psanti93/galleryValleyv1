CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE REFERENCES users (id) , --one option of creating foreign keys
    token_hash TEXT UNIQUE NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users (id) --second option
);

-- will tell our db that users id are entered intot eh sessions table map to real users objects
-- if we delete a user they can't have a session associated with them


-- altering an existing table to add another constraint
ALTER TABLE sessions(
    ADD CONSTRAINT sessions_user_id_fkey FOREIGN KEY(user_id) REFERENCES users (id);
)