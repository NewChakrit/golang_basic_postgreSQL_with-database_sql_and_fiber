package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host         = "localhost "
	port         = 5432
	databaseName = "mydatabase"
	username     = "myuser"
	password     = "mypassword"
)

var db *sql.DB

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	sdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	db = sdb

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Connection Database Successful
	fmt.Println("Connection Database Successful")

	// ----- Create Product ----- //
	//err = createProduct(&Product{Name: "Go product 2", Price: 400})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Create Successful !")

	// ----- Get Product ----- //

	product, err := getProduct(2)

	fmt.Println("Get Successful !", product)
}

func createProduct(product *Product) error {
	_, err := db.Exec(
		"Insert into public.products(name, price) values ($1, $2);",
		product.Name,
		product.Price)

	return err
}

func getProduct(id int) (Product, error) {
	var p Product
	row := db.QueryRow("select id, name, price from products where id=$1;", id)

	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}

	return p, nil
}
