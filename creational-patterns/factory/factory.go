package main

import (
	"fmt"
)

// Burger interface
type Burger interface {
	Ingredients() []string
}

// Burger structs
type CheeseBurger struct{}

func (c CheeseBurger) Ingredients() []string {
	return []string{"Bun", "Cheese", "Beef Patty", "Lettuce", "Tomato"}
}

type VeggieBurger struct{}

func (v VeggieBurger) Ingredients() []string {
	return []string{"Bun", "Lettuce", "Tomato", "Cucumber", "Veggie Patty"}
}

type ChickenBurger struct{}

func (c ChickenBurger) Ingredients() []string {
	return []string{"Bun", "Chicken Patty", "Lettuce", "Tomato", "Mayonnaise"}
}

type FishBurger struct{}

func (f FishBurger) Ingredients() []string {
	return []string{"Bun", "Fish Patty", "Lettuce", "Tomato", "Tartar Sauce"}
}

// BurgerFactory struct
type BurgerFactory struct{}

func (bf BurgerFactory) CreateBurger(burgerType string) Burger {
	switch burgerType {
	case "CheeseBurger":
		return CheeseBurger{}
	case "VeggieBurger":
		return VeggieBurger{}
	case "ChickenBurger":
		return ChickenBurger{}
	case "FishBurger":
		return FishBurger{}
	default:
		// return nil or throw an error in case of invalid type
		return nil
	}
}

func main() {
	factory := BurgerFactory{}

	burger1 := factory.CreateBurger("CheeseBurger")
	fmt.Println(burger1.Ingredients())

	burger2 := factory.CreateBurger("VeggieBurger")
	fmt.Println(burger2.Ingredients())

	burger3 := factory.CreateBurger("ChickenBurger")
	fmt.Println(burger3.Ingredients())

	burger4 := factory.CreateBurger("FishBurger")
	fmt.Println(burger4.Ingredients())
}
