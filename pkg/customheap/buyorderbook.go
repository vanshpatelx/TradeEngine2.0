// dummyengine/pkg/customheap/buyorderbook.go

package customheap

type BuyOrderBook struct {
	CommonHeap
}

func (b *BuyOrderBook) Less(i, j int) bool {
	return b.CommonHeap[i].Price > b.CommonHeap[j].Price
}
