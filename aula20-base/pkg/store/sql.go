package store

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	_ "github.com/go-sql-driver/mysql"
)

type sqlStore struct {
	db *sql.DB
	*jsonStore
}

func NewSQLStore() StoreInterface {
	database, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/my_db")
	if err != nil {
		log.Fatalln(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &sqlStore{
		db: database,
	}
}

func (s *sqlStore) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM products WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}

const QUERY_INSERT    = `INSERT INTO products (
	name, 
	quantity, 
	code_value, 
	is_published,	
	expiration,	
	price) VALUES (?, ?, ?, ?, ?, ?)`


func (s *sqlStore) Create(p domain.Product) error {

	stmt, err := s.db.Prepare(QUERY_INSERT)

	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		p.Name,
		p.Quantity,
		p.CodeValue,
		p.IsPublished,
		p.Expiration,
		p.Price)
	if err != nil {
		return err
	}

	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return fmt.Errorf("Fail to save")
	}

	return nil
}

func (s *sqlStore) GetById(id int) (domain.Product, error) {

	queryGetById = "SELECT id, name, quantity, code_value, is_published, expiration, price FROM products where id = ?"

	row := s.db.QueryRowContext(queryGetById, id)

	product := domain.Product{}

	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Quantity,
		&product.CodeValue,
		&product.IsPublished,
		&product.Expiration,
		&product.Price,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return product, fmt.Errorf("product %d not found", id)
	}

	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *sqlStore) Update(id int, p domain.Product) (domain.Product, error) {

	queryUpdate  := "UPDATE products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? WHERE id = ?"

	product, err := s.GetById(id)
	if err != nil {
		return domain.Product{}, fmt.Errorf("product %d not found", id)
	}

	if p.Name != "" {
		product.Name = p.Name
	}
	if p.Quantity > 0 {
		product.Quantity = p.Quantity
	}
	if p.CodeValue != "" {
		product.CodeValue = p.CodeValue
	}
	if p.Expiration != "" {
		product.Expiration = p.Expiration
	}
	if p.Price > 0.0 {
		product.Price = p.Price
	}

	product.IsPublished = p.IsPublished

	result, err := s.db.Exec(
		queryUpdate,
		product.Name,
		product.Quantity,
		product.CodeValue,
		product.IsPublished,
		product.Expiration,
		product.Price,
		id,
	)
	if err != nil {
		return domain.Product{}, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return domain.Product{}, err
	}
	log.Println(affectedRows)

	return product, nil
}
