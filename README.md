# MyCandy's Orders Microservice

## Description

This is the Orders Microservice for the MyCandy application. It is a RESTful API that allows users to create, read,
update, and delete orders.
It also allows users to view all orders, view orders by user, and view orders by status.

## Environment Variables

The following environment variables must be set in order for the application to run:

| Variable Name            | Description                               |
|--------------------------|-------------------------------------------|
| PORT                     | The port to run the application on.       |
| DATABASE_URL             | The URL of the database to connect to.    |
| DATABASE_NAME            | The name of the database to connect to.   |
| PRODUCT_SERVICE_URL      | The URL of the Product Microservice.      |
| NOTIFICATION_SERVICE_URL | The URL of the Notification Microservice. |
| AUTH_SERVICE_URL         | The URL of the Auth Microservice.         |

**Example file**

```
    PORT=8080
    DATABASE_URL=mongodb://localhost:27017
    DATABASE_NAME=orders
    PRODUCT_SERVICE_URL=http://localhost:8081
    NOTIFICATION_SERVICE_URL=http://localhost:8082
    AUTH_SERVICE_URL=http://localhost:8083
```

## Running the Application

To run the application, you must have Docker installed on your machine. Once you have Docker installed, you can run the
following command to start the application:

```bash
docker-compose up
```

## API Documentation

After running the application the API documentation can be found at the following
link: http://localhost:8080/swagger/index.html

