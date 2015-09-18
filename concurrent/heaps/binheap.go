package binheap

import (
	"errors"
)

type BinHeap struct {
	value     float64
	degree    int
	root      *BinHeap
	childrens []*BinHeap
}

func MakeBinHeap() *BinHeap {
	bh := new(BinHeap)
	bh.root = nil
	bh.childrens = []*BinHeap{}
	return bh
}

func (bh *BinHeap) Insert(value float64) {
	newheap := &BinHeap{root: nil, childrens: []*BinHeap{}, value: value, degree: 0}
	oldroot := bh.root
	if oldroot == nil {
		bh.root = newheap
		return
	}

	bh.root = newheap
	bh.root.root = oldroot

}

func (bh *BinHeap) ExtractMin() (float64, error) {
	if bh.root == nil {
		return 0, errors.New("Binomial heap is empty")
	}

	minvalue := bh.value
	heap := bh.root
	for heap != nil {
		item := heap.value
		if item < minvalue {
			minvalue = item
		}
		heap = heap.root
	}

	return minvalue, nil
}

func (bh *BinHeap) Merge(heap1, heap2 *BinHeap) {

}

func (bh *BinHeap) mergeTree(heap1 *BinHeap) {
	if bh.value < heap1.value {
		bh.addToChildren(heap1)
	} else {
		heap1.addToChildren(bh)
		bh = heap1
	}
}

func (bh *BinHeap) addSubTree(bh1 *BinHeap) {

}

func (bh *BinHeap) addToChildren(bh1 *BinHeap) {
	bh.childrens = append(bh.childrens, bh1)
	bh.degree += bh1.degree
}
