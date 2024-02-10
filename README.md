## Directory Descriptions

- **api**: Contains modules related to API endpoints.

  - **products**: Module for managing product-related endpoints.

    - `productRoutes.go`: Routes for product endpoints.

  - **users**: Module for managing user-related endpoints.

    - `userRoutes.go`: Routes for user endpoints.

- **config**: Configuration files for the project.

  - `config.go`: Main configuration file.

- **infrastructure**: Infrastructure-related files.

  - `.env`: Environment variables file.

  - `docker-compose.yaml`: Docker Compose configuration file.

  - `Dockerfile`: Dockerfile for building the application container.

- **middleware**: Middleware components for request processing.

  - **authentication**: Middleware for authentication.

    - `authentication.go`: Authentication middleware logic.

    - `refreshtoken.go`: Logic for refreshing tokens.

    - `revoketoken.go`: Logic for revoking tokens.

  - **authorization**: Middleware for authorization.

    - `AuthorizeRequest.go`: Logic for authorizing requests.

  - **hashpassword**: Middleware for hashing passwords.

    - `hashpassword.go`: Logic for hashing passwords.

  - **logger**: Middleware for logging.

    - `logger.go`: Logging middleware logic.

  - **validations**: Middleware for request validations.

    - `validations.go`: Logic for request validations.

- **models**: Contains data models for the application.

  - `authmodel.go`: Authentication-related data models.

  - `UserModel.go`: User-related data models.

- **services**: Business logic services.

  - `userservices.go`: Business logic related to user management.

- **.gitignore**: Git ignore file to specify intentionally untracked files.
- **go.mod**: Go module file specifying the module's module path and its dependencies.
- **go.sum**: Go dependency file containing the hashed versions of dependencies.
- **main.go**: Entry point of the application.

## Steps to Run the Project

1. Navigate to the `infrastructure` directory:
    ```
    cd infrastructure
    ```

2. Build the Docker images using Docker Compose:
    ```
    docker-compose build
    ```

3. Once the images are built, start the Docker containers:
    ```
    docker-compose up
    ```

This will start the project and make it accessible according to the configuration specified in the Docker Compose file.

## Rebuilding the API

If you make changes to the API code and need to rebuild it, follow these steps:

1. Make your changes to the API code.

2. Stop the running Docker containers:
    ```
    docker-compose down
    ```

3. Rebuild the Docker images:
    ```
    docker-compose build
    ```

4. Start the Docker containers again:
    ```
    docker-compose up
    ```

This will rebuild the API with your changes and start the project again.

## Curl Commands for API Endpoints

1. SignUp Endpoint:
    ```
   curl -i -X POST http://localhost:8080/signup -d "{\"Email\":\"ashokmamilla899@gmail.com\", \"Password\":\"ashok53323\"}" -H "Content-Type: application/json"
    ```
2. SignIn Endpoint:
    ```
    curl -i -X POST http://localhost:8080/signin -d "{\"Email\":\"ashokmamilla899@gmail.com\", \"Password\":\"ashok53323\"}" -H "Content-Type: application/json"
   ```
3. Authorize Request:
   ```
     curl -i -X GET http://localhost:8080/protected \-H "Authorization: Bearer <access_token>"
   ```
   Note: Make the command as single line otherwise curl throws an error.

4. Refresh Token Endpoint:
  ``` 
   curl -i -X POST http://localhost:8080/refresh-token -H "Authorization: Bearer <refresh_token>" 
   ```
5. Revoke Token:
   ```
    curl -i -X POST http://localhost:8080/revoke-token -H "Authorization: Bearer <access_token>"
   ```   
