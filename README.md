# E-commerce API

## Overview
This project is a RESTful API for an e-commerce application built using Go (version 1.22.2). The API provides functionalities for user management, product management, and order management, adhering to best practices in software design and architecture.

## Features
- **User Management**: Register and authenticate users using JWT for secure session management.
- **Product Management**: Admin users can create, read, update, and delete products.
- **Order Management**: Authenticated users can place orders, view their orders, and cancel pending orders.

## Technologies Used
- **Programming Language**: Go (1.22.2)
- **Framework**: Gin (or Echo, depending on implementation)
- **Database**: PostgreSQL or MySQL
- **Documentation**: Swagger for API documentation

## Project Structure
```
ecommerce-api
├── cmd
│   └── main.go
├── internal
│   ├── auth
│   │   ├── jwt.go
│   │   └── middleware.go
│   ├── config
│   │   └── config.go
│   ├── controllers
│   │   ├── order_controller.go
│   │   ├── product_controller.go
│   │   └── user_controller.go
│   ├── models
│   │   ├── order.go
│   │   ├── product.go
│   │   └── user.go
│   ├── repository
│   │   ├── order_repository.go
│   │   ├── product_repository.go
│   │   └── user_repository.go
│   ├── routes
│   │   └── routes.go
│   ├── services
│   │   ├── order_service.go
│   │   ├── product_service.go
│   │   └── user_service.go
│   └── utils
│       └── utils.go
├── docs
│   └── swagger.yaml
├── go.mod
├── go.sum
└── README.md
```