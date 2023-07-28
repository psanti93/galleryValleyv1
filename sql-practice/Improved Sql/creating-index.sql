CREATE INDEX sessions_token_hash_idx ON sessions(token_hash,user_id,id);

CREATE TABLE dogs (
    id SERIAL PRIMARY KEY,
    name TEXT,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    INDEX dogs_user_id_idx(user_id) --doesn't work on postgresql
)

CREATE TABLE dogs (
    id SERIAL PRIMARY KEY,
    name TEXT,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
   
)
