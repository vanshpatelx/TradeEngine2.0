// dummyengine/pkg/orderbook/orderbook.go
package orderbook

import (
	"container/heap"
	"dummyengine/pkg/customheap"
	"dummyengine/pkg/pricelevel"
)

type OrderBook struct{
	BuyOrders customheap.BuyOrderBook
	SellOrders customheap.SellOrderBook
}

func NewOrderBook() *OrderBook{
	buyHeap := &customheap.BuyOrderBook{}
	sellHeap := &customheap.SellOrderBook{}
	heap.Init(buyHeap)
	heap.Init(sellHeap)

	return &OrderBook{
		BuyOrders: *buyHeap,
		SellOrders: *sellHeap,
	}
}

func (ob *OrderBook) AddBuyOrder(price float64, quantity int) {
	order := &pricelevel.PriceLevel{Price: price, Quantity: quantity}
	heap.Push(&ob.BuyOrders, order)
}

func (ob *OrderBook) AddSellOrder(price float64, quantity int) {
	order := &pricelevel.PriceLevel{Price: price, Quantity: quantity}
	heap.Push(&ob.SellOrders, order)
}											

func (ob *OrderBook) GetTopBuyOrder() *pricelevel.PriceLevel {
	if len(ob.BuyOrders.CommonHeap) > 0 {
		return ob.BuyOrders.CommonHeap[0]
	}
	return nil
}

func (ob *OrderBook) GetAllBuyOrder() *pricelevel.PriceLevel {
	if len(ob.BuyOrders.CommonHeap) > 0 {
		for i in range(len(ob.BuyOrders.CommonHeap)):
			PrintF()
	}
}

func (ob *OrderBook) GetTopSellOrder() *pricelevel.PriceLevel {
	if len(ob.SellOrders.CommonHeap) > 0 {
		return ob.SellOrders.CommonHeap[0]
	}
	return nil
}
