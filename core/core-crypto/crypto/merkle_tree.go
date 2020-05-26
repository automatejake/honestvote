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
		node.Hash = data
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

func TraverseTransaction(transaction []byte, root *MerkleTree) bool {
	return IsIntroverse(transaction, root.RootNode)
}

// func IsIntroverse(transaction []byte, node *MerkleNode) bool {
// 	if node.Left == nil && node.Right == nil {
// 		return bytes.Equal(node.Hash, transaction)
// 	}

// 	var l = false
// 	var r = false

// 	if node.Left != nil {
// 		l = IsIntreverse(transaction, node.Left)
// 	}

// 	if node.Right != nil {
// 		r = IsIntreverse(transaction, node.Right)
// 	}

// 	return r || l
// }

func IsIntroverse(transaction []byte, node *MerkleNode) bool {
	var arr []MerkleNode

	if node.Left != nil {
		arr = append(arr, *node.Left)
	}

	if node.Right != nil {
		arr = append(arr, *node.Right)
	}

	for len(arr) > 0 {
		var arr2 []MerkleNode

		arr2 = arr
		arr = nil

		for _, node := range arr2 {
			if bytes.Equal(node.Hash, transaction) {
				return true
			}

			if node.Left != nil {
				arr = append(arr, *node.Left)
			}

			if node.Right != nil {
				arr = append(arr, *node.Right)
			}
		}
	}

	return false
}
