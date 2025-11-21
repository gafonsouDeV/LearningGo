package models

import (
	"fmt"
	"strings"
)

type Product struct {
	id    uint
	name  string
	price float32
	stock uint32
}

func (product *Product) CreateProduct(id uint, name string, price float32, stock uint32) {
	product.id = id
	product.name = strings.TrimSpace(name)
	product.price = price
	product.stock = stock
}

func (product *Product) ShowProductInfo() {
	const format = "%-3d %-20s %-10.2f %-10d\n"
	fmt.Printf(format, product.id, product.name, product.price, product.stock)
}
