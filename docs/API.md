# API Documentation

This document describes the HTTP API implemented in the current codebase.

Base URL for local development:

```text
http://localhost:8080
```

Base API path:

```text
/api/v1
```

## Common Behavior

- Request and response bodies are JSON where applicable.
- Server errors return HTTP `500` with a simple JSON message.
- Invalid food creation payloads return HTTP `400`.
- Routes are registered with Gin in `api/routes/`.

## GET /customers/getall

Lists customers ordered by `customerid` ascending.

```http
GET /api/v1/customers/getall?page=1&pageSize=10
```

Query parameters:

| Name | Type | Default | Description |
| --- | --- | ---: | --- |
| `page` | integer | `1` | Page index. Values lower than 1 fall back to 1. |
| `pageSize` | integer | `10` | Number of rows per page. Values lower than 1 fall back to 10. |

Successful response:

```json
{
  "customers": [
    {
      "customer_id": 1,
      "first_name": "Amir",
      "last_name": "Akbari",
      "email": "amir@email.com",
      "phone_number": "09121111111",
      "customer_type": "registered",
      "join_date": "2026-06-12T10:00:00Z"
    }
  ],
  "pagination": {
    "pageIndex": 1,
    "pageSize": 10,
    "total": 10,
    "totalPage": 1
  }
}
```

Implementation files:

- `api/routes/customer_routes.go`
- `api/handlers/customers/queries_handler.go`
- `data/sql-scripts/queries/cutomers/cutomerQueries.go`
- `models/customerModels.go`

## GET /foods/getfoods

Lists food items with their category names.

```http
GET /api/v1/foods/getfoods
```

Optional category filter:

```http
GET /api/v1/foods/getfoods?foodCategory=Pizza
```

Query parameters:

| Name | Type | Default | Description |
| --- | --- | --- | --- |
| `foodCategory` | string | empty | Filters by `food_category.category_name`. |

Successful response:

```json
{
  "foods": {
    "Data": [
      {
        "FoodName": "Pepperoni Pizza",
        "FoodCategory": "Pizza",
        "Description": "Spicy Italian pizza with special sauce",
        "Calories": 1200,
        "ImagePath": "/images/pepperoni.jpg",
        "UnitPrice": 220000,
        "Inventory": 30
      }
    ]
  }
}
```

Note: `FoodItem` and `FoodItemsListResponse` currently do not define JSON tags,
so the response uses exported Go field names.

Implementation files:

- `api/routes/food_routes.go`
- `api/handlers/foods/foods_handler.go`
- `data/sql-scripts/queries/foods/foodQueries.go`
- `models/foodModels.go`

## POST /foods/createfood

Creates a new food item and returns the generated `food_id`.

```http
POST /api/v1/foods/createfood
Content-Type: application/json
```

Request body:

```json
{
  "foodName": "Mushroom Burger",
  "description": "Beef burger with mushrooms and cheese",
  "calories": 780,
  "imagePath": "/images/mushroom-burger.jpg",
  "unitPrice": 210000,
  "inventory": 25,
  "foodCategory": 2
}
```

Request fields:

| Name | Type | Required | Description |
| --- | --- | --- | --- |
| `foodName` | string | yes | Food item name. |
| `description` | string | no | Food description. |
| `calories` | integer | no | Calorie count. Database requires non-negative values. |
| `imagePath` | string | no | Image path or URL. |
| `unitPrice` | number | yes | Price. Database requires non-negative values. |
| `inventory` | integer | yes | Available quantity. Database requires non-negative values. |
| `foodCategory` | integer | yes | Foreign key to `food_category.category_id`. |

Successful response:

```json
{
  "FoodId": 11,
  "Message": "Created food"
}
```

Note: `CreateFoodItemResponse` currently does not define JSON tags, so the
response uses exported Go field names.

Possible errors:

| Status | Cause |
| ---: | --- |
| `400` | Required JSON fields are missing or the JSON body is invalid. |
| `500` | Database insert failed, including invalid category id or constraint failure. |

Implementation files:

- `api/routes/food_routes.go`
- `api/handlers/foods/mutation_handlers.go`
- `data/sql-scripts/mutations/foods/foodMutations.go`
- `models/foodModels.go`

## Route Map

| Method | Path | Handler |
| --- | --- | --- |
| `GET` | `/api/v1/customers/getall` | `customers.GetAllCustomers` |
| `GET` | `/api/v1/foods/getfoods` | `foods.GetAllFoods` |
| `POST` | `/api/v1/foods/createfood` | `foods.HandleCreateFood` |
