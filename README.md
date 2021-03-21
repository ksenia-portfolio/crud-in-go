# CRUD IN GO WITH POSTGRES
<hr>
This simple application is build with go and postgres db. 

It contains following functionalities:
* Create new customer
* Read information about a customer
* Update customer details
* Delete customer
* Search customer by first or/and last name
* Sort list by any parameter in a table ascending and descending

## Installation

1. Clone repository and import the link into your IDE.
   <br><br>
2. If you run app for the first time, please, uncomment the following methods in <strong> main.go </strong> file:
   1.  createTable() 
   2.  addRandomCustomersToDB()
   <br><br>
2.  Do not forget to comment them back to avoid loosing new data from the database.
    <br><br>


## Project imperfections

Application has some bugs: 
* Bootstrap example page was used, js do not work as should and cause some errors in HTML elements alignment.
* Sorting function in searching mode returns all customers list and not searched ones.
* Pagination is not implemented.
* No tests implemented.

<hr>