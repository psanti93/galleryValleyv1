INSERT INTO users VALUES (1, 22, 'Paul', 'Santiago', 'paul@santiago.com');

Result:
INSERT 0 1
galleyvalley=# select * from users;
 id | age | first_name | last_name |       email       
----+-----+------------+-----------+-------------------
  1 |  22 | Paul       | Santiago  | paul@santiago.com
(1 row)


INSERT INTO users(age,email,first_name,last_name) VALUES (25, 'john@santiago.com', 'John','Santiago');

Result:
INSERT 0 1     
galleyvalley=# select * from users;
 id | age | first_name | last_name |       email       
----+-----+------------+-----------+-------------------
  1 |  22 | Paul       | Santiago  | paul@santiago.com
  2 |  25 | John       | Santiago  | john@santiago.com
  (2 rows)