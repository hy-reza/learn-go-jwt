# Assignment 10 FGA Go Hactiv8 - REST API Product with Login and Register

This project is a RESTful API for managing products with user authentication and authorization features. The API is built using Golang and provides the following functionality:

## Features

- User authentication and registration
- Create, read, update, and delete (CRUD) operations for products
- Middleware for authentication
- Multi-level user authorization
- Authorization to access a product by ID

## Endpoints

| Method | Endpoint                | Middleware                      | Description                |
|--------|-------------------------|---------------------------------|----------------------------|
| POST   | /users/register         |                                 | Register a new user        |
| POST   | /users/login            |                                 | Login a user               |
| GET    | /products/              | Authentication                  | Get all products           |
| GET    | /products/:productId    | Authentication, ProductAuthorization | Get a product by ID      |
| POST   | /products/              | Authentication                  | Create a new product       |
| PUT    | /products/:productId    | Authentication, ProductAuthorization | Update a product by ID  |
| DELETE | /products/:productId    | Authentication, ProductAuthorization | Delete a product by ID  |


## Middleware

- Authentication: Validates the JSON Web Token (JWT) from the `Authorization` header.
- Multi-level user authorization: Checks the user's role (e.g., admin, user) and grants access accordingly.
- Authorization access product by ID: Ensures that the user has permission to access the requested product.

## Postman Documentation

Please refer to the [Postman documentation](https://documenter.getpostman.com/view/15041975/2s93Xu36WW) for detailed API usage and example requests.

## Getting Started

To run the project locally, follow these steps:

1. Clone the repository.
2. Install the required dependencies.
3. Set up the database and environment variables.
4. Run the application using `go run main.go`.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
