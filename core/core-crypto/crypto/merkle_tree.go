package crypto

import (
	"encoding/hex"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func NodeRehash(node *database.MerkleNode) string {
	leafHashes := []byte(node.Left.Hash + node.Right.Hash)
	rehashed := CalculateHash(leafHashes)
	return hex.EncodeToString(rehashed)
}

//NewMerkleNode Takes in bytes and encodes bytes to hex
func NewMerkleNode(left *database.MerkleNode, right *database.MerkleNode, data string) *database.MerkleNode {
	node := database.MerkleNode{}

	if left == nil && right == nil {
		node.Hash = data
		node.Hierarchy = 1
	} else {
		prevHashes := []byte(left.Hash + right.Hash)
		newHash := CalculateHash(prevHashes)
		node.Hash = hex.EncodeToString(newHash)
		node.Hierarchy = left.Hierarchy + 1
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

		if len(nodes)%2 != 0 {
			nodes = append(nodes, nodes[len(nodes)-1])
		}

		for j := 0; j < len(nodes)-1; j += 2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], "")
			level = append(level, *node)
		}

		nodes = level
	}

	tree := database.MerkleTree{&nodes[0]}

	return &tree
}

//MerkleProof Verifies a transaction in the merkle tree
func MerkleProof(transaction string, root *database.MerkleNode) bool {
	var arr []database.MerkleNode

	if root.Left.Left != nil {
		arr = append(arr, *root.Left)
	}

	if root.Right.Right != nil {
		arr = append(arr, *root.Right)
	}

	for len(arr) > 0 {
		var tempArr []database.MerkleNode

		tempArr = arr
		arr = nil

		for _, node := range tempArr {
			if node.Left.Left != nil {
				arr = append(arr, *node.Left)
			} else {
				if node.Left.Hash == transaction {
					rehash := NodeRehash(&node)
					return RecursiveMerkleProof(rehash, node.Hierarchy, root)
				}
			}
			if node.Right.Right != nil {
				arr = append(arr, *node.Right)
			} else {
				if node.Right.Hash == transaction {
					rehash := NodeRehash(&node)
					return RecursiveMerkleProof(rehash, node.Hierarchy, root)
				}
			}
		}
	}

	//Tree only has a depth of 1, check two leaves
	if root.Left.Hash == transaction {
		rehash := NodeRehash(root)
		return RecursiveMerkleProof(rehash, root.Hierarchy, root)
	} else if root.Right.Hash == transaction {
		rehash := NodeRehash(root)
		return RecursiveMerkleProof(rehash, root.Hierarchy, root)
	}

	return false
}

func RecursiveMerkleProof(rehash string, hierarchy int, root *database.MerkleNode) bool {
	var arr []database.MerkleNode

	if hierarchy == root.Hierarchy {
		return rehash == root.Hash
	}

	if root.Left.Hierarchy != hierarchy && root.Right.Hierarchy != hierarchy {
		arr = append(arr, *root.Left)
		arr = append(arr, *root.Right)
	} else {
		rehash := NodeRehash(root)
		return RecursiveMerkleProof(rehash, root.Hierarchy, root)
	}

	for len(arr) > 0 {
		var tempArr []database.MerkleNode

		tempArr = arr
		arr = nil

		for _, node := range tempArr {
			if node.Left.Hierarchy != hierarchy || node.Right.Hierarchy != hierarchy {
				arr = append(arr, *node.Left)
				arr = append(arr, *node.Right)
			} else {
				if node.Left.Hash == rehash {
					rehash := NodeRehash(&node)
					return RecursiveMerkleProof(rehash, node.Hierarchy, root)
				} else if node.Right.Hash == rehash {
					rehash := NodeRehash(&node)
					return RecursiveMerkleProof(rehash, node.Hierarchy, root)
				}
			}
		}
	}

	return false
}
