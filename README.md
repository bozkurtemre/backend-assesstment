# Backend Assessment

This project is designed to handle high-traffic scenarios with two services. 

The project is complete and ready for use, and the instructions below outline how to run and test it.

## Project Structure

- **src/frontend**: Frontend application code.
- **src/worker**: Background workers and related code.
- **src/docker-compose.yml**: Docker Compose configuration file, defines the project's services.
- **src/resources**: Resource files, such as SQL scripts.
- **src/frontend/scripts/e2e-testing.sh**: Bash script to run end-to-end (E2E) tests.

## Getting Started

### Requirements

- Docker and Docker Compose
- [Go](https://golang.org/)

### Setup and Running

1. **Clone the Project**: Clone the project from GitHub or another repository source.
   ```bash
   git clone https://github.com/bozkurtemre/backend-assesstment.git 
   ```

2. **Navigate to the Project Directory**: Move to the project directory.
   ```bash
   cd backend-assessment 
   ```

3. **Start the Services**: Start the services using Docker Compose.
   ```bash
   make run
   ```
   
Now you can access the frontend application at `http://localhost:8080`.

### Testing

1. End-to-End (E2E) Testing: Run the E2E testing script.
   ```bash
   make e2e-test
   ```
   
2. Unit Testing: Run the unit tests.
   ```bash
   make run-test
   ```

### API Collection
Postman collection is included in the project for testing the API endpoints.

## License

Licensed under [Apache 2.0](LICENSE).
