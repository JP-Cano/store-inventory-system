## Store Inventory Management System

This is a README file for the Store Inventory Management System built using Go and Echo framework, PostgreSQL for
database management, and hosted on localhost.

### Getting Started

To start the server, run the following command:

```sh
go run main.go
```

### Endpoints

### Product

- **GET** `http://localhost:8080/products`
    - Retrieves a list of products.


- **GET** `http://localhost:8080/products/search?value=coca`
    - Searches for a product by name or category.

### Sales

- **GET** `http://localhost:8080/sales/search?date=2024-06-01`
    - Retrieves sales for a specific date.


- **GET** `http://localhost:8080/sales/report?start-date=2024-06-01&end-date=2024-06-03`
    - Generates a sales report within a specified date range.


- **POST** `http://localhost:8080/sales`
    - Creates a new sale.
    - Content-Type: application/json
  ```json
  {
    "product_id": "700c5a90-c308-4547-a1d0-952d995585e3",
    "customer_name": "Pedro",
    "quantity_sold": 3
  }

### About the Application

The Store Inventory Management System is a web application designed to manage product inventory and sales efficiently.
Built using Go and Echo framework, PostgreSQL for database management, and hosted in Google Cloud App Engine.

### About Echo Framework

Echo is a high-performance, minimalist Go web framework inspired by Sinatra. It's known for its simplicity and speed,
making it a popular choice for building web applications in Go. With its lightweight design and powerful features, Echo
provides developers with a convenient and efficient way to create robust web applications.
