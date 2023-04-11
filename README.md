# Assignment 10 - REST API Product with Login and Register

This project is a RESTful API for managing products with user authentication and authorization features. The API is built using Golang and provides the following functionality:

## Features

- User authentication and registration
- Create, read, update, and delete (CRUD) operations for products
- Middleware for authentication
- Multi-level user authorization
- Authorization to access a product by ID

## Endpoints

1. Register - `POST /register`
2. Login - `POST /login`
3. Create Product - `POST /products`
4. Get All Products - `GET /products`
5. Get Product by ID - `GET /products/:id`
6. Update Product - `PUT /products/:id`
7. Delete Product - `DELETE /products/:id`

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
