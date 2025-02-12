// Order DS
// - OrderLevels by heaps - Buy order and sell order
// - HashMap[price] => orderlist[0]
// - OrderList for each prices

// dummyengine/cmd/main.go
package main

import (
	"dummyengine/pkg/orderbook"
	"dummyengine/pkg/pricelevel"
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	orderBook := orderbook.NewOrderBook()

// 	orderBook.AddBuyOrder(&pricelevel.Order{ID: 1, Price: 100.0, Quantity: 10})
// 	orderBook.AddSellOrder(&pricelevel.Order{ID: 2, Price: 105.0, Quantity: 5})
// 	orderBook.AddBuyOrder(&pricelevel.Order{ID: 3, Price: 95.0, Quantity: 10})
// 	orderBook.AddSellOrder(&pricelevel.Order{ID: 4, Price: 105.0, Quantity: 5})
// 	orderBook.AddBuyOrder(&pricelevel.Order{ID: 5, Price: 93.0, Quantity: 10})
// 	orderBook.AddSellOrder(&pricelevel.Order{ID: 6, Price: 952.0, Quantity: 5})
// 	orderBook.AddBuyOrder(&pricelevel.Order{ID: 7, Price: 22.0, Quantity: 10})
// 	orderBook.AddSellOrder(&pricelevel.Order{ID: 8, Price: 105.0, Quantity: 5})
// 	orderBook.AddBuyOrder(&pricelevel.Order{ID: 9, Price: 223.0, Quantity: 10})
// 	orderBook.AddSellOrder(&pricelevel.Order{ID: 10, Price: 23.0, Quantity: 5})

// 	orderBook.AddBuyOrder(&pricelevel.Order{ID: 11, Price: 100.0, Quantity: 10})
// 	orderBook.AddSellOrder(&pricelevel.Order{ID: 12, Price: 105.0, Quantity: 5})
// 	orderBook.AddBuyOrder(&pricelevel.Order{ID: 13, Price: 95.0, Quantity: 10})
// 	orderBook.AddSellOrder(&pricelevel.Order{ID: 14, Price: 105.0, Quantity: 5})
// 	orderBook.AddBuyOrder(&pricelevel.Order{ID: 15, Price: 93.0, Quantity: 10})
// 	orderBook.AddSellOrder(&pricelevel.Order{ID: 16, Price: 952.0, Quantity: 5})
// 	orderBook.AddBuyOrder(&pricelevel.Order{ID: 17, Price: 22.0, Quantity: 10})
// 	orderBook.AddSellOrder(&pricelevel.Order{ID: 18, Price: 105.0, Quantity: 5})
// 	orderBook.AddBuyOrder(&pricelevel.Order{ID: 19, Price: 223.0, Quantity: 10})
// 	orderBook.AddSellOrder(&pricelevel.Order{ID: 20, Price: 23.0, Quantity: 5})

// 	fmt.Println("Top Buy Order:", orderBook.GetTopBuyOrder())
// 	fmt.Println("Top Sell Order:", orderBook.GetTopSellOrder())

// 	fmt.Println("All Buy Orders:")
// 	orderBook.GetAllBuyOrders()
// 	fmt.Println("All Sell Orders:")
// 	orderBook.GetAllSellOrders()
// }

func main() {
	orderBook := orderbook.NewOrderBook()
	rand.Seed(time.Now().UnixNano())

	orderID := 1
	basePrice := 100
	priceFluctuation := 20

	for i := 0; i < 100000; i++ {
		buyPrice := basePrice - priceFluctuation + rand.Intn(priceFluctuation)
		sellPrice := basePrice + rand.Intn(priceFluctuation)
		buyQuantity := rand.Intn(20) + 1  // Random quantity between 1 and 50
		sellQuantity := rand.Intn(20) + 1 // Random quantity between 1 and 50

		orderBook.AddBuyOrder(&pricelevel.Order{ID: orderID, Price: float64(buyPrice), Quantity: buyQuantity})
		orderID++
		orderBook.AddSellOrder(&pricelevel.Order{ID: orderID, Price: float64(sellPrice), Quantity: sellQuantity})
		orderID++
	}

	fmt.Println("Top Buy Order:", orderBook.GetTopBuyOrder())
	fmt.Println("Top Sell Order:", orderBook.GetTopSellOrder())

	fmt.Println("All Buy Orders:")
	orderBook.GetAllBuyOrders()
	fmt.Println("All Sell Orders:")
	orderBook.GetAllSellOrders()
}

// func main() {
// 	// Create an order
// 	order := hello.NewOrder(1, 100.50)

// 	// Print order details
// 	order.PrintOrder()
// }
