SELECT * FROM users;
 id | age | first_name  | last_name |       email       
----+-----+-------------+-----------+-------------------
  3 |  56 | don         | planty    | don@planty.com
  2 |  25 | young-blood | Santiago  | john@santiago.com
  1 |  22 | young-blood | Bunyan    | paul@santiago.com
(3 rows)

DELETE FROM users WHERE first_name='don';

RESULT: 

SELECT * FROM users;
 id | age | first_name  | last_name |       email       
----+-----+-------------+-----------+-------------------
  2 |  25 | young-blood | Santiago  | john@santiago.com
  1 |  22 | young-blood | Bunyan    | paul@santiago.com
(2 rows)


Note: IDs will keep increasing when you insert new items. 

INSERT INTO users(age,email,first_name,last_name)
VALUES(25, 'don@planty.com', 'don', 'planty');

Result: re-inserting makes the id to 4

 id | age | first_name  | last_name |       email       
----+-----+-------------+-----------+-------------------
  2 |  25 | young-blood | Santiago  | john@santiago.com
  1 |  22 | young-blood | Bunyan    | paul@santiago.com
  4 |  25 | don         | planty    | don@planty.com