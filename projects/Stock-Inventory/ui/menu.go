package ui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/gafonsouDeV/LearningGo/projects/Stock-Inventory/models"
	"github.com/gafonsouDeV/LearningGo/projects/Stock-Inventory/utils"
)

func Menu() {
	fmt.Println("You're the new manager of the store. First create and inventory and manage the products you'll sell")

	// closure for id generation
	idIndex := func() func() uint {
		var index uint = 0
		return func() uint {
			index += 1
			return index
		}
	}()
	var inventory models.Inventory
	inventory.CreateInventory()

	for {
		showOptionsText()
		operation, err := utils.ReadNumber[int32]("Enter the operation")
		if err != nil || operation < 0 {
			fmt.Println("That's not a valid operation")
			time.Sleep(1 * time.Second)
			continue
		}

		if operation == 0 {
			break
		}

		operate(uint(operation), &inventory, idIndex)
		cleanMenu()
	}

}

func showOptionsText() {
	fmt.Println("Enter the option: ")
	fmt.Println("1. Create new Product.\n2. List all Products.")
	fmt.Println("3. Update a Product.\n4. Remove a Product")
	fmt.Println("0. Exit the system")
}

func operate(option uint, inv *models.Inventory, idIndex func() uint) {

	operations := map[uint]func(){
		1: func() { create(idIndex, inv) },
		2: func() { list(inv) },
		3: func() { update(inv) },
		4: func() { remove(inv) },
	}

	if operation, exist := operations[uint(option)]; exist {
		operation()
		WaitForKeypress()
	} else {
		fmt.Println("Invalid option")
	}
}

func create(idIndex func() uint, inv *models.Inventory) {
	var newProduct models.Product
	id := idIndex()
	name := utils.ReadString("Enter the product name: ")
	price := utils.ReadValidatedNumber[float32]("Enter the product price")
	stock := uint32(utils.ReadValidatedNumber[int32]("Enter the stock"))

	newProduct.CreateProduct(id, name, price, stock)
	inv.AddProduct(newProduct)
}

func list(inv *models.Inventory) {
	inv.ListAll()
}

func update(inv *models.Inventory) {
	productId := utils.ReadValidatedNumber[int32]("Enter the product id you want to update: ")
	fmt.Println("1. Product Name\n2. Product Price\n3.Product Stock ")
	update := utils.ReadValidatedNumber[int32]("Enter the product what you want to update: ")
	fieldsMap := map[int32]string{
		1: "name",
		2: "price",
		3: "stock",
	}
	var updatedField any
	switch update {
	case 1:
		updatedField = utils.ReadString("Type the new product name: ")
	case 2:
		updatedField = utils.ReadValidatedNumber[float32]("Type the new product price: ")
	case 3:
		updatedField = uint32(utils.ReadValidatedNumber[int32]("Type the new product stock: "))
	}

	inv.UpdateProduct(uint32(productId), fieldsMap[update], updatedField)
}

func remove(inv *models.Inventory) {
	productId := utils.ReadValidatedNumber[int32]("Enter the product id you want to remove: ")

	inv.RemoveProduct(uint32(productId))
}

func cleanMenu() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default: // Linux, macOS
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func WaitForKeypress() {
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
}
