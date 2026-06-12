# Database Documentation

The project uses PostgreSQL. Schema creation is stored in
`ResturnDB_Schema.sql`, and development seed data is stored in
`mockdata_migration.sql`.

## Development Connection

The development config points to:

```text
host=localhost
port=5433
user=postgres
password=12345678
dbname=restaurant_db
sslmode=disable
```

Connection string:

```bash
postgres://postgres:12345678@localhost:5433/restaurant_db?sslmode=disable
```

## Setup Commands

Create infrastructure:

```bash
docker compose -f docker/docker-compose.yaml up -d postgres
```

Apply schema:

```bash
psql "postgres://postgres:12345678@localhost:5433/restaurant_db?sslmode=disable" \
  -f ResturnDB_Schema.sql
```

Apply seed data:

```bash
psql "postgres://postgres:12345678@localhost:5433/restaurant_db?sslmode=disable" \
  -f mockdata_migration.sql
```

## Main Tables

### Customer Domain

- `customer`: customer profile data, credit, and customer type.
- `customer_address`: one or more labeled addresses per customer.

Customer type enum:

```text
registered, guest, vip, corporate
```

### Restaurant Tables

- `tablestatus`: lookup table for table state.
- `tables`: physical restaurant tables, capacity, location, and status.

### Menu Domain

- `food_category`: menu categories.
- `food_item`: menu items, price, inventory, image path, and category.

The currently implemented food endpoints read and write this domain.

### Ordering Domain

- `orders`: order header, order type, customer, invoice, discount, table, and
  delivery address.
- `order_food_contains`: order line items and quantity.
- `orderstatus`: lookup table for order state.

Order type enum:

```text
dine_in, takeaway, delivery
```

Delivery orders must include an address because of the schema check constraint.

### Payments And Invoices

- `invoice`: calculated money totals for an order.
- `payment_transaction`: payment amount, tracking code, date, order, and status.
- `transaction_status`: lookup table for payment state.

### Discounts

- `discount`: discount type, value, validity dates, usage limit, and status.

Discount type enum:

```text
percentage, fixed_amount, free_delivery
```

Discount status enum:

```text
active, expired, disabled
```

### Employees

- `employee`: staff identity, credentials, role, shift, and manager relation.

Employee role enum:

```text
admin, manager, cashier, waiter, chef
```

Work shift enum:

```text
morning, afternoon, night
```

## Indexes

The schema creates indexes for common relationships and filters:

- `idx_order_customer` on `orders(customerid)`
- `idx_order_date` on `orders(orderdatetime)`
- `idx_payment_order` on `payment_transaction(order_id)`
- `idx_food_category` on `food_item(foodcategory)`
- `idx_discount_status` on `discount(status)`

## Tables Used By Current API

Current handlers directly query:

| API area | Tables |
| --- | --- |
| Customers | `customer` |
| Foods | `food_item`, `food_category` |

The rest of the schema is available for future API expansion.

## Naming Notes

PostgreSQL folds unquoted identifiers to lowercase. The schema uses names such
as `CustomerID`, `FoodCategory`, and `UnitPrice`, but the SQL queries use their
lowercase forms such as `customerid`, `foodcategory`, and `unitprice`.
