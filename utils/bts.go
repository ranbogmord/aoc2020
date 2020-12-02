package utils

type bstnode struct {
	Value int
	Left  *bstnode
	Right *bstnode
}

type BTS struct {
	root *bstnode
}

func (b *BTS) Reset() {
	b.root = nil
}

func (b *BTS) Insert(value int) {
	b.insertRec(b.root, value)
}

func (b *BTS) insertRec(node *bstnode, value int) *bstnode {
	if b.root == nil {
		b.root = &bstnode{
			Value: value,
		}
		return b.root
	}
	if node == nil {
		return &bstnode{
			Value: value,
		}
	}
	if value <= node.Value {
		node.Left = b.insertRec(node.Left, value)
	}
	if value > node.Value {
		node.Right = b.insertRec(node.Right, value)
	}

	return node
}

func (b *BTS) Find(value int) *bstnode {
	return b.findRec(b.root, value)
}

func (b *BTS) findRec(node *bstnode, value int) *bstnode {
	if node == nil {
		return nil
	}

	if node.Value == value {
		return node
	}

	if value < node.Value {
		return b.findRec(node.Left, value)
	} else {
		return b.findRec(node.Right, value)
	}
}
