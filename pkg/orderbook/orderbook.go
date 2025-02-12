// dummyengine/pkg/orderbook/orderbook.go
package orderbook

import (
	"container/heap"
	"dummyengine/pkg/customheap"
	"dummyengine/pkg/pricelevel"
	"fmt"
)

type OrderBook struct {
	BuyOrders  *customheap.BuyOrderBook
	SellOrders *customheap.SellOrderBook
}

func NewOrderBook() *OrderBook {
	buyHeap := &customheap.BuyOrderBook{}
	sellHeap := &customheap.SellOrderBook{}
	heap.Init(buyHeap)
	heap.Init(sellHeap)

	return &OrderBook{
		BuyOrders:  buyHeap,
		SellOrders: sellHeap,
	}
}

func (ob *OrderBook) AddBuyOrder(order *pricelevel.Order) {
	for _, level := range ob.BuyOrders.CommonHeap {
		if level.Price == order.Price {
			level.Orders = append(level.Orders, order)
			level.Quantity += order.Quantity
			heap.Fix(ob.BuyOrders, level.Index)
			return
		}
	}

	newLevel := &pricelevel.PriceLevel{
		Price:    order.Price,
		Quantity: order.Quantity,
		Orders:   []*pricelevel.Order{order},
	}

	heap.Push(ob.BuyOrders, newLevel)
}

func (ob *OrderBook) AddSellOrder(order *pricelevel.Order) {
	for _, level := range ob.SellOrders.CommonHeap {
		if level.Price == order.Price {
			level.Orders = append(level.Orders, order)
			level.Quantity += order.Quantity
			heap.Fix(ob.BuyOrders, level.Index)
			return
		}
	}

	newLevel := &pricelevel.PriceLevel{
		Price:    order.Price,
		Quantity: order.Quantity,
		Orders:   []*pricelevel.Order{order},
	}

	heap.Push(ob.SellOrders, newLevel)
}

func (ob *OrderBook) GetTopBuyOrder() *pricelevel.PriceLevel {
	if len(ob.BuyOrders.CommonHeap) > 0 {
		return ob.BuyOrders.CommonHeap[0]
	}
	return nil
}

func (ob *OrderBook) GetTopSellOrder() *pricelevel.PriceLevel {
	if len(ob.SellOrders.CommonHeap) > 0 {
		return ob.SellOrders.CommonHeap[0]
	}
	return nil
}

func (ob *OrderBook) GetAllBuyOrders() {
	var totalOrders int = 0
	if len(ob.BuyOrders.CommonHeap) > 0 {
		for i := range ob.BuyOrders.CommonHeap {
			orders := ob.BuyOrders.CommonHeap[i].Orders
			fmt.Printf("Buy Order: Total Q: %v  Price: %v \n", ob.BuyOrders.CommonHeap[i].Quantity, ob.BuyOrders.CommonHeap[i].Price)
			fmt.Printf("Numbers of Unique Orders: %v\n", len(orders))

			totalOrders += len(orders)
			// for _, order := range orders {
			// 	fmt.Printf("%v %v\n", order.Price, order.Quantity)
			// }
		}
	}
	fmt.Printf("Total Orders: %v\n", totalOrders)
}

func (ob *OrderBook) GetAllSellOrders() {
	var totalOrders int = 0
	if len(ob.SellOrders.CommonHeap) > 0 {
		for i := range ob.SellOrders.CommonHeap {
			orders := ob.SellOrders.CommonHeap[i].Orders
			fmt.Printf("Sell Order: Total Q: %v  Price: %v \n", ob.SellOrders.CommonHeap[i].Quantity, ob.SellOrders.CommonHeap[i].Price)
			fmt.Printf("Numbers of Unique Orders: %v\n", len(orders))

			totalOrders += len(orders)
			// for id, order := range orders {
			// 	fmt.Printf("%v %v %v\n", id, order.Price, order.Quantity)
			// }
		}
	}
	fmt.Printf("Total Orders: %v\n", totalOrders)
}
