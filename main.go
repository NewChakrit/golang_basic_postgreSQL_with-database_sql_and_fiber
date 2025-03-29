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

	// เมื่อดึง row เสร็จจะทำ db.Close() = ปิด connection database เพื่อปรัหยัด connection ที่ต่อไปยัง database (connection จะ ไม่สะสม)
	defer db.Close()

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
	//product, err := getProduct(2)

	//fmt.Println("Get Successful !", product)

	// ----- Get All Products ----- //

	products, err := getProducts()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(products)

	// ----- Update Product ----- //
	//product, err := updateProduct(6, &Product{Name: "New", Price: 555})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Update Product Successful !", product)

	// ----- Delete Product ----- //
	//err = deleteProduct(7)
	//
	//fmt.Println("Delete Product Successful !")

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

func getProducts() ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price from products")
	if err != nil {
		return nil, err
	}

	var products []Product

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func updateProduct(id int, product *Product) (Product, error) {
	var p Product
	//_, err := db.Exec(
	//	"UPDATE public.products SET name=$1, price=$2 WHERE id=$3;",
	//	product.Name,
	//	product.Price,
	//	id)

	row := db.QueryRow(
		"UPDATE public.products SET name=$1, price=$2 WHERE id=$3 RETURNING id, name, price;",
		product.Name,
		product.Price,
		id)

	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}

	return p, err
}

func deleteProduct(id int) error {
	_, err := db.Exec(
		"DELETE FROM public.products WHERE id=$1;", id)

	return err
}
