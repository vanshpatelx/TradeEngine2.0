// dummyengine/pkg/pricelevel/pricelevel.go


package pricelevel


type Order struct {
	ID       int
	Price    float64
	Quantity int
}

type PriceLevel struct {
	Price    float64
	Quantity int
	Index    int // Heap index
	Orders   []Order
}
