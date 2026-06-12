package foods

var CREATE_FOOD = `
INSERT INTO food_item
(name, description, calories, imagepath, unitprice, inventory, foodcategory)
VALUES
($1, $2, $3, $4, $5, $6, $7)
RETURNING food_id
`
