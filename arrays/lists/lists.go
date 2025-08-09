package lists

import "fmt"

type Product struct {
	title string
	id    int32
	price float64
}

func main() {

	var productNames [3]string
	prices := [4]float64{1, 2.8, 3, 4.1}

	productNames[2] = "Pratik"

	// featuredPrices := prices[1 : 4]
	// featuredPrices := prices[: 3] // starts with 0
	featuredPrices := prices[:2] // starts with 1 ends with 4 (exclusive)
	randomSlice := featuredPrices[0:]
	fmt.Println("RandomSlice from 0 -> 1", randomSlice)
	randomSlice = featuredPrices[:4]
	fmt.Println("RandomSlice from 0 -> 3", randomSlice)

	fmt.Println(prices)
	fmt.Println(prices[1])
	fmt.Println(productNames)

	featuredPrices[1] = 999.99 // reference to the original arr -> prices
	fmt.Println(prices)
	fmt.Println("Length of prices array", len(prices))
	fmt.Println("Capacity of featuredPrices array", cap(featuredPrices))
	fmt.Println("Slice of an array", featuredPrices)

	/*
		// slice literal
		prices := []float64{1.2, 32, 9.8, 11, 12.3}
		newSicedPrice := prices[:3]
		// a := append(newSicedPrice, 1)
		// fmt.Println(prices, a)
		newPrices := append(newSicedPrice, 9.99) // returns a new
		fmt.Println(newPrices, newSicedPrice, prices)
	*/

	// 1
	hobbies := [3]string{"cooking", "gaming", "youtube"}
	fmt.Println(hobbies)

	// 2
	fmt.Println(hobbies[0], hobbies[1:])

	// 3
	slice1 := hobbies[:2]
	slice2 := []string{hobbies[0], hobbies[1]}
	fmt.Println(slice1, slice2)

	// 5
	goals := []string{"Good job", "a gaming setup"}
	newGoals := append(goals, "good physique")
	fmt.Println(goals, newGoals)

	// 7
	productList := []Product{{title: "tooth paste", id: 91, price: 22.5}, {title: "tooth brush", id: 91, price: 22.5}, {title: "pen", id: 92, price: 20}}

	updatedList := append(productList, Product{title: "Rice", id: 12, price: 100})
	fmt.Println("Product list", productList)
	fmt.Println("updated product list", updatedList)

	/*
	// unpack slice values
	prices := []float64{1.2, 32, 9.8, 11, 12.3}
	Newprices := []float64{111, 22, 8, 1, 12}

	// discountPrices := append(prices, 2, 3, 4)
	discountPrices := append(prices, Newprices...)
	fmt.Println(discountPrices)
	*/

}
