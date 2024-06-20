package main

import (
	"fmt"
	"strings"
)

// Burger struct
type Burger struct {
	ingredients []string
}

func (b Burger) String() string {
	return strings.Join(b.ingredients, ", ")
}

// BurgerBuilder interface
type BurgerBuilder interface {
	AddBuns(bunType string) BurgerBuilder
	AddPatty(pattyType string) BurgerBuilder
	AddCheese(cheeseType string) BurgerBuilder
	AddLettuce() BurgerBuilder
	AddTomato() BurgerBuilder
	Build() Burger
}

// CustomBurgerBuilder struct
type CustomBurgerBuilder struct {
	burger Burger
}

func (cbb *CustomBurgerBuilder) AddBuns(bunType string) BurgerBuilder {
	cbb.burger.ingredients = append(cbb.burger.ingredients, bunType+" Buns")
	return cbb
}

func (cbb *CustomBurgerBuilder) AddPatty(pattyType string) BurgerBuilder {
	cbb.burger.ingredients = append(cbb.burger.ingredients, pattyType+" Patty")
	return cbb
}

func (cbb *CustomBurgerBuilder) AddCheese(cheeseType string) BurgerBuilder {
	cbb.burger.ingredients = append(cbb.burger.ingredients, cheeseType+" Cheese")
	return cbb
}

func (cbb *CustomBurgerBuilder) AddLettuce() BurgerBuilder {
	cbb.burger.ingredients = append(cbb.burger.ingredients, "Lettuce")
	return cbb
}

func (cbb *CustomBurgerBuilder) AddTomato() BurgerBuilder {
	cbb.burger.ingredients = append(cbb.burger.ingredients, "Tomato")
	return cbb
}

func (cbb *CustomBurgerBuilder) Build() Burger {
	return cbb.burger
}

func main() {
	builder := &CustomBurgerBuilder{}

	burger1 := builder.AddBuns("Sesame").
		AddPatty("Beef").
		AddCheese("Cheddar").
		AddLettuce().
		AddTomato().
		Build()

	burger2 := builder.AddBuns("Whole Wheat").
		AddPatty("Chicken").
		AddCheese("Swiss").
		AddLettuce().
		Build()

	fmt.Println(burger1)
	fmt.Println(burger2)
}
