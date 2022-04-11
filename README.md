## 177.-Picus-Security-Go-Bootcamp-Bitirme-Projesi



* Login and register\
*http://localhost:8090/user/login
http://localhost:8090/user/register*

```
{
"email": "xxx@hotmail.cpö",
"password":"xxxdxxxxx"
}
```

***
***
* Adds item to cart\
User can add an item to cart with using this body
*http://localhost:8090/cart/add*

```
{
"ID": 1, // product id
"Amount":5
}
```
***
***
* Deletes item from cart\
User can delete an item from the cart by product id\
*http://localhost:8090/cart/delete*
`http://localhost:8090/cart/delete?id=1`
***
***

* Creates product\
If user is admin, it can create product with using this body\
*http://localhost:8090/product/create*
```
{
"product_name" : "addibffds",
"price" :10,
"stock": 30,
"category_name" : "giysi",
"sku" : "32422374"
}
```
***
***
* Deletes product\
If user is admin, it deletes product by product id\
*http://localhost:8090/product/delete*
`http://localhost:8090/product/delete?id=2`
***
***
* Lists products\
Lists products by pagination\
*http://localhost:8090/product/list*
`http://localhost:8090/product/list?pageSize=5&page=2`
***
***

* Updates product\
If user's role is admin, then product can be updated. id is required.\
*http://localhost:8090/product/update*
```
id
stock
price
name
sku
http://localhost:8090/product/update?sku=2&id=4&price=15
```
***
***
* Lists products in cart\
User can list all products in cart\
`http://localhost:8090/cart/list`
***
***
* Updates product in cart\
User can update amount of product in cart by product id\
`http://localhost:8090/cart/update`
```
{
"ID": 2, // product id
"Amount":3
}
```
***
***
* Searches products\
It searches by product name, and sku and returns all products that match the search\
*http://localhost:8090/product/search*
`http://localhost:8090/product/search?search=324&pageSize=3&page=2`
***
***
* Lists user's all orders\
`http://localhost:8090/order/list`
***
***
* Lists order details\
User can see his order details by order id\
*http://localhost:8090/order/detail*
` http://localhost:8090/order/details?id=1`
***
***
* Creates order\
It creates order and updates user's cart\
`http://localhost:8090/order/create`
***
***
* Cancels order\
In 14 days after order, it can be cancelled. id is order id and required \
*http://localhost:8090/order/cancel*
`http://localhost:8090/order/cancel?id=1`
***
***
* Creates category\
If user's role is admin then user can create category by reading csv file\
*http://localhost:8090/category/create*
***
***
* List all categories\
*http://localhost:8090/category/list*