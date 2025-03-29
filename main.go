package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"log"
	"strconv"
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
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
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

	// ----- Add Fiber ----- //
	app := fiber.New()

	app.Post("product", createProductHandler)
	app.Get("/product/:id", getProductHandler)
	app.Get("/products", getAllProductHandler)
	app.Put("product/:id", updateProductHandler)
	app.Delete("product/:id", deleteProductHandler)

	app.Listen(":8080")

	// Connection Database Successful
	//fmt.Println("Connection Database Successful")

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

	//products, err := getProducts()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(products)

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

func createProductHandler(c *fiber.Ctx) error {
	p := new(Product)
	if err := c.BodyParser(p); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := createProduct(p)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(p)
}

func getProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	product, err := getProduct(productId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(product)
}

func getAllProductHandler(c *fiber.Ctx) error {
	products, err := getProducts()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(products)
}

func updateProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	p := new(Product)
	if err := c.BodyParser(p); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	product, err := updateProduct(productId, p)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(product)
}

func deleteProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = deleteProduct(productId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
