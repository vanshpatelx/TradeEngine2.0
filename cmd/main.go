// Order DS
// - OrderLevels by heaps - Buy order and sell order
// - HashMap[price] => orderlist[0]
// - OrderList for each prices


// dummyengine/cmd/main.go
package main

import (
	"dummyengine/pkg/orderbook"
	"fmt"
)


func main() {
	orderBook := orderbook.NewOrderBook()

	orderBook.AddBuyOrder(100.0, 10)
	orderBook.AddSellOrder(105.0, 5)

	orderBook.AddBuyOrder(95.0, 10)
	orderBook.AddSellOrder(105.0, 5)
	orderBook.AddBuyOrder(93.0, 10)
	orderBook.AddSellOrder(952.0, 5)
	orderBook.AddBuyOrder(22.0, 10)
	orderBook.AddSellOrder(105.0, 5)
	orderBook.AddBuyOrder(223.0, 10)
	orderBook.AddSellOrder(23.0, 5)

	fmt.Println("Top Buy Order:", orderBook.GetTopBuyOrder())
	fmt.Println("Top Sell Order:", orderBook.GetTopSellOrder())
}



// func main() {
// 	// Create an order
// 	order := hello.NewOrder(1, 100.50)

// 	// Print order details
// 	order.PrintOrder()
// }
