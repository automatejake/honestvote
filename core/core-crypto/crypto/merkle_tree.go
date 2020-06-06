package crypto

import (
	"encoding/hex"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

//NewMerkleNode Takes in bytes and encodes bytes to hex
func NewMerkleNode(left *database.MerkleNode, right *database.MerkleNode, data string) *database.MerkleNode {
	node := database.MerkleNode{}

	if left == nil && right == nil {
		node.Hash = data
	} else {
		prevHashes := []byte(left.Hash + right.Hash)
		newHash := CalculateHash(prevHashes)
		node.Hash = hex.EncodeToString(newHash)
	}

	node.Left = left
	node.Right = right

	return &node
}

//NewMerkleRoot Creates a merkle tree with the given bytes
func NewMerkleRoot(data []string) *database.MerkleTree {
	var nodes []database.MerkleNode

	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	for _, d := range data {
		node := NewMerkleNode(nil, nil, d)
		nodes = append(nodes, *node)
	}

	for len(nodes) != 1 {
		var level []database.MerkleNode

		for j := 0; j < len(nodes)-1; j += 2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], "")
			level = append(level, *node)
		}

		nodes = level
	}

	tree := database.MerkleTree{&nodes[0]}

	return &tree
}

//IsIntroverse Verifies a transaction in the merkle tree
func IsIntroverse(transaction string, node *database.MerkleNode) bool {
	var arr []database.MerkleNode

	if node.Left != nil {
		arr = append(arr, *node.Left)
	}

	if node.Right != nil {
		arr = append(arr, *node.Right)
	}

	for len(arr) > 0 {
		var arr2 []database.MerkleNode

		arr2 = arr
		arr = nil

		for _, node := range arr2 {
			if node.Hash == transaction {
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
