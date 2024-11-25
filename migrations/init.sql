CREATE DATABASE IF NOT EXISTS transactions_db;

USE transactions_db;

CREATE TABLE transactions (
                              id INT AUTO_INCREMENT PRIMARY KEY,
                              transaction_date DATE NOT NULL,
                              amount DECIMAL(10, 2) NOT NULL,
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE accounts (
                          id INT AUTO_INCREMENT PRIMARY KEY,
                          total_balance DECIMAL(10, 2) NOT NULL,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
