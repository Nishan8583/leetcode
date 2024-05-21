/*
Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.


Our goal of serialization
1. Just do dfs
2. Append urself, then left subtree, then right subtree
	1
  /    \
2		3
	   /   \
	4		5

This will be 1,2,nil,nil,3,4,,nil,nil,5,nil,nil
1 then 2, and both subchild is empty, 3, its left is 4, then 4s child is emtpy and so on

For deserialize
1. Do dfs
2. recursive
3. If current is emtpy, return nil, but increase counter
4. If value, append, then recursively do same for left
5. THen for right
6. Return the final node create
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	serializedValue []string
}

func Constructor() Codec {
	return Codec{
		serializedValue: []string{},
	}
}

func dfsSerialize(node *TreeNode, value *[]string) {
	if node == nil {
		*value = append(*value, "nil")
		return
	}
	v := fmt.Sprintf("%d", node.Val)
	*value = append(*value, v)
	dfsSerialize(node.Left, value)
	dfsSerialize(node.Right, value)
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	dfsSerialize(root, &this.serializedValue)
	v := strings.Join(this.serializedValue, ",")
	return v
}

func dfsDeserialize(i *int, values []string) *TreeNode {
	if values[*i] == "nil" {
		*i++
		return nil
	}
	v, _ := strconv.Atoi(values[*i])
	node := TreeNode{Val: v}
	*i++
	node.Left = dfsDeserialize(i, values)
	node.Right = dfsDeserialize(i, values)
	return &node
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	values := strings.Split(data, ",")
	counter := 0
	return dfsDeserialize(&counter, values)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
