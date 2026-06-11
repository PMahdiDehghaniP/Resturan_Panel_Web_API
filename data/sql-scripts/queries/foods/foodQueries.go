package foods

var GET_ALL_FOOD = `SELECT name AS foodName,description,calories,imagepath,unitprice,inventory,fc.category_name
    AS foodCategory  FROM food_item fi JOIN food_category fc
ON fi.foodcategory = fc.category_id`
