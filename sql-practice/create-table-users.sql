CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    age INT,
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE NOT NULL
);

Result:
galleyvalley=# SELECT * FROM users;
 id | age | first_name | last_name | email 
----+-----+------------+-----------+-------
(0 rows)