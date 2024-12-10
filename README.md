# Banking System

This project provides APIs for a banking system backend with the following functionalities:

## Features

1. **Bank Account Management**  
   - Create and manage bank accounts.  
   - Bank accounts consist of:
     - Owner’s Name  
     - Balance  
     - Currency  

2. **Transaction Records**  
   - Record every balance change for each account.  
   - Create an account entry record for every deposit, withdrawal, or transfer.

3. **Money Transfers**  
   - Perform money transfers between two accounts.  
   - Ensure transactional integrity: updates to both accounts occur together, or none of the updates are applied.

---

## Tools and Dependencies

### Core Tools

- **[Docker](https://docs.docker.com/engine/)**  
  For running postgres image.

- **[TablePlus](https://tableplus.com/)**  
  A GUI tool for interacting with the database.

- **[Golang](https://golang.org/)**  
  The programming language used for backend development.

### Database Migrations

- **[Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)**  
  Used for database schema migrations.

  **Installation:**
  ```bash
  curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
  echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
  apt-get update
  apt-get install -y migrate
  ```

### SQL Code Generation

- **[Sqlc](https://github.com/kyleconroy/sqlc#installation)**  
  Used to generate type-safe Go code from SQL queries.

  **Installation:**
  ```bash
  sudo snap install sqlc
  ```

### Mock Testing

- **[Gomock](https://github.com/golang/mock)**  
  Used for generating mock objects for unit testing.

  **Installation:**
  ```bash
  go install github.com/golang/mock/mockgen
  ```

---

## Setup Instructions

1. **Install Dependencies**  
   Ensure all tools mentioned above are installed on your system.

2. **Run Migrations**  
   Use `migrate` to apply database migrations. For example:
   ```bash
   migrate -path db/migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up
   ```

3. **Run Tests**  
   To run tests:
   ```bash
   make test
   ```

4. **Build and Run the Application**  
   To run the application and start the server:
   ```bash
   make server
   ```

---

## Contributing

- Follow the Git workflow for feature development.
- Ensure your code is covered by unit tests.
- Run all tests before submitting a pull request.

---

This README serves as a guide for setting up and understanding the Banking System project.
