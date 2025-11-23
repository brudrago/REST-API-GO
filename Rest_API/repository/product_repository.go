package repository

import (
	"database/sql"
	"fmt"
	"rest-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(conn *sql.DB) *ProductRepository {
	return &ProductRepository{
		connection: conn,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}
	defer rows.Close()

	products := []model.Product{}
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}
		products = append(products, p)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	return products, nil
}

func (pr *ProductRepository) CreateProduct(p model.Product) (int, error) {
	query := "INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id"
	var lastInsertID int
	err := pr.connection.QueryRow(query, p.Name, p.Price).Scan(&lastInsertID)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return lastInsertID, nil
}

func (pr *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	query := "SELECT id, product_name, price FROM product WHERE id=$1"
	var p model.Product
	err := pr.connection.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}
	return &p, nil
}

func (pr *ProductRepository) UpdateProduct(p model.Product) error {
	query := "UPDATE product SET product_name=$1, price=$2 WHERE id=$3"
	_, err := pr.connection.Exec(query, p.Name, p.Price, p.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (pr *ProductRepository) DeleteProduct(id int) error {
	query := "DELETE FROM product WHERE id=$1"
	_, err := pr.connection.Exec(query, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
