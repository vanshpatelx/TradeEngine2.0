package exchange

import (
	"dummyengine/pkg/orderbook"
	"log"
	"math/big"
	"github.com/streadway/amqp"
)

type Exchange struct {
	OrderBooks     map[string]*orderbook.OrderBook
	TradeQueue     string
	PriceQueue     string
	OrderBookQueue string
	Ch             *amqp.Channel
}

func NewExchange(ch *amqp.Channel, tradeQueue, priceQueue, orderBookQueue string) *Exchange {
	return &Exchange{
		OrderBooks:     make(map[string]*orderbook.OrderBook),
		TradeQueue:     tradeQueue,
		PriceQueue:     priceQueue,
		OrderBookQueue: orderBookQueue,
		Ch:             ch,
	}
}

func (e *Exchange)AddEvent(eventID *big.Int) {
	eventKey := eventID.String()
	if _, exists := e.OrderBooks[eventKey]; exists{
		log.Printf("Event %s already exists", eventKey)
		return
	}

	// create orderbook
	e.OrderBooks[eventKey] = orderbook.NewOrderBook(e.Ch, eventID, e.TradeQueue, e.PriceQueue, e.OrderBookQueue)
	log.Printf("New OrderBook created for event %s", eventKey)
}

func (e *Exchange)AddBuyOrder(eventID *big.Int, orderID *big.Int, orderPrice int, orderQuantity int, orderUserID *big.Int){
	eventKey := eventID.String()
	orderBook, exists := e.OrderBooks[eventKey]
	if !exists {
		log.Printf("OrderBook for event %s not found", eventKey)
		return
	}

	orderBook.AddBuyOrder(orderID, orderPrice, orderQuantity, orderUserID)
	log.Printf("📦 Added buy order %s to event %s", orderID.String(), eventKey)

}
func (e *Exchange)AddSellOrder(eventID *big.Int, orderID *big.Int, orderPrice int, orderQuantity int, orderUserID *big.Int){
	eventKey := eventID.String()
	orderBook, exists := e.OrderBooks[eventKey]
	if !exists {
		log.Printf("OrderBook for event %s not found", eventKey)
		return
	}

	orderBook.AddSellOrder(orderID, orderPrice, orderQuantity, orderUserID)
	log.Printf("📦 Added sell order %s to event %s", orderID.String(), eventKey)
}

func (e *Exchange)Settlement(eventID *big.Int){
	
}
