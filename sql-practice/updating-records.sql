BEFORE UPDATE:

SELECT * FROM users;
 id | age | first_name | last_name |       email       
----+-----+------------+-----------+-------------------
  1 |  22 | Paul       | Santiago  | paul@santiago.com
  2 |  25 | John       | Santiago  | john@santiago.com
  3 |  56 | don        | planty    | don@planty.com


UPDATE users SET first_name = 'Pauly', last_name ='Bunyan' WHERE id=1;

AFTER UPDATE:

id | age | first_name | last_name |       email       
----+-----+------------+-----------+-------------------
  2 |  25 | John       | Santiago  | john@santiago.com
  3 |  56 | don        | planty    | don@planty.com
  1 |  22 | Pauly      | Bunyan    | paul@santiago.com


UPDATING MULTIPLE users

BEFORE MULTI UPDATE
id | age | first_name | last_name |       email       
----+-----+------------+-----------+-------------------
  2 |  25 | John       | Santiago  | john@santiago.com
  3 |  56 | don        | planty    | don@planty.com
  1 |  22 | Pauly      | Bunyan    | paul@santiago.com


  UPDATE users SET first_name='young-blood' WHERE age < 50;

AFTER MULTI UPDATE:

  id | age | first_name  | last_name |       email       
----+-----+-------------+-----------+-------------------
  3 |  56 | don         | planty    | don@planty.com
  2 |  25 | young-blood | Santiago  | john@santiago.com
  1 |  22 | young-blood | Bunyan    | paul@santiago.com