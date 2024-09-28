package dto

import "github.com/codeedu/go-hexagonal/application"

type Product struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Status string  `json:"status"`
	Price  float64 `json:"price"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) Bind(product *application.Product) (*application.Product, error) {

	if p.Id != "" {
		product.Id = p.Id
	}
	product.Name = p.Name
	product.Status = p.Status
	product.Price = p.Price
	_, err := product.IsValid()
	if err != nil {
		return &application.Product{}, err
	}

	return product, nil
}