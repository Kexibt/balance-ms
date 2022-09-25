CREATE TABLE balances (
    userID VARCHAR(100) PRIMARY KEY NOT NULL,
    balance FLOAT NOT NULL,
    commnt TEXT,
    date TIMESTAMP
);

SELECT * FROM balances;

DELETE FROM balances;

INSERT INTO balances ("userid", "balance", "commnt", "date") 
VALUES ('user2_test', 200, ' ', now()),
('user3_test', 300, ' ', now()),
('user4_test', 600, ' ', now()),
('user5_test', 100, ' ', now()),
('user6_test', 50, ' ', now());

UPDATE balances SET balance = 0 WHERE userid = 'user4_test';