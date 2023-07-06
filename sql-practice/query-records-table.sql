SELECT * FROM users;
Result: gets everything
id | age | first_name | last_name |       email
----+-----+------------+-----------+-------------------
  1 |  22 | Paul       | Santiago  | paul@santiago.com
  2 |  25 | John       | Santiago  | john@santiago.com

choosing specific fields
SELECT id, email FROM users;
Result: 
id |       email       
----+-------------------
  1 | paul@santiago.com
  2 | john@santiago.com
(2 rows)