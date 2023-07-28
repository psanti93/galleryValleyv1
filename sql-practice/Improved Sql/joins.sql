-- Inner JOIN example
SELECT * FROM users
 JOIN sessions ON users.id = sessions.user_id; -- match id on users table to sessions' user_id field --JOIN is known as INNER JOIN 


select * from users;
 id |         email         |                        password_hash
----+-----------------------+--------------------------------------------------------------
  1 | paul@santiago1.com    | $2a$10$Bx7LIRtWfh5X4hWW5hvm6.Wf0MDLg27ZQSot2uP.JWcZS4n7VeI5S
  2 | em@velasquez1.com     | $2a$10$AjxUaiyL3nsvYwHKbOJ.Qe5ZkzL356fmVfWtastb9gVWLMJs1AvTa
  3 | john@santiago1.com    | $2a$10$RSJoPc9KMhDy6v0Vl7HJIeKfxpFcylrAiO4sfjjZpHJx1SjLDL/dS
  4 | francis@santiago1.com | $2a$10$tqgW/gRc/AMt31Fz0MKcH.T0jOUBKQTAvxGJI4jAJ6nkst8oG2S0G
  5 | thea@gonazalez1.com   | $2a$10$H0ySs8tx71tIMfYj8JcnKu0t0A4H5IRf0vRqSQiEBVIVz4j1PRlSa
(5 rows)

select * from sessions;
 id | user_id |                  token_hash
----+---------+----------------------------------------------
  1 |       1 | 11irIAh40anc8lBeadXW4u7NS03mESI8Rcb3JWGolVM=
  2 |       2 | bYo1In4M57JqmtnJ2r8pq0N6gOaz8VfwEn3r9Njr6lI=
  3 |       3 | S4QdT2F7G6RKX1YHvCwOLeGcd3_cJyEiyI9J_nauKDI=
  4 |       4 | RM0UNytkiME94GlM5Yv_peK-jV0jpI-beleSEi9FT5E=
  5 |       5 | rGdpE1xwL1DGytIrYLllVRDIJPr4OxM3GwN8xQG1hh0=
(5 rows)

RESULT: 

 id |         email         |                        password_hash                         | id | user_id |                  token_hash
----+-----------------------+--------------------------------------------------------------+----+---------+----------------------------------------------
  1 | paul@santiago1.com    | $2a$10$Bx7LIRtWfh5X4hWW5hvm6.Wf0MDLg27ZQSot2uP.JWcZS4n7VeI5S |  1 |       1 | 11irIAh40anc8lBeadXW4u7NS03mESI8Rcb3JWGolVM=
  2 | em@velasquez1.com     | $2a$10$AjxUaiyL3nsvYwHKbOJ.Qe5ZkzL356fmVfWtastb9gVWLMJs1AvTa |  2 |       2 | bYo1In4M57JqmtnJ2r8pq0N6gOaz8VfwEn3r9Njr6lI=
  3 | john@santiago1.com    | $2a$10$RSJoPc9KMhDy6v0Vl7HJIeKfxpFcylrAiO4sfjjZpHJx1SjLDL/dS |  3 |       3 | S4QdT2F7G6RKX1YHvCwOLeGcd3_cJyEiyI9J_nauKDI=
  4 | francis@santiago1.com | $2a$10$tqgW/gRc/AMt31Fz0MKcH.T0jOUBKQTAvxGJI4jAJ6nkst8oG2S0G |  4 |       4 | RM0UNytkiME94GlM5Yv_peK-jV0jpI-beleSEi9FT5E=
  5 | thea@gonazalez1.com   | $2a$10$H0ySs8tx71tIMfYj8JcnKu0t0A4H5IRf0vRqSQiEBVIVz4j1PRlSa |  5 |       5 | rGdpE1xwL1DGytIrYLllVRDIJPr4OxM3GwN8xQG1hh0=
(5 rows)


 -- returns data that is common on both tables

 -- LEFT JOIN 
 -- all the items that we're starting with (users) and only records that we're joining (sessions) 
SELECT * FROM users
     LEFT JOIN sessions ON users.id = sessions.user_id;

 id |         email         |                        password_hash
----+-----------------------+--------------------------------------------------------------
  1 | paul@santiago1.com    | $2a$10$Bx7LIRtWfh5X4hWW5hvm6.Wf0MDLg27ZQSot2uP.JWcZS4n7VeI5S
  2 | em@velasquez1.com     | $2a$10$AjxUaiyL3nsvYwHKbOJ.Qe5ZkzL356fmVfWtastb9gVWLMJs1AvTa
  3 | john@santiago1.com    | $2a$10$RSJoPc9KMhDy6v0Vl7HJIeKfxpFcylrAiO4sfjjZpHJx1SjLDL/dS
  5 | thea@gonazalez1.com   | $2a$10$H0ySs8tx71tIMfYj8JcnKu0t0A4H5IRf0vRqSQiEBVIVz4j1PRlSa
  6 | jim@dundermifflin.com | $2a$10$BCBfSuYzZaRXf79IFlsfTOtAugBnJD/Q/3Lf/ZSHkXIQiZbJrwlU6
  7 | pam@dundermifflin.com | $2a$10$cy2xcLuQT7UmCEQXw3ff0uG0Ilt4I4gIzRM/QGIDGsO4tmzUI5xUi



select * from sessions;
 id | user_id |                  token_hash
----+---------+----------------------------------------------
  1 |       1 | 11irIAh40anc8lBeadXW4u7NS03mESI8Rcb3JWGolVM=
  2 |       2 | bYo1In4M57JqmtnJ2r8pq0N6gOaz8VfwEn3r9Njr6lI=
  3 |       3 | S4QdT2F7G6RKX1YHvCwOLeGcd3_cJyEiyI9J_nauKDI=
  5 |       5 | rGdpE1xwL1DGytIrYLllVRDIJPr4OxM3GwN8xQG1hh0=



RESULT: 
--Joined table returns back all of the users even if they don't have a session

 id |         email         |                        password_hash                         | id | user_id |                  token_hash
----+-----------------------+--------------------------------------------------------------+----+---------+----------------------------------------------
  1 | paul@santiago1.com    | $2a$10$Bx7LIRtWfh5X4hWW5hvm6.Wf0MDLg27ZQSot2uP.JWcZS4n7VeI5S |  1 |       1 | 11irIAh40anc8lBeadXW4u7NS03mESI8Rcb3JWGolVM=
  2 | em@velasquez1.com     | $2a$10$AjxUaiyL3nsvYwHKbOJ.Qe5ZkzL356fmVfWtastb9gVWLMJs1AvTa |  2 |       2 | bYo1In4M57JqmtnJ2r8pq0N6gOaz8VfwEn3r9Njr6lI=
  3 | john@santiago1.com    | $2a$10$RSJoPc9KMhDy6v0Vl7HJIeKfxpFcylrAiO4sfjjZpHJx1SjLDL/dS |  3 |       3 | S4QdT2F7G6RKX1YHvCwOLeGcd3_cJyEiyI9J_nauKDI=
  5 | thea@gonazalez1.com   | $2a$10$H0ySs8tx71tIMfYj8JcnKu0t0A4H5IRf0vRqSQiEBVIVz4j1PRlSa |  5 |       5 | rGdpE1xwL1DGytIrYLllVRDIJPr4OxM3GwN8xQG1hh0=
  6 | jim@dundermifflin.com | $2a$10$BCBfSuYzZaRXf79IFlsfTOtAugBnJD/Q/3Lf/ZSHkXIQiZbJrwlU6 |    |         |
  7 | pam@dundermifflin.com | $2a$10$cy2xcLuQT7UmCEQXw3ff0uG0Ilt4I4gIzRM/QGIDGsO4tmzUI5xUi |    |         | 

-- RIGHT JOIN

SELECT * FROM users
     RIGHT JOIN sessions ON users.id = sessions.user_id;


RESULT
-- returns data from the right table (sessions) that have existing mapping on the left handside (left)

id |        email        |                        password_hash                         | id | user_id |                  token_hash
----+---------------------+--------------------------------------------------------------+----+---------+----------------------------------------------
  1 | paul@santiago1.com  | $2a$10$Bx7LIRtWfh5X4hWW5hvm6.Wf0MDLg27ZQSot2uP.JWcZS4n7VeI5S |  1 |       1 | 11irIAh40anc8lBeadXW4u7NS03mESI8Rcb3JWGolVM=
  2 | em@velasquez1.com   | $2a$10$AjxUaiyL3nsvYwHKbOJ.Qe5ZkzL356fmVfWtastb9gVWLMJs1AvTa |  2 |       2 | bYo1In4M57JqmtnJ2r8pq0N6gOaz8VfwEn3r9Njr6lI=
  3 | john@santiago1.com  | $2a$10$RSJoPc9KMhDy6v0Vl7HJIeKfxpFcylrAiO4sfjjZpHJx1SjLDL/dS |  3 |       3 | S4QdT2F7G6RKX1YHvCwOLeGcd3_cJyEiyI9J_nauKDI=
  5 | thea@gonazalez1.com | $2a$10$H0ySs8tx71tIMfYj8JcnKu0t0A4H5IRf0vRqSQiEBVIVz4j1PRlSa |  5 |       5 | rGdpE1xwL1DGytIrYLllVRDIJPr4OxM3GwN8xQG1hh0=

-- if we flip it around:

SELECT * FROM sessions
    RIGHT JOIN users ON users.id= sessions.user_id;

-- gets back all of the users even if there is no session

 id | user_id |                  token_hash                  | id |         email         |                        password_hash
----+---------+----------------------------------------------+----+-----------------------+--------------------------------------------------------------
  1 |       1 | 11irIAh40anc8lBeadXW4u7NS03mESI8Rcb3JWGolVM= |  1 | paul@santiago1.com    | $2a$10$Bx7LIRtWfh5X4hWW5hvm6.Wf0MDLg27ZQSot2uP.JWcZS4n7VeI5S
  2 |       2 | bYo1In4M57JqmtnJ2r8pq0N6gOaz8VfwEn3r9Njr6lI= |  2 | em@velasquez1.com     | $2a$10$AjxUaiyL3nsvYwHKbOJ.Qe5ZkzL356fmVfWtastb9gVWLMJs1AvTa
  3 |       3 | S4QdT2F7G6RKX1YHvCwOLeGcd3_cJyEiyI9J_nauKDI= |  3 | john@santiago1.com    | $2a$10$RSJoPc9KMhDy6v0Vl7HJIeKfxpFcylrAiO4sfjjZpHJx1SjLDL/dS
  5 |       5 | rGdpE1xwL1DGytIrYLllVRDIJPr4OxM3GwN8xQG1hh0= |  5 | thea@gonazalez1.com   | $2a$10$H0ySs8tx71tIMfYj8JcnKu0t0A4H5IRf0vRqSQiEBVIVz4j1PRlSa
    |         |                                              |  6 | jim@dundermifflin.com | $2a$10$BCBfSuYzZaRXf79IFlsfTOtAugBnJD/Q/3Lf/ZSHkXIQiZbJrwlU6
    |         |                                              |  7 | pam@dundermifflin.com | $2a$10$cy2xcLuQT7UmCEQXw3ff0uG0Ilt4I4gIzRM/QGIDGsO4tmzUI5xUi

-- FULL OUTER JOIN

SELECT * FROM users
     FULL OUTER JOIN sessions ON users.id = sessions.user_id;

-- return everything from both tables even if they don't have the same realtionhips


 id |         email         |                        password_hash                         | id | user_id |                  token_hash
----+-----------------------+--------------------------------------------------------------+----+---------+----------------------------------------------
  1 | paul@santiago1.com    | $2a$10$Bx7LIRtWfh5X4hWW5hvm6.Wf0MDLg27ZQSot2uP.JWcZS4n7VeI5S |  1 |       1 | 11irIAh40anc8lBeadXW4u7NS03mESI8Rcb3JWGolVM=
  2 | em@velasquez1.com     | $2a$10$AjxUaiyL3nsvYwHKbOJ.Qe5ZkzL356fmVfWtastb9gVWLMJs1AvTa |  2 |       2 | bYo1In4M57JqmtnJ2r8pq0N6gOaz8VfwEn3r9Njr6lI=
  3 | john@santiago1.com    | $2a$10$RSJoPc9KMhDy6v0Vl7HJIeKfxpFcylrAiO4sfjjZpHJx1SjLDL/dS |  3 |       3 | S4QdT2F7G6RKX1YHvCwOLeGcd3_cJyEiyI9J_nauKDI=
  5 | thea@gonazalez1.com   | $2a$10$H0ySs8tx71tIMfYj8JcnKu0t0A4H5IRf0vRqSQiEBVIVz4j1PRlSa |  5 |       5 | rGdpE1xwL1DGytIrYLllVRDIJPr4OxM3GwN8xQG1hh0=
  6 | jim@dundermifflin.com | $2a$10$BCBfSuYzZaRXf79IFlsfTOtAugBnJD/Q/3Lf/ZSHkXIQiZbJrwlU6 |    |         |
  7 | pam@dundermifflin.com | $2a$10$cy2xcLuQT7UmCEQXw3ff0uG0Ilt4I4gIzRM/QGIDGsO4tmzUI5xUi |    |         |