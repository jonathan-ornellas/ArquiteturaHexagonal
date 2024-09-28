package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/codeedu/go-hexagonal/adapters/db"
	"github.com/codeedu/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _= sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}


func createTable(db *sql.DB){
	createTableSql := `
	CREATE TABLE products (
		id TEXT NOT NULL PRIMARY KEY,
		name TEXT,
		price REAL,
		status TEXT
	);
	`
	stmt , err := db.Prepare(createTableSql)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
	
}


func createProduct (db *sql.DB){
	insertProductSql := `
	INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?)
	`
	stmt, err := db.Prepare(insertProductSql)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = stmt.Exec("abc", "Product Test", 1500.0, "disabled")
	if err != nil {
		log.Fatal(err.Error())
	}
}


func TestProductDb_Get(t *testing.T){
	setUp()
	defer Db.Close() // Close the connection after the test
	productDb :=  db.NewProductDb(Db)
	product, err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t,  1500.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}


func TestProducDb(t *testing.T){
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 1500.0


	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())


	product.Status = "enabled"

	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Status, productResult.GetStatus())

	


}