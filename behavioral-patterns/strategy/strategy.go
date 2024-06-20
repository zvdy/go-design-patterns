package main

import "fmt"

// Strategy interface
type FilterStrategy interface {
	Filter([]int) []int
}

// Concrete strategy 1
type RemoveNegativeStrategy struct{}

func (rns *RemoveNegativeStrategy) Filter(nums []int) []int {
	result := []int{}
	for _, num := range nums {
		if num >= 0 {
			result = append(result, num)
		}
	}
	return result
}

// Concrete strategy 2
type RemovePositiveStrategy struct{}

func (rps *RemovePositiveStrategy) Filter(nums []int) []int {
	result := []int{}
	for _, num := range nums {
		if num <= 0 {
			result = append(result, num)
		}
	}
	return result
}

// Concrete strategy 3
type RemoveMultiplesOfThreeStrategy struct{}

func (rmots *RemoveMultiplesOfThreeStrategy) Filter(nums []int) []int {
	result := []int{}
	for _, num := range nums {
		if num%3 != 0 {
			result = append(result, num)
		}
	}
	return result
}

// Concrete strategy 4
type RemoveMultiplesOfFiveStrategy struct{}

func (rmofs *RemoveMultiplesOfFiveStrategy) Filter(nums []int) []int {
	result := []int{}
	for _, num := range nums {
		if num%5 != 0 {
			result = append(result, num)
		}
	}
	return result
}

// Context
type NumberFilter struct {
	strategy FilterStrategy
}

func (nf *NumberFilter) SetStrategy(s FilterStrategy) {
	nf.strategy = s
}

func (nf *NumberFilter) ExecuteStrategy(nums []int) []int {
	return nf.strategy.Filter(nums)
}

func main() {
	nums := []int{-2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	nf := &NumberFilter{}

	nf.SetStrategy(&RemoveNegativeStrategy{})
	fmt.Println(nf.ExecuteStrategy(nums))

	nf.SetStrategy(&RemovePositiveStrategy{})
	fmt.Println(nf.ExecuteStrategy(nums))

	nf.SetStrategy(&RemoveMultiplesOfThreeStrategy{})
	fmt.Println(nf.ExecuteStrategy(nums))

	nf.SetStrategy(&RemoveMultiplesOfFiveStrategy{})
	fmt.Println(nf.ExecuteStrategy(nums))
}
