// dummyengine/pkg/customheap/sellorderbook.go

package customheap

type SellOrderBook struct {
	CommonHeap
}

func (s *SellOrderBook) Less(i, j int) bool {
	return s.CommonHeap[i].Price < s.CommonHeap[j].Price
}
