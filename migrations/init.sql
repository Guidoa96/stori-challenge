CREATE DATABASE IF NOT EXISTS transactions_db;

USE transactions_db;


-- Accounts Table
CREATE TABLE IF NOT EXISTS accounts (
        id INT AUTO_INCREMENT PRIMARY KEY,
        account_number VARCHAR(20) NOT NULL UNIQUE,
        total_balance DECIMAL(10, 2) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Transactions Table
CREATE TABLE IF NOT EXISTS transactions (
      id INT AUTO_INCREMENT PRIMARY KEY,
      amount DECIMAL(10, 2) NOT NULL,
      account_id INT NOT NULL,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Insert initial accounts
INSERT INTO accounts (account_number, total_balance) VALUES
     ('123456789', 1000.00),
     ('987654321', 500.50),
     ('555555555', 750.75);

-- Insert initial transactions
INSERT INTO transactions (amount, account_id) VALUES
  (200.00, 1),   -- Deposit to account 123456789
  (-50.00, 1),   -- Withdrawal from account 123456789
  (300.00, 2),   -- Deposit to account 987654321
  (-100.50, 2),  -- Withdrawal from account 987654321
  (500.00, 3),   -- Deposit to account 555555555
  (-250.75, 3);  -- Withdrawal from account 555555555