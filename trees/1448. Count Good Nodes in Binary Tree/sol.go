/*
https://leetcode.com/problems/count-good-nodes-in-binary-tree/

Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.

Return the number of good nodes in the binary tree.

Example 1:

Input: root = [3,1,4,3,null,1,5]
Output: 4
Explanation: Nodes in blue are good.
Root Node (3) is always a good node.
Node 4 -> (3,4) is the maximum value in the path starting from the root.
Node 5 -> (3,4,5) is the maximum value in the path
Node 3 -> (3,1,3) is the maximum value in the path.

Example 2:

Input: root = [3,3,null,4,2]
Output: 3
Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.

Example 3:

Input: root = [1]
Output: 1
Explanation: Root is considered as good.

Constraints:

	The number of nodes in the binary tree is in the range [1, 10^5].
	Each node's value is between [-10^4, 10^4].
*/

/*
Solution:
Traverse, pass in max value as arguement as well
if current node is >= max value till current, its a good node, increse counter update the max value
else not a good value, do not update counter and max value, keep going
*/
package main

import (
	"fmt"
	"math"
)

// “ Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	counter := 0
	fmt.Println("Starting with counter", counter)
	dfs(root, &counter, math.MinInt64)
	return counter
}

func dfs(node *TreeNode, counter *int, maxValueInPath int) {

	// if node is nil, return
	if node == nil {
		return
	}

	// first case, if current value is gte or equals to the max value in parent
	// then its a good node, and current value is the max value till now, so increase the counter
	cur_max := node.Val
	if node.Val >= maxValueInPath {
		*counter++
		fmt.Printf("Increase for node value %d because parent vlaue is %d counter =%d\n", node.Val, maxValueInPath, *counter)
	} else {
		// second case, if current value is less then the max in parent
		// not a good node, current max is still the max value from parent
		cur_max = maxValueInPath
	}

	dfs(node.Left, counter, cur_max)
	dfs(node.Right, counter, cur_max)
}
