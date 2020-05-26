package crypto

import (
	"bytes"
)

type MerkleTree struct {
	RootNode *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Hash  []byte
}

func NewMerkleNode(left *MerkleNode, right *MerkleNode, data []byte) *MerkleNode {
	node := MerkleNode{}

	if left == nil && right == nil {
		node.Hash = CalculateHash(data)
	} else {
		prevHashes := append(left.Hash, right.Hash...)
		node.Hash = CalculateHash(prevHashes)
	}

	node.Left = left
	node.Right = right

	return &node
}

func NewMerkleRoot(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	for _, d := range data {
		node := NewMerkleNode(nil, nil, d)
		nodes = append(nodes, *node)
	}

	for len(nodes) != 1 {
		var level []MerkleNode

		for j := 0; j < len(nodes)-1; j += 2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			level = append(level, *node)
		}

		nodes = level
	}

	tree := MerkleTree{&nodes[0]}

	return &tree
}

func VerifyTransaction(transaction []byte, root *MerkleTree) bool {
	if TraverseTransaction(transaction, root.RootNode) {
		return true
	}

	return false
}

func TraverseTransaction(transaction []byte, node *MerkleNode) bool {
	if node.Left != nil {
		return TraverseTransaction(transaction, node.Left)
	}

	if node.Right != nil {
		return TraverseTransaction(transaction, node.Right)
	}

	return bytes.Equal(node.Hash, transaction)
}
