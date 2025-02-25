// dummyengine/pkg/rabbitmqQueue/consumer.go
package rabbitmqQueue

import (
	"dummyengine/pkg/exchange"
	"encoding/json"
	"log"
	"math/big"
	"time"

	"github.com/streadway/amqp"
)

// RabbitMQConsumer handles message consumption
type RabbitMQQueue struct {
	Conn     *amqp.Connection
	Ch       *amqp.Channel
	Queue    string
	URL      string
	Exchange *exchange.Exchange
}

// EventMessage represents the structure received from RabbitMQ
type EventMessage struct {
	Task          string `json:"task"`     // "Order" || "CreateEvent" || "Settlement"
	ID            string `json:"eventId"`  // EventID
	OrderID       string `json:"orderId"`  // OrderID
	OrderPrice    string `json:"price"`    // OrderID
	OrderUserID   string `json:"userId"`   // OrderID
	OrderQuantity string `json:"quantity"` // OrderID
	Type          string `json:"type"`     // "BUY" || "SELL"
}

// NewRabbitMQConsumer initializes a new consumer
func NewRabbitMQQueue(url, queue, tradeQueue, priceQueue, orderBookQueue string) *RabbitMQQueue {
	return &RabbitMQQueue{
		Queue:    queue,
		URL:      url,
		Exchange: nil, // Initialize to nil, will be set after connection
	}
}

// Connect establishes the RabbitMQ connection
func (c *RabbitMQQueue) Connect() {
	var err error

	// Attempt connection
	c.Conn, err = amqp.Dial(c.URL)
	if err != nil {
		log.Fatalf("❌ Failed to connect to RabbitMQ: %v", err)
	}

	c.Ch, err = c.Conn.Channel()
	if err != nil {
		log.Fatalf("❌ Failed to open a channel: %v", err)
	}

	log.Printf("✅ Connected to RabbitMQ: %s (Queue: %s)", c.URL, c.Queue)

	// Auto-reconnect on failure
	c.Conn.NotifyClose(make(chan *amqp.Error))
	c.reconnect()

	// Start consuming
	c.Consume()
}

// Reconnect handles RabbitMQ reconnection logic
func (c *RabbitMQQueue) reconnect() {
	for {
		log.Printf("🔄 Reconnecting to RabbitMQ...")
		time.Sleep(5 * time.Second)

		var err error
		c.Conn, err = amqp.Dial(c.URL)
		if err == nil {
			c.Ch, err = c.Conn.Channel()
			if err == nil {
				log.Println("✅ Reconnected to RabbitMQ")
				c.Exchange = exchange.NewExchange(c.Ch, "trade_queue", "price_queue", "orderBook_queue")
				c.Consume()
				return
			}
		}
	}
}

// Consume listens to the queue
func (c *RabbitMQQueue) Consume() {
	msgs, err := c.Ch.Consume(
		c.Queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("❌ Failed to register a consumer: %v", err)
	}

	log.Printf("📥 Consuming messages from queue: %s", c.Queue)

	for msg := range msgs {
		c.processMessage(msg)
	}
}

func (c *RabbitMQQueue) processMessage(msg amqp.Delivery) {
	var orderMsg EventMessage
	if err := json.Unmarshal(msg.Body, &orderMsg); err != nil {
		log.Printf("❌ Failed to parse message: %v", err)
		msg.Nack(false, false) // Reject message
		return
	}

	// Convert string values to big.Int
	ID := new(big.Int)
	if err := ID.UnmarshalText([]byte(orderMsg.ID)); err != nil {
		log.Printf("❌ Failed to convert orderMsg.ID to big.Int: %v Event: %v", err, orderMsg)
		return
	}

	orderID := new(big.Int)
	if err := orderID.UnmarshalText([]byte(orderMsg.OrderID)); err != nil {
		log.Printf("❌ Failed to convert orderMsg.OrderID to big.Int: %v Event: %v", err, orderMsg)
		return
	}

	OrderUserID := new(big.Int)
	if err := OrderUserID.UnmarshalText([]byte(orderMsg.OrderUserID)); err != nil {
		log.Printf("❌ Failed to convert orderMsg.OrderUserID to big.Int: %v Event: %v", err, orderMsg)
		return
	}

	switch orderMsg.Task {
	case "CreateEvent":
		log.Printf("📌 Creating event: %s", ID.String())
		c.Exchange.AddEvent(ID)

	case "Settlement":
		log.Printf("💰 Processing settlement for event: %s", ID.String())
		c.Exchange.Settlement(ID)

	case "Order":
		log.Printf("📦 Processing order for event: %s", orderID.String())
		switch orderMsg.Type {
		case "BUY":
			c.Exchange.AddBuyOrder(ID, orderID, &orderMsg.OrderPrice, &orderMsg.OrderQuantity, &orderMsg.Type, OrderUserID)
		case "SELL":
			c.Exchange.AddSellOrder(ID, orderID, &orderMsg.OrderPrice, &orderMsg.OrderQuantity, &orderMsg.Type, OrderUserID)
		default:
			log.Printf("⚠️ Unknown order type: %s", orderMsg.Type)
			msg.Nack(false, false)
			return
		}

	default:
		log.Printf("⚠️ Unknown task type: %s", orderMsg.Task)
		msg.Nack(false, false)
		return
	}

	msg.Ack(false) // Acknowledge message
}
