CREATE TABLE user_statement
(
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(255) NOT NULL UNIQUE,
    salary INT,
    loan_amount INT
);