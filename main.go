package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// true if ascending order and false if descending
type Column struct{
	Name string
	Asc bool
}

var Options []Column


func main(){
	/* 	N.B! Comment out createTable() and addRandomCustomersToDB() after the first start of the app.
	If you leave those 2 methods uncommented on the second start, you will lose all new data from the db. */

	// creates empty table
	// createTable()

	// imports 5 new customers to the table
	// addRandomCustomersToDB()

	handleFunctions()
}


// app setters
func createTable() {
	// CONNECTING TO THE DATABASE
	connStr := "user=postgres dbname=crud-in-go password=12345 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// CLOSING CONNECTION
	defer db.Close()

	// DROPPING TABLE IF EXISTS
	drop, err := db.Query("DROP TABLE IF EXISTS customers")
	if err != nil {
		panic(err)
	}

	// CLOSING DROP-QUERY CONNECTION
	drop.Close()

	// CREATING NEW TABLE
	create, err := db.Query("CREATE TABLE IF NOT EXISTS customers ( id serial PRIMARY KEY, first_name varchar(100) NOT NULL, last_name varchar(100) NOT NULL, dob DATE NOT NULL, gender varchar(10) NOT NULL, email varchar(100) NOT NULL UNIQUE, address varchar(200));")
	if err != nil {
		panic(err)
	} else {
	}

	// CLOSING CREATE-QUERY CONNECTION
	defer create.Close()
}

func addRandomCustomersToDB() {
	// CONNECTING TO THE DATABASE
	connStr := "user=postgres dbname=crud-in-go password=12345 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// CLOSING CONNECTION
	defer db.Close()

	// INSERTING DATA INTO THE TABLE
	insert, err := db.Query("INSERT INTO customers (first_name, last_name, dob, gender, email, address) " +
		"VALUES('Clint', 'Yamchinsky', '1974-05-17', 'Male', 'clint@gmail.com', '18 Armley road, Leeds')," +
		"('Vicky', 'Van Strauzberg', '1996-12-07', 'Female', 'vicky@gmail.com', '32-154 Gutten road, Amsterdam')," +
		"('Johny', 'Bravo', '1968-10-18', 'Male', 'johny@bravo.com', '214 Golden Beach Apartments, Orlando')," +
		"('Ben', 'Golberg', '1975-02-21', 'Male', 'ben.golberg@gmail.com', '13 Swansley avenue, Beverley Hills')," +
		"('Amanda', 'Bower', '1980-05-07', 'Female', 'amanda.bower@gmail.com', '377 Skyline scrapper, Miami') ON CONFLICT DO NOTHING")
	if err != nil {
		panic(err)
	}

	// CLOSING INSERT-QUERY
	defer insert.Close()
}

func handleFunctions() {
	// loading options for column ordering
	Options = LoadOptions()

	router := mux.NewRouter()

	// rendering local css files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// importing mux router
	http.Handle("/", router)


	// rendering pages
	router.HandleFunc("/", homePageHandler).Methods("GET")
	router.HandleFunc("/create", createCustomerPageHandler).Methods("GET")
	router.HandleFunc("/create-customer", createCustomer).Methods("POST")
	router.HandleFunc("/confirmation", confirmationPageHandler).Methods("GET")
	router.HandleFunc("/view-customers", viewCustomersPageHandler).Methods("GET")
	router.HandleFunc("/view-customer/{id:[0-9]+}", customerInfoPageHandler).Methods("GET")
	router.HandleFunc("/update-customer/{id:[0-9]+}", updateCustomerPageHandler).Methods("GET")
	router.HandleFunc("/confirm-update-customer/{id:[0-9]+}", updateCustomer).Methods("POST")
	router.HandleFunc("/delete-customer/{id:[0-9]+}", deleteCustomer).Methods("POST")
	router.HandleFunc("/search-customer", searchCustomer).Methods("POST")

	//sorting list of customers
	router.HandleFunc("/view-customers/{order}", sortByOrder).Methods("GET")


	//server starts
	log.Println("listens on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}



// page handlers
func homePageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/home.tmpl", "templates/header.tmpl", "templates/footer.tmpl", "templates/script.tmpl")
	if err != nil {
		panic(err)
	}
	Options = LoadOptions()

	_ = t.ExecuteTemplate(w, "home", nil)
}

func createCustomerPageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.tmpl", "templates/header.tmpl", "templates/footer.tmpl", "templates/script.tmpl")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	Options = LoadOptions()
	_ = t.ExecuteTemplate(w, "create", nil)

}

func confirmationPageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/confirmation.tmpl", "templates/header.tmpl", "templates/footer.tmpl", "templates/script.tmpl")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	Options = LoadOptions()
	_ = t.ExecuteTemplate(w, "confirmation", nil)
}

func viewCustomersPageHandler(w http.ResponseWriter, r *http.Request) {
	var Customers []Customer

	// RENDER TEMPLATES
	t, err := template.ParseFiles("templates/view-customers.tmpl", "templates/header.tmpl", "templates/footer.tmpl", "templates/script.tmpl")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	// CONNECTING TO THE DATABASE
	connStr := "user=postgres dbname=crud-in-go password=12345 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// CLOSING CONNECTION
	defer db.Close()

	// EXPORTING DATA FROM THE DATABASE
	export, err := db.Query("SELECT * FROM customers ORDER BY id ASC")
	if err != nil {
		panic(err)
	}

	for export.Next() {
		var c Customer
		err = export.Scan(&c.Id, &c.FirstName, &c.LastName, &c.DOB, &c.Gender, &c.Email, &c.Address)
		if err != nil {
			panic(err)
		}

		Customers = append(Customers, c)
	}

	// CLOSING EXPORT-QUERY
	defer export.Close()

	Options = LoadOptions()
	_ = t.ExecuteTemplate(w, "view_customers", Customers)
}

func customerInfoPageHandler(w http.ResponseWriter, r *http.Request) {
	var c Customer

	// RENDER TEMPLATES
	t, err := template.ParseFiles("templates/view-customer.tmpl", "templates/header.tmpl", "templates/footer.tmpl", "templates/script.tmpl")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	// GETTING ID FROM THE ROUTER
	vars := mux.Vars(r)
	//w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w, "ID: %v\n", vars["id"])

	// CONNECTING TO THE DATABASE
	connStr := "user=postgres dbname=crud-in-go password=12345 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// CLOSING CONNECTION
	defer db.Close()

	// EXPORTING DATA FROM THE DATABASE
	export, err := db.Query(fmt.Sprintf("SELECT * FROM customers WHERE id='%v'", vars["id"]))
	if err != nil {
		panic(err)
	}

	for export.Next() {
		err = export.Scan(&c.Id, &c.FirstName, &c.LastName, &c.DOB, &c.Gender, &c.Email, &c.Address)
		if err != nil {
			panic(err)
		}
	}

	// CLOSING EXPORT-QUERY
	defer export.Close()

	Options = LoadOptions()
	_ = t.ExecuteTemplate(w, "customer_info", c)

}

func updateCustomerPageHandler(w http.ResponseWriter, r *http.Request) {
	var c Customer
	var cf CustomerForm

	// RENDER TEMPLATES
	t, err := template.ParseFiles("templates/update-details.tmpl", "templates/header.tmpl", "templates/footer.tmpl", "templates/script.tmpl")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	// GETTING ID FROM THE ROUTER
	vars := mux.Vars(r)

	// CONNECTING TO THE DATABASE
	connStr := "user=postgres dbname=crud-in-go password=12345 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// CLOSING CONNECTION
	defer db.Close()

	// EXPORTING DATA FROM THE DATABASE
	export, err := db.Query(fmt.Sprintf("SELECT * FROM customers WHERE id='%v'", vars["id"]))
	if err != nil {
		panic(err)
	}

	for export.Next() {
		err = export.Scan(&c.Id, &c.FirstName, &c.LastName, &c.DOB, &c.Gender, &c.Email, &c.Address)
		if err != nil {
			panic(err)
		}
	}

	// CLONING CUSTOMER DETAILS TO CUSTOMER FORM
	cf.Id = c.Id
	cf.FirstName = c.FirstName
	cf.LastName = c.LastName
	cf.DOB = c.FormatDOB()
	cf.Gender = c.Gender
	cf.Email = c.Email
	cf.Address = c.Address

	// CLOSING EXPORT-QUERY
	defer export.Close()

	Options = LoadOptions()
	_ = t.ExecuteTemplate(w, "update", cf)
}


// sorting group of methods
func sortByOrder(w http.ResponseWriter, r *http.Request) {
	// GETTING ORDER FROM THE ROUTER
	vars := mux.Vars(r)

	switch strings.ToLower(vars["order"]){
	case "by-id": sortById(w, r)
	case "by-first-name": sortByFirstName(w, r)
	case "by-last-name": sortByLastName(w, r)
	case "by-birthday": sortByDOB(w, r)
	case "by-gender": sortByGender(w, r)
	case "by-email": sortByEmail(w, r)
	case "by-address": sortByAddress(w, r)
	default:
		// NAVIGATE TO BAD STATUS PAGE
		http.Redirect(w, r, "/", http.StatusBadRequest)
	}
}

func sortById(w http.ResponseWriter, r *http.Request) {
	var Customers []Customer
	colName := "id"
	var col Column

	for _, c := range Options{
		if c.Name == colName{
			col = c
			break
		}
	}


	// RENDER TEMPLATES
	t, err := template.ParseFiles("templates/view-customers.tmpl", "templates/header.tmpl", "templates/footer.tmpl", "templates/script.tmpl")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	// CONNECTING TO THE DATABASE
	connStr := "user=postgres dbname=crud-in-go password=12345 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// CLOSING CONNECTION
	defer db.Close()


	// EXPORTING DATA FROM THE DATABASE
	export, err := db.Query(fmt.Sprintf("SELECT * FROM customers ORDER BY id %s", col.Order()))
	if err != nil {
		panic(err)
	}

	for export.Next() {
		var c Customer
		err = export.Scan(&c.Id, &c.FirstName, &c.LastName, &c.DOB, &c.Gender, &c.Email, &c.Address)
		if err != nil {
			panic(err)
		}

		Customers = append(Customers, c)
		// CLOSING EXPORT-QUERY
		defer export.Close()
	}

	_ = t.ExecuteTemplate(w, "view_customers", Customers)
}

func sortByFirstName(w http.ResponseWriter, r *http.Request) {
	var Customers []Customer
	colName := "first"
	var col Column

	for _, c := range Options{
		if c.Name == colName{
			col = c
			break
		}
	}


	// RENDER TEMPLATES
	t, err := template.ParseFiles("templates/view-customers.tmpl", "templates/header.tmpl", "templates/footer.tmpl", "templates/script.tmpl")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	// CONNECTING TO THE DATABASE
	connStr := "user=postgres dbname=crud-in-go password=12345 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// CLOSING CONNECTION
	defer db.Close()


	// EXPORTING DATA FROM THE DATABASE
	export, err := db.Query(fmt.Sprintf("SELECT * FROM customers ORDER BY first_name %s", col.Order()))
	if err != nil {
		panic(err)
	}

	for export.Next() {
		var c Customer
		err = export.Scan(&c.Id, &c.FirstName, &c.LastName, &c.DOB, &c.Gender, &c.Email, &c.Address)
		if err != nil {
			panic(err)
		}

		Customers = append(Customers, c)
		// CLOSING EXPORT-QUERY
		defer export.Close()
	}

	_ = t.ExecuteTemplate(w, "view_customers", Customers)
}

func sortByLastName(w http.ResponseWriter, r *http.Request) {
	var Customers = []Customer{}
	colName := "last"
	var col Column

	for _, c := range Options{
		if c.Name == colName{
			col = c
			break
		}
	}

	// RENDER TEMPLATES
	t, err := template.ParseFiles("templates/view-customers.tmpl", "templates/header.tmpl", "templates/footer.tmpl", "templates/script.tmpl")
	if err != nil {
		panic(err)
	}

	// CONNECTING TO THE DATABASE
	connStr := "user=postgres dbname=crud-in-go password=12345 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// CLOSING CONNECTION
	defer db.Close()

	// EXPORTING DATA FROM THE DATABASE
	export, err := db.Query(fmt.Sprintf("SELECT * FROM customers ORDER BY last_name %s", col.Order()))
	if err != nil {
		panic(err)
	}

	for export.Next() {
		var c Customer
		err = export.Scan(&c.Id, &c.FirstName, &c.LastName, &c.DOB, &c.Gender, &c.Email, &c.Address)
		if err != nil {
			panic(err)
		}

		Customers = append(Customers, c)
	}

	// CLOSING EXPORT-QUERY
	defer export.Close()

	_ = t.ExecuteTemplate(w, "view_customers", Customers)
}

func sortByGender(w http.ResponseWriter, r *http.Request) {
	var Customers = []Customer{}
	colName := "gender"
	var col Column

	for _, c := range Options{
		if c.Name == colName{
			col = c
			break
		}
	}

	// RENDER TEMPLATES
	t, err := template.ParseFiles("templates/view-customers.tmpl", "templates/header.tmpl", "templates/footer.tmpl", "templates/script.tmpl")
	if err != nil {
		panic(err)
	}

	// CONNECTING TO THE DATABASE
	connStr := "user=postgres dbname=crud-in-go password=12345 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// CLOSING CONNECTION
	defer db.Close()

	// EXPORTING DATA FROM THE DATABASE
	export, err := db.Query(fmt.Sprintf("SELECT * FROM customers ORDER BY gender %s", col.Order()))
	if err != nil {
		panic(err)
	}

	for export.Next() {
		var c Customer
		err = export.Scan(&c.Id, &c.FirstName, &c.LastName, &c.DOB, &c.Gender, &c.Email, &c.Address)
		if err != nil {
			panic(err)
		}

		Customers = append(Customers, c)
	}

	// CLOSING EXPORT-QUERY
	defer export.Close()

	_ = t.ExecuteTemplate(w, "view_customers", Customers)
}

func sortByDOB(w http.ResponseWriter, r *http.Request) {
	var Customers = []Customer{}
	colName := "dob"
	var col Column

	for _, c := range Options{
		if c.Name == colName{
			col = c
			break
		}
	}

	// RENDER TEMPLATES
	t, err := template.ParseFiles("templates/view-customers.tmpl", "templates/header.tmpl", "templates/footer.tmpl", "templates/script.tmpl")
	if err != nil {
		panic(err)
	}

	// CONNECTING TO THE DATABASE
	connStr := "user=postgres dbname=crud-in-go password=12345 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// CLOSING CONNECTION
	defer db.Close()

	// EXPORTING DATA FROM THE DATABASE
	export, err := db.Query(fmt.Sprintf("SELECT * FROM customers ORDER BY dob %s", col.Order()))
	if err != nil {
		panic(err)
	}

	for export.Next() {
		var c Customer
		err = export.Scan(&c.Id, &c.FirstName, &c.LastName, &c.DOB, &c.Gender, &c.Email, &c.Address)
		if err != nil {
			panic(err)
		}

		Customers = append(Customers, c)
	}

	// CLOSING EXPORT-QUERY
	defer export.Close()

	_ = t.ExecuteTemplate(w, "view_customers", Customers)
}

func sortByEmail(w http.ResponseWriter, r *http.Request) {
	var Customers = []Customer{}
	colName := "email"
	var col Column

	for _, c := range Options{
		if c.Name == colName{
			col = c
			break
		}
	}

	// RENDER TEMPLATES
	t, err := template.ParseFiles("templates/view-customers.tmpl", "templates/header.tmpl", "templates/footer.tmpl", "templates/script.tmpl")
	if err != nil {
		panic(err)
	}

	// CONNECTING TO THE DATABASE
	connStr := "user=postgres dbname=crud-in-go password=12345 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// CLOSING CONNECTION
	defer db.Close()

	// EXPORTING DATA FROM THE DATABASE
	export, err := db.Query(fmt.Sprintf("SELECT * FROM customers ORDER BY email %s", col.Order()))
	if err != nil {
		panic(err)
	}

	for export.Next() {
		var c Customer
		err = export.Scan(&c.Id, &c.FirstName, &c.LastName, &c.DOB, &c.Gender, &c.Email, &c.Address)
		if err != nil {
			panic(err)
		}

		Customers = append(Customers, c)
	}

	// CLOSING EXPORT-QUERY
	defer export.Close()

	_ = t.ExecuteTemplate(w, "view_customers", Customers)
}

func sortByAddress(w http.ResponseWriter, r *http.Request) {
	var Customers = []Customer{}
	colName := "address"
	var col Column

	for _, c := range Options{
		if c.Name == colName{
			col = c
			break
		}
	}

	// RENDER TEMPLATES
	t, err := template.ParseFiles("templates/view-customers.tmpl", "templates/header.tmpl", "templates/footer.tmpl", "templates/script.tmpl")
	if err != nil {
		panic(err)
	}

	// CONNECTING TO THE DATABASE
	connStr := "user=postgres dbname=crud-in-go password=12345 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// CLOSING CONNECTION
	defer db.Close()

	// EXPORTING DATA FROM THE DATABASE
	export, err := db.Query(fmt.Sprintf("SELECT * FROM customers ORDER BY address %s", col.Order()))
	if err != nil {
		panic(err)
	}

	for export.Next() {
		var c Customer
		err = export.Scan(&c.Id, &c.FirstName, &c.LastName, &c.DOB, &c.Gender, &c.Email, &c.Address)
		if err != nil {
			panic(err)
		}

		Customers = append(Customers, c)
	}

	// CLOSING EXPORT-QUERY
	defer export.Close()

	_ = t.ExecuteTemplate(w, "view_customers", Customers)
}


// rest controllers
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	// GETTING ID FROM THE ROUTER
	vars := mux.Vars(r)
	if len(vars) < 1 {
		panic("vars is empty")
	}

	// CONNECTING TO THE DATABASE
	connStr := "user=postgres dbname=crud-in-go password=12345 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	} else {
		// CLOSING CONNECTION
		defer db.Close()

		// DELETE DATA FROM THE DATABASE
		remove, err := db.Query(fmt.Sprintf("DELETE FROM customers WHERE id='%s'", vars["id"]))
		if err != nil {
			panic(err)
		}

		// CLOSING EXPORT-QUERY
		defer remove.Close()

		// NAVIGATE TO THE LIST OF CUSTOMERS
		Options = LoadOptions()
		http.Redirect(w, r, "/view-customers/by-id", http.StatusSeeOther)

	}

}

func searchCustomer(w http.ResponseWriter, r *http.Request){
	var Customers []Customer

	// RENDER TEMPLATES
	t, err := template.ParseFiles("templates/view-customers.tmpl", "templates/header.tmpl", "templates/footer.tmpl", "templates/script.tmpl")
	if err != nil {
		panic(err)
	}

	// SAVING DATA FROM FORM
	sf := &SearchForm{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
	}

	// VALIDATING SEARCH FORM
	if sf.ValidateSearchForm() == false {
		log.Println("Search validation failed")
		// REDIRECT TO THE SAME PAGE WITHOUT ANY CHANGE
		http.HandleFunc("/search-customer", viewCustomersPageHandler)

	}else{

		// CONNECTING TO THE DATABASE
		connStr := "user=postgres dbname=crud-in-go password=12345 host=localhost port=5432 sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}

		// CLOSING CONNECTION
		defer db.Close()

		if len(sf.FirstName) > 0 && len(sf.LastName) > 0{
			// SEARCH FOR BOTH NAMES
			export, err := db.Query(fmt.Sprintf("SELECT * FROM customers WHERE first_name='%s' AND last_name='%s' ORDER BY id ASC", sf.FirstName, sf.LastName ))
			if err != nil {
				panic(err)
			}

			for export.Next() {
				var c Customer
				err = export.Scan(&c.Id, &c.FirstName, &c.LastName, &c.DOB, &c.Gender, &c.Email, &c.Address)
				if err != nil {
					panic(err)
				}

				Customers = append(Customers, c)
			}
			// CLOSING EXPORT-QUERY
			defer export.Close()

		}else if len(sf.FirstName) > 0{
			// SEARCH FOR FIRST NAME ONLY
			export, err := db.Query(fmt.Sprintf("SELECT * FROM customers WHERE first_name='%s' ORDER BY id ASC", sf.FirstName))
			if err != nil {
				panic(err)
			}

			for export.Next() {
				var c Customer
				err = export.Scan(&c.Id, &c.FirstName, &c.LastName, &c.DOB, &c.Gender, &c.Email, &c.Address)
				if err != nil {
					panic(err)
				}

				Customers = append(Customers, c)
			}
			// CLOSING EXPORT-QUERY
			defer export.Close()

		}else{
			// SEARCH FOR LAST NAME ONLY
			export, err := db.Query(fmt.Sprintf("SELECT * FROM customers WHERE last_name='%s' ORDER BY id ASC", sf.LastName))
			if err != nil {
				panic(err)
			}

			for export.Next() {
				var c Customer
				err = export.Scan(&c.Id, &c.FirstName, &c.LastName, &c.DOB, &c.Gender, &c.Email, &c.Address)
				if err != nil {
					panic(err)
				}

				Customers = append(Customers, c)
			}
			// CLOSING EXPORT-QUERY
			defer export.Close()

		}
		Options = LoadOptions()
		_ = t.ExecuteTemplate(w, "view_customers", Customers)
	}

}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	// SAVING DATA FROM FORM
	cf := &CustomerForm{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Gender:    r.FormValue("gender"),
		DOB:       r.FormValue("dob"),
		Email:     r.FormValue("email"),
		Address:   r.FormValue("address"),
	}

	// VALIDATING CONTACT FORM FIELDS
	if cf.ValidateCreateForm() == false {
		t, err := template.ParseFiles("templates/create.tmpl", "templates/header.tmpl", "templates/footer.tmpl", "templates/script.tmpl")
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		_ = t.ExecuteTemplate(w, "create", cf)
		return
	} else {
		// CONNECTING TO THE DATABASE
		connStr := "user=postgres dbname=crud-in-go password=12345 host=localhost port=5432 sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}

		// CLOSING CONNECTION
		defer db.Close()

		// INSERTING DATA INTO THE TABLE
		insert, err := db.Query(fmt.Sprintf("INSERT INTO customers (first_name, last_name, dob, gender, email, address) "+
			"VALUES('%s', '%s', '%s', '%s', '%s', '%s');", cf.FirstName, cf.LastName, cf.DOB, cf.Gender, cf.Email, cf.Address))
		if err != nil {
			panic(err)
		}

		// CLOSING INSERT-QUERY
		defer insert.Close()

		// NAVIGATE TO CONFIRMATION PAGE
		Options = LoadOptions()
		http.Redirect(w, r, "/confirmation", http.StatusSeeOther)
	}
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	var c Customer

	// GETTING ID FROM THE ROUTER
	vars := mux.Vars(r)

	// CONNECTING TO THE DATABASE
	connStr := "user=postgres dbname=crud-in-go password=12345 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// CLOSING CONNECTION
	defer db.Close()

	// EXPORTING DATA FROM THE DATABASE
	export, err := db.Query(fmt.Sprintf("SELECT * FROM customers WHERE id='%v'", vars["id"]))
	if err != nil {
		panic(err)
	}

	for export.Next() {
		err = export.Scan(&c.Id, &c.FirstName, &c.LastName, &c.DOB, &c.Gender, &c.Email, &c.Address)
		if err != nil {
			panic(err)
		}
	}

	// CLOSING EXPORT-QUERY
	defer export.Close()

	// SAVING DATA FROM FORM
	cf := &CustomerForm{
		Id:        c.Id,
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		DOB:       r.FormValue("dob"),
		Gender:    c.Gender,
		Email:     r.FormValue("email"),
		Address:   r.FormValue("address"),
	}

	// VALIDATING CONTACT FORM FIELDS
	if cf.ValidateUpdateForm() == false {
		t, err := template.ParseFiles("templates/update-details.tmpl", "templates/header.tmpl", "templates/footer.tmpl", "templates/script.tmpl")
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		_ = t.ExecuteTemplate(w, "update", cf)
		return
	} else {
		// CONNECTING TO THE DATABASE
		connStr := "user=postgres dbname=crud-in-go password=12345 host=localhost port=5432 sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}

		// CLOSING CONNECTION
		defer db.Close()

		// UPDATING DATA IN THE TABLE
		update, err := db.Query(fmt.Sprintf("UPDATE customers SET first_name='%s', last_name='%s', dob='%s', email='%s', address='%s' WHERE  id='%d'", cf.FirstName, cf.LastName, cf.DOB, cf.Email, cf.Address, cf.Id))
		if err != nil {
			panic(err)
		}

		// CLOSING INSERT-QUERY
		defer update.Close()

		// NAVIGATE TO THE LIST OF CUSTOMERS
		Options = LoadOptions()
		http.Redirect(w, r, "/view-customers/by-id", http.StatusSeeOther)
	}
}


// helper methods
func LoadOptions() []Column {
	id := Column{"id", true}
	firstName := Column{"first",true}
	lastName := Column{"last",true}
	gender := Column{"gender",true}
	dob := Column{"dob",true}
	email := Column{"email", true}
	address := Column{"address",true}

	res := []Column{id,
		firstName,
		lastName,
		gender,
		dob,
		email,
		address,
	}

	return  res
}

// changing order for each column based on user choice and return asc or desc
func (c *Column) Order() string {
	var result string

	for i, col := range Options{
		if col.Name == c.Name{
			result = AscValue(col)
			Options = LoadOptions()
			(&Options[i]).Asc = c.ChangeAsc()
		}
	}
	return result
}

// gets string value of asc parameter of the column
func AscValue(col Column) string{
	if col.Asc == true{
		return "ASC"
	}else{
		return "DESC"
	}
}

// swap values from asc to desc and vice versa
func (c *Column) ChangeAsc() bool{
	if c.Asc{
		return false
	}
		return true
}


