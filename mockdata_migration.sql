INSERT INTO Transaction_Status (Status_ID, Name) VALUES
(1, 'Successful'),
(2, 'Failed'),
(3, 'Pending Payment'),
(4, 'Cancelled');

INSERT INTO OrderStatus (Status_ID, Name) VALUES
(1, 'Placed'),
(2, 'Preparing'),
(3, 'Ready for Pickup/Delivery'),
(4, 'Delivered'),
(5, 'Cancelled');

INSERT INTO TableStatus (Status_ID, Name) VALUES
(1, 'Available'),
(2, 'Occupied'),
(3, 'Reserved'),
(4, 'Cleaning');

INSERT INTO Food_Category (Category_ID, Category_name) VALUES
(1, 'Persian Food'),
(2, 'Fast Food'),
(3, 'Pizza'),
(4, 'Beverages'),
(5, 'Appetizers'),
(6, 'Salads'),
(7, 'Desserts'),
(8, 'Seafood'),
(9, 'Pasta'),
(10, 'Vegetarian');

INSERT INTO Employee
(Employee_ID, SSN, First_name, Last_name, Username, Password, WorkShift, Role, Super_Employee_ID)
VALUES
(1, '001', 'Ali', 'Rezaei', 'admin_ali', 'hashed_pass_1', 'morning', 'admin', NULL),
(2, '002', 'Mohammad', 'Ahmadi', 'manager_m', 'hashed_pass_2', 'morning', 'manager', 1),
(3, '003', 'Sara', 'Karami', 'cash_sara', 'hashed_pass_3', 'afternoon', 'cashier', 2),
(4, '004', 'Reza', 'Mohammadi', 'waiter_reza', 'hashed_pass_4', 'night', 'waiter', 2),
(5, '005', 'Mina', 'Jalali', 'chef_mina', 'hashed_pass_5', 'morning', 'chef', 2),
(6, '006', 'Hossein', 'Taheri', 'waiter_h', 'hashed_pass_6', 'afternoon', 'waiter', 2),
(7, '007', 'Omid', 'Nouri', 'chef_omid', 'hashed_pass_7', 'night', 'chef', 2),
(8, '008', 'Zahra', 'Hosseini', 'cash_zahra', 'hashed_pass_8', 'morning', 'cashier', 2),
(9, '009', 'Mehdi', 'Sadeghi', 'waiter_m', 'hashed_pass_9', 'morning', 'waiter', 2),
(10, '010', 'Nima', 'Rad', 'manager_n', 'hashed_pass_10', 'night', 'manager', 1);

INSERT INTO Customer
(CustomerID, First_name, Last_name, Email, Phone_number, Credit, Customer_type)
VALUES
(1, 'Amir', 'Akbari', 'amir@email.com', '09121111111', 50000.00, 'registered'),
(2, 'Maryam', 'Ghasemi', 'maryam@email.com', '09122222222', 0.00, 'guest'),
(3, 'Niloufar', 'Shafiei', 'nilou@email.com', '09123333333', 150000.00, 'vip'),
(4, 'Web Development', 'Co.', 'info@webdev.com', '02188888888', 1000000.00, 'corporate'),
(5, 'Sina', 'Parsa', 'sina@email.com', '09124444444', 0.00, 'registered'),
(6, 'Elnaz', 'Azimi', 'elnaz@email.com', '09125555555', 20000.00, 'registered'),
(7, 'Kamran', 'Najafi', 'kamran@email.com', '09126666666', 0.00, 'guest'),
(8, 'Aria Commerce', 'Inc.', 'contact@ariat.com', '02177777777', 500000.00, 'corporate'),
(9, 'Fatemeh', 'Moradi', 'fatemeh@email.com', '09128888888', 75000.00, 'vip'),
(10, 'Pouya', 'Jahandar', 'pouya@email.com', '09129999999', 0.00, 'registered');

INSERT INTO Customer_Address
(Customer_ID, Address_Label, City, Postal_code, Street)
VALUES
(1, 'home', 'Tehran', '1111122222', 'Valiasr St, 1st Alley, No. 10'),
(1, 'work', 'Tehran', '1111133333', 'Vanak, Molla Sadra St, No. 22'),
(3, 'home', 'Tehran', '3333344444', 'Saadat Abad, Kaj Square, No. 5'),
(4, 'office', 'Tehran', '5555566666', 'Jordan, Nahid St, Office Tower, Unit 2'),
(5, 'home', 'Karaj', '7777788888', 'Azimiyeh, Mehran Square, No. 12'),
(6, 'home', 'Tehran', '9999900000', 'Tehranpars, 1st Circle, 104th Alley, No. 3'),
(8, 'office', 'Tehran', '1234567890', 'Arjantin Square, Alvand St, No. 88'),
(9, 'home', 'Shiraz', '0987654321', 'Ghasrodasht, 20th Alley, No. 4'),
(10, 'home', 'Isfahan', '1357924680', 'Chaharbagh Abbasi, Sepahan Alley, No. 1'),
(3, 'work', 'Tehran', '3333355555', 'Niavaran, Ammar St, No. 7');

INSERT INTO Tables
(Table_ID, Table_Number, Capacity, Location, TableStatus)
VALUES
(1, 1, 2, 'Window Side', 1),
(2, 2, 4, 'Main Hall', 2),
(3, 3, 4, 'Main Hall', 1),
(4, 4, 6, 'VIP', 3),
(5, 5, 2, 'Outdoor (Terrace)', 1),
(6, 6, 8, 'VIP', 1),
(7, 7, 4, 'Window Side', 2),
(8, 8, 2, 'Outdoor (Terrace)', 4),
(9, 9, 10, 'Meeting Room', 3),
(10, 10, 4, 'Main Hall', 1);

INSERT INTO Discount
(Discount_ID, DiscountType, DiscountValue, StartDate, ExpiryDate, MinOrderAmount, MaxUsageCount, Status)
VALUES
(1, 'percentage', 10.00, '2023-01-01', '2025-12-31', 100000.00, 1000, 'active'),
(2, 'fixed_amount', 50000.00, '2024-03-01', '2024-04-01', 200000.00, 500, 'expired'),
(3, 'free_delivery', 0.00, '2024-01-01', '2025-01-01', 150000.00, 200, 'active'),
(4, 'percentage', 20.00, '2024-12-20', '2024-12-22', 0.00, 100, 'expired'),
(5, 'percentage', 15.00, '2024-06-01', '2025-06-01', 300000.00, 50, 'active'),
(6, 'fixed_amount', 25000.00, '2024-01-01', '2025-12-31', 100000.00, 5000, 'active');

INSERT INTO Food_item
(Food_ID, Name, Description, Calories, ImagePath, Inventory, UnitPrice, FoodCategory)
VALUES
(1, 'Chelo Kebab Koobideh', 'Two skewers of minced lamb kebab with Persian rice', 850, '/images/koobideh.jpg', 50, 250000.00, 1),
(2, 'Pepperoni Pizza', 'Spicy Italian pizza with special sauce', 1200, '/images/pepperoni.jpg', 30, 220000.00, 3),
(3, 'Alfredo Pasta', 'Penne with Alfredo sauce and grilled chicken', 950, '/images/pasta.jpg', 40, 180000.00, 9),
(4, 'Caesar Salad', 'Lettuce, grilled chicken, Parmesan cheese, and Caesar dressing', 450, '/images/caesar.jpg', 60, 120000.00, 6),
(5, 'Coca-Cola Can', '330ml carbonated soft drink', 140, '/images/coca.jpg', 200, 25000.00, 4),
(6, 'French Fries', 'Crispy potatoes with ketchup', 350, '/images/fries.jpg', 100, 75000.00, 5),
(7, 'Chelo Joojeh Kebab', 'One skewer of saffron chicken kebab with rice', 750, '/images/joojeh.jpg', 45, 210000.00, 1),
(8, 'Special Burger', '100% pure meat burger with Gouda cheese', 800, '/images/burger.jpg', 40, 190000.00, 2),
(9, 'Local Doogh (Glass)', 'Local mint yogurt drink', 80, '/images/doogh.jpg', 150, 20000.00, 4),
(10, 'Chocolate Cake', 'One slice of moist chocolate cake', 400, '/images/cake.jpg', 20, 80000.00, 7);

INSERT INTO Invoice
(Invoice_ID, TaxAndServiceAmount, TotalItemsAmount, DiscountAmount, FinalPayableAmount)
VALUES
(1, 27500.00, 275000.00, 0.00, 302500.00),
(2, 41000.00, 410000.00, 41000.00, 410000.00),
(3, 14500.00, 145000.00, 0.00, 159500.00),
(4, 52000.00, 520000.00, 50000.00, 522000.00),
(5, 12000.00, 120000.00, 0.00, 132000.00),
(6, 64000.00, 640000.00, 64000.00, 640000.00),
(7, 21000.00, 210000.00, 0.00, 231000.00),
(8, 85000.00, 850000.00, 0.00, 935000.00),
(9, 31500.00, 315000.00, 25000.00, 321500.00),
(10, 19000.00, 190000.00, 0.00, 209000.00);

INSERT INTO Orders
(Order_ID, PhoneNumber, OrderType, TotalAmountBeforeDiscount, TotalAmountAfterDiscount, OrderStatus, CustomerID, InvoiceID, DiscountID, TableID, Address)
VALUES
(1, '09121111111', 'dine_in', 275000.00, 275000.00, 4, 1, 1, NULL, 2, NULL),
(2, '09123333333', 'delivery', 410000.00, 369000.00, 3, 3, 2, 1, NULL, 'Saadat Abad, Kaj Square, No. 5'),
(3, '09122222222', 'takeaway', 145000.00, 145000.00, 4, 2, 3, NULL, NULL, NULL),
(4, '02188888888', 'delivery', 520000.00, 470000.00, 4, 4, 4, 6, NULL, 'Jordan, Nahid St, Office Tower, Unit 2'),
(5, '09125555555', 'dine_in', 120000.00, 120000.00, 2, 6, 5, NULL, 7, NULL),
(6, '09128888888', 'dine_in', 640000.00, 576000.00, 1, 9, 6, 1, 4, NULL),
(7, '09124444444', 'takeaway', 210000.00, 210000.00, 4, 5, 7, NULL, NULL, NULL),
(8, '02177777777', 'delivery', 850000.00, 850000.00, 2, 8, 8, NULL, NULL, 'Arjantin Square, Alvand St, No. 88'),
(9, '09129999999', 'delivery', 315000.00, 290000.00, 4, 10, 9, 6, NULL, 'Isfahan, Chaharbagh Abbasi, Sepahan Alley, No. 1'),
(10, '09126666666', 'takeaway', 190000.00, 190000.00, 5, 7, 10, NULL, NULL, NULL);

INSERT INTO Order_Food_Contains
(Order_ID, Food_ID, UnitPriceAtOrder, Quantity)
VALUES
(1, 1, 250000.00, 1),
(1, 5, 25000.00, 1),
(2, 2, 220000.00, 1),
(2, 8, 190000.00, 1),
(3, 4, 120000.00, 1),
(3, 5, 25000.00, 1),
(4, 3, 180000.00, 2),
(4, 6, 75000.00, 2);

INSERT INTO Payment_transaction
(Transaction_ID, Amount, Tracking_code, Order_ID, Payment_Status)
VALUES
(1, 302500.00, 'TRX-987654321', 1, 1),
(2, 410000.00, 'TRX-123456789', 2, 1),
(3, 159500.00, 'TRX-456123789', 3, 1),
(4, 522000.00, 'TRX-789456123', 4, 1),
(5, 132000.00, 'TRX-321654987', 5, 3),
(6, 640000.00, 'TRX-654987321', 6, 1),
(7, 231000.00, 'TRX-159753486', 7, 1),
(8, 935000.00, 'TRX-753159846', 8, 1),
(9, 321500.00, 'TRX-852963741', 9, 1),
(10, 209000.00, 'TRX-963852147', 10, 2);

