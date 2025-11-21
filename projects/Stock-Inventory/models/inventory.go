package models

import (
	"fmt"
	"slices"
	"strings"
)

type Inventory struct {
	products []Product
}

type InventoryRepository interface {
	AddProduct(product Product)
	ListAll()
}

func (inventory *Inventory) CreateInventory() {
	inventory.products = make([]Product, 0)
}

func (inventory *Inventory) AddProduct(product Product) {
	inventory.products = append(inventory.products, product)
}

func (inventory *Inventory) RemoveProduct(id uint32) {
	inventory.products = slices.DeleteFunc(inventory.products, func(product Product) bool {
		return product.id == uint(id)
	})
}

func (inventory *Inventory) UpdateProduct(id uint32, field string, value any) {

	for i := range inventory.products {
		if inventory.products[i].id == uint(id) {
			switch field {
			case "name":
				inventory.products[i].name = strings.TrimSpace(value.(string))
			case "price":
				inventory.products[i].price = value.(float32)
			case "stock":
				inventory.products[i].stock = value.(uint32)
			}
			break
		}
	}

}

func (inventory *Inventory) ListAll() {
	const format = "%-3s %-20s %-10s %-10s\n"

	fmt.Printf(format, "ID", "Name", "Price", "Stock")
	fmt.Println()
	for _, product := range inventory.products {
		product.ShowProductInfo()
		fmt.Println()
	}
}
