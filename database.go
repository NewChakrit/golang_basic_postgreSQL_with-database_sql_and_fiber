package main

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
