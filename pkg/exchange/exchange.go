package exchange

import (
	"dummyengine/pkg/orderbook"
	"math/big"
	"strings"

	"github.com/streadway/amqp"
)

type Exchange struct {
	eventID *big.Int
    OrderBooks map[string]*orderbook.OrderBook
	TradeQueue     string
	PriceQueue     string
	OrderBookQueue string
}


func NewExchange(ch *amqp.Channel, tradeQueue, priceQueue, orderBookQueue string) *Exchange {

	return &Exchange{
		eventID: big.Int,
		OrderBooks: nil,
		tradeQueue: tradeQueue,
		priceQueue: priceQueue,
		orderBookQueue: orderBookQueue
	}
}
