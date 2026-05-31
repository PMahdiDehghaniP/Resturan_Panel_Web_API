package customers

var GET_ALL_CUSTOMERS = `SELECT customerid,first_name,
last_name,email,phone_number,customer_type,join_date
from customer order by customerid ASC LIMIT $1 OFFSET $2`
