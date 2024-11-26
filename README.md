# Transaction Summary Application

This application processes financial transactions and sends transaction summaries via email. It uses **MailHog** for email testing in a development environment.

## Requirements

To run this application, ensure the following software is installed on your machine:

1. **Docker**
    - Install Docker from [https://www.docker.com/get-started](https://www.docker.com/get-started).

2. **Docker Compose**
    - Docker Compose typically comes with Docker Desktop. Verify its installation with:
      ```bash
      docker-compose --version
      ```

3. **Git**
    - For cloning the repository. Install it from [https://git-scm.com/](https://git-scm.com/). Verify its installation with:
      ```bash
      git --version
      ```


## How to Run

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/transaction-summary.git
   cd transaction-summary

2. **Set up and run the application:**
    ```bash
   # Build and start the containers:
    docker-compose up --build
   # Stop the application when done:
    docker-compose down
3. **Access MailHog: Open your browser and navigate to:**
    ```bash
    http://localhost:8025
    ```

Here, you can view all emails sent by the application.

## Endpoints

### The application has the following endpoints:

1. **Send an email**
   ```bash
   http://localhost:8080/send_email
   ```
   This endpoint will send a stylized email with a transaction summary based on the contents of transactionvs.csv.
   Navigate to the MailHog interface to view the email.


2. **Create Account**
   ```bash
   http://localhost:8080/account?account_number=123456789
   ```
   This endpoint creates a new account in the database. It requires an account number as a query parameter. The Method must be a POST.


3. **Get Account Balance**
   ```bash
   http://localhost:8080/account?account_number=123456789
   ```
   This endpoint gets the account balance for a given account_number in the database. It requires an account number as a query parameter. The Method must be a GET.


4. **Save Transaction**
   ```bash
   http://localhost:8080/transactions
   ```
   This endpoint saves a transaction for a specific account_number in the database. It requires an account number and an amount in a JSON body. The Method must be a POST.
   The body should look like this:
   ```json
   {
       "account_number": "123456789",
       "amount": 100
   }
   ```

5. **Get Transactions**
   ```bash
    http://localhost:8080/transactions?account_number=123456789
   ```
   This endpoint gets all transactions for a specific account_number in the database. It requires an account number as a query parameter.
   The Method must be a GET.

To check existent accounts and transactions in the database, check the init.sql file.





