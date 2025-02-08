# Golang API Product Management

This is a RESTful API built using Golang with the Gin framework, GORM as the ORM for MySQL.

### Features
* Create, Update, and Get Product
* Create, and Get a Product Transaction
* Create, Update, and Get Inventory
* One to Many relations for Product ID
* GORM-based relational database modeling

# Technologies
![Golang](https://img.shields.io/badge/golang-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)  ![REST API](https://img.shields.io/badge/restapi-%23000000.svg?style=for-the-badge&logo=swagger&logoColor=white)   ![MySQL](https://img.shields.io/badge/mysql-%234479A1.svg?style=for-the-badge&logo=mysql&logoColor=white)  

Uses golang as main frameworks for build an API, with RESTful API for communication with mySQL database

* Golang (Gin-Gonic framework) 
* MySQL (Database)
* GORM (ORM for database operations)

# API Endpoints Documentation

| Method     | API Endpoint               | Description                                      | Table             |
|------------|----------------------------|--------------------------------------------------|-------------------|
| **POST**   | `/products`                | Create a new product                             | Products Table     |
| **GET**    | `/products/:id`            | Get details of a specific product by ID          | Products Table     |
| **PUT**    | `/products/:id`            | Update an existing product by ID                 | Products Table     |
| **DELETE** | `/products/:id`            | Delete a product by ID                           | Products Table     |
| **GET**    | `/inventories/:id`         | Get inventory details by ID                      | Inventory Table    |
| **PUT**    | `/inventories/:id`         | Update inventory details by ID                   | Inventory Table    |
| **POST**   | `/orders`                  | Create a new order                               | Orders Table       |
| **GET**    | `/orders/:id`              | Get details of a specific order by ID            | Orders Table       |
