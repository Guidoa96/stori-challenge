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






