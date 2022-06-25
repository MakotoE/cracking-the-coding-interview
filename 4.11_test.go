package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

type BinaryNodeWithSize struct {
	Item  int
	Left  *BinaryNodeWithSize
	Right *BinaryNodeWithSize
	Size  int
}

func NewBinaryNodeWithSize(item int) *BinaryNodeWithSize {
	return &BinaryNodeWithSize{
		Item: item,
		Size: 1,
	}
}

func (b *BinaryNodeWithSize) Insert(item int) {
	if item < b.Item {
		if b.Left == nil {
			b.Left = NewBinaryNodeWithSize(item)
		} else {
			b.Left.Insert(item)
		}
	} else if b.Right == nil {
		b.Right = NewBinaryNodeWithSize(item)
	} else {
		b.Right.Insert(item)
	}

	b.Size++
	return
}

// Root item will not be deleted.
func (b *BinaryNodeWithSize) Delete(item int) {
	b.deleteWithParent(item, nil)
}

func (b *BinaryNodeWithSize) deleteWithParent(item int, parent *BinaryNodeWithSize) {
	if parent == nil {
		if item < b.Item {
			if b.Left != nil {
				b.Left.deleteWithParent(item, b)
			}
		} else if b.Right != nil {
			b.Right.deleteWithParent(item, b)
		}
	} else if item == b.Item {
		if b.Left == nil && b.Right == nil {
			if parent.Left != nil && parent.Left.Item == item {
				parent.Left = nil
			} else if parent.Right != nil && parent.Right.Item == item {
				parent.Right = nil
			} else {
				panic("unexpected")
			}
		} else if b.Left != nil && b.Right == nil {
			b.Item = b.Left.Item
			b.Left = nil
		} else if b.Left == nil && b.Right != nil {
			b.Item = b.Right.Item
			b.Right = nil
		} else {
			b.Item = b.Right.minItem()
			b.Right.deleteWithParent(b.Item, b)
		}
	} else if item < b.Item {
		if b.Left != nil {
			b.Left.deleteWithParent(item, b)
		}
	} else if b.Right != nil {
		b.Right.deleteWithParent(item, b)
	}

	b.Size--
}

func (b *BinaryNodeWithSize) minItem() int {
	curr := b
	for curr.Left != nil {
		curr = b.Left
	}
	return curr.Item
}

// GetRandomNode returns a random item from the tree with uniform distribution.
func (b *BinaryNodeWithSize) GetRandomNode() int {
	n := rand.Intn(b.Size)
	if n == 0 {
		return b.Item
	}

	if b.Left != nil && n < b.Left.Size+1 {
		return b.Left.GetRandomNode()
	}

	return b.Right.GetRandomNode()
}

func (b *BinaryNodeWithSize) List() []*BinaryNodeWithSize {
	var result []*BinaryNodeWithSize

	if b.Left != nil {
		result = b.Left.List()
	}
	result = append(result, b)
	if b.Right != nil {
		result = append(result, b.Right.List()...)
	}
	return result
}

func TestBinaryNodeWithSize_Insert(t *testing.T) {
	tree := NewBinaryNodeWithSize(0)
	list := tree.List()
	assert.Len(t, list, 1)
	assert.Equal(t, 0, list[0].Item)
	assert.Equal(t, 1, list[0].Size)

	tree.Insert(1)
	list = tree.List()
	assert.Len(t, list, 2)
	assert.Equal(t, 1, list[1].Item)
	assert.Equal(t, 2, list[0].Size)
	assert.Equal(t, 1, list[1].Size)

	tree.Insert(2)
	list = tree.List()
	assert.Len(t, list, 3)
	assert.Equal(t, 2, list[2].Item)
	assert.Equal(t, 3, list[0].Size)
	assert.Equal(t, 2, list[1].Size)

	tree.Insert(-1)
	list = tree.List()
	assert.Len(t, list, 4)
	assert.Equal(t, -1, list[0].Item)
	assert.Equal(t, 4, list[1].Size)
}

func TestBinaryNodeWithSize_Delete(t *testing.T) {
	tree := NewBinaryNodeWithSize(5)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)

	tree.Delete(8)
	list := tree.List()
	assert.Len(t, list, 7)
	for _, a := range list {
		assert.NotEqual(t, 8, a.Item)
	}

	tree.Delete(6)
	list = tree.List()
	assert.Len(t, list, 6)
	for _, a := range list {
		assert.NotEqual(t, 6, a.Item)
	}

	tree.Delete(2)
	list = tree.List()
	assert.Len(t, list, 5)
	for _, a := range list {
		assert.NotEqual(t, 2, a.Item)
	}
	assert.Equal(t, 5, list[3].Item)
	assert.Equal(t, 5, list[3].Size)
	assert.Equal(t, 3, list[1].Item)
	assert.Equal(t, 3, list[1].Size)
}

func TestBinaryNodeWithSize_GetRandomNode(t *testing.T) {
	tree := NewBinaryNodeWithSize(5)
	assert.Equal(t, 5, tree.GetRandomNode())

	tree.Insert(6)
	tree.GetRandomNode()

	tree.Insert(7)
	tree.Insert(8)
	tree.GetRandomNode()

	tree.Insert(4)
	tree.GetRandomNode()
}
