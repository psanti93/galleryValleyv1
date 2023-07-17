SELECT * FROM users WHERE email='paul@santiago.com';
 
 Result: 
 id | age | first_name | last_name |       email       
----+-----+------------+-----------+-------------------
  1 |  22 | Paul       | Santiago  | paul@santiago.com


SELECT * FROM users WHERE age > 22;

Result: 
 id | age | first_name | last_name |       email       
----+-----+------------+-----------+-------------------
  2 |  25 | John       | Santiago  | john@santiago.com


Example of using OR:
SELECT * FROM users WHERE age > 22 OR last_name='planty';
 
 Result: 
 id | age | first_name | last_name |       email       
----+-----+------------+-----------+-------------------
  2 |  25 | John       | Santiago  | john@santiago.com
  3 |  56 | don        | planty    | don@planty.com


Example of no result:
SELECT * FROM users WHERE age > 22 AND  last_name='james';
 
 Result:
 id | age | first_name | last_name | email 
----+-----+------------+-----------+-------


Example of using LIMIT
SELECT * FROM users WHERE age > 22 OR last_name='planty'LIMIT 1;

Result: 
 id | age | first_name | last_name |       email       
----+-----+------------+-----------+-------------------
  2 |  25 | John       | Santiago  | john@santiago.com
