CREATE TYPE customer_type_enum AS ENUM ('registered','guest','vip','corporate');

CREATE TYPE order_type_enum AS ENUM ('dine_in','takeaway','delivery');

CREATE TYPE discount_type_enum AS ENUM ('percentage','fixed_amount','free_delivery');

CREATE TYPE discount_status_enum AS ENUM ('active','expired','disabled');

CREATE TYPE employee_role_enum AS ENUM ('admin','manager','cashier','waiter','chef');

CREATE TYPE work_shift_enum AS ENUM ('morning','afternoon','night');

CREATE TABLE Customer(
    CustomerID BIGSERIAL PRIMARY KEY,
    First_name VARCHAR(100) NOT NULL,
    Last_name VARCHAR(100) NOT NULL,
    Email VARCHAR(255) UNIQUE,
    Phone_number VARCHAR(20) NOT NULL,
    Join_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Credit NUMERIC(12,2) DEFAULT 0 CHECK (Credit >= 0),
    Customer_type customer_type_enum NOT NULL
);

CREATE TABLE Customer_Address(
    Customer_ID BIGINT NOT NULL,
    Address_Label VARCHAR(50) NOT NULL,
    City VARCHAR(100) NOT NULL,
    Postal_code VARCHAR(20),
    Street VARCHAR(255) NOT NULL,

    PRIMARY KEY (Customer_ID, Address_Label),

    FOREIGN KEY (Customer_ID)
        REFERENCES Customer(CustomerID)
        ON DELETE CASCADE
);

CREATE TABLE Transaction_Status(
    Status_ID SMALLSERIAL PRIMARY KEY,
    Name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE OrderStatus(
    Status_ID SMALLSERIAL PRIMARY KEY,
    Name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE TableStatus(
    Status_ID SMALLSERIAL PRIMARY KEY,
    Name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE Tables(
    Table_ID SERIAL PRIMARY KEY,
    Table_Number INT UNIQUE NOT NULL,
    Capacity SMALLINT NOT NULL CHECK (Capacity > 0),
    Location VARCHAR(255),
    TableStatus SMALLINT REFERENCES TableStatus(Status_ID)
);

CREATE TABLE Discount(
    Discount_ID BIGSERIAL PRIMARY KEY,
    DiscountType discount_type_enum NOT NULL,
    DiscountValue NUMERIC(10,2) CHECK (DiscountValue >= 0),
    StartDate TIMESTAMP NOT NULL,
    ExpiryDate TIMESTAMP NOT NULL,
    MinOrderAmount NUMERIC(10,2) DEFAULT 0,
    MaxUsageCount INT CHECK (MaxUsageCount >= 0),
    Status discount_status_enum NOT NULL,
    CHECK (ExpiryDate > StartDate)
);

CREATE TABLE Invoice(
    Invoice_ID BIGSERIAL PRIMARY KEY,
    TaxAndServiceAmount NUMERIC(10,2) DEFAULT 0 CHECK (TaxAndServiceAmount >= 0),
    TotalItemsAmount NUMERIC(10,2) DEFAULT 0 CHECK (TotalItemsAmount >= 0),
    DiscountAmount NUMERIC(10,2) DEFAULT 0 CHECK (DiscountAmount >= 0),
    FinalPayableAmount NUMERIC(10,2) NOT NULL CHECK (FinalPayableAmount >= 0),
    IssueDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Orders(
    Order_ID BIGSERIAL PRIMARY KEY,
    PhoneNumber VARCHAR(20),
    OrderDateTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    OrderType order_type_enum NOT NULL,

    TotalAmountBeforeDiscount NUMERIC(12,2) DEFAULT 0 CHECK (TotalAmountBeforeDiscount >= 0),
    TotalAmountAfterDiscount NUMERIC(12,2) DEFAULT 0 CHECK (TotalAmountAfterDiscount >= 0),

    OrderStatus SMALLINT REFERENCES OrderStatus(Status_ID),
    CustomerID BIGINT REFERENCES Customer(CustomerID) ON DELETE SET NULL,
    InvoiceID BIGINT UNIQUE REFERENCES Invoice(Invoice_ID),
    DiscountID BIGINT REFERENCES Discount(Discount_ID),
    TableID INT REFERENCES Tables(Table_ID),

    Address TEXT,

    CHECK (
        (OrderType != 'delivery') OR (Address IS NOT NULL)
    )
);

CREATE TABLE Payment_transaction(
    Transaction_ID BIGSERIAL PRIMARY KEY,
    Amount NUMERIC(12,2) NOT NULL CHECK (Amount >= 0),
    Tracking_code VARCHAR(255) UNIQUE NOT NULL,
    Payment_Datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Order_ID BIGINT REFERENCES Orders(Order_ID) ON DELETE CASCADE,
    Payment_Status SMALLINT REFERENCES Transaction_Status(Status_ID)
);

CREATE TABLE Food_Category(
    Category_ID SMALLSERIAL PRIMARY KEY,
    Category_name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE Food_item(
    Food_ID BIGSERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Description TEXT,
    Calories INT CHECK (Calories >= 0),
    ImagePath TEXT,
    Inventory INT NOT NULL CHECK (Inventory >= 0),
    UnitPrice NUMERIC(10,2) NOT NULL CHECK (UnitPrice >= 0),
    FoodCategory SMALLINT REFERENCES Food_Category(Category_ID)
);

CREATE TABLE Order_Food_Contains(
    Order_ID BIGINT NOT NULL,
    Food_ID BIGINT NOT NULL,

    UnitPriceAtOrder NUMERIC(10,2) NOT NULL
        CHECK (UnitPriceAtOrder >= 0),

    Quantity INT NOT NULL
        CHECK (Quantity > 0),

    PRIMARY KEY (Order_ID, Food_ID),

    FOREIGN KEY (Order_ID)
        REFERENCES Orders(Order_ID)
        ON DELETE CASCADE,

    FOREIGN KEY (Food_ID)
        REFERENCES Food_item(Food_ID)
);

CREATE TABLE Employee(
    Employee_ID BIGSERIAL PRIMARY KEY,

    SSN VARCHAR(20) UNIQUE NOT NULL,

    First_name VARCHAR(100) NOT NULL,
    Last_name VARCHAR(100) NOT NULL,

    Username VARCHAR(100) UNIQUE NOT NULL,
    Password TEXT NOT NULL,

    WorkShift work_shift_enum,
    Role employee_role_enum,

    Super_Employee_ID BIGINT,

    FOREIGN KEY (Super_Employee_ID)
        REFERENCES Employee(Employee_ID)
);

CREATE INDEX idx_order_customer ON Orders(CustomerID);

CREATE INDEX idx_order_date ON Orders(OrderDateTime);

CREATE INDEX idx_payment_order ON Payment_transaction(Order_ID);

CREATE INDEX idx_food_category ON Food_item(FoodCategory);

CREATE INDEX idx_discount_status ON Discount(Status);

