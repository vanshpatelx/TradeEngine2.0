package exchange

import (
	"dummyengine/pkg/orderbook"
	"github.com/streadway/amqp"
	"math/big"
)

type Exchange struct {
	eventID        *big.Int
	OrderBooks     map[string]*orderbook.OrderBook
	TradeQueue     string
	PriceQueue     string
	OrderBookQueue string
}

func NewExchange(ch *amqp.Channel, tradeQueue, priceQueue, orderBookQueue string) *Exchange {
	return &Exchange{
		eventID:        big.NewInt(0),
		OrderBooks:     make(map[string]*orderbook.OrderBook),
		TradeQueue:     tradeQueue,
		PriceQueue:     priceQueue,
		OrderBookQueue: orderBookQueue,
	}
}

func (e *Exchange)AddEvent{

}

func (e *Exchange)Settlement{
	
}

func (e *Exchange)AddBuyOrder{
	
}
func (e *Exchange)AddSellOrder{
	
}