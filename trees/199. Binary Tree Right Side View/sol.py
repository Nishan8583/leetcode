# Definition for a binary tree node.
class TreeNode:
     def __init__(self, val=0, left=None, right=None):
         self.val = val
         self.left = left
         self.right = right

# return the values of the nodes you can see ordered from top to bottom.
# So DFS, only on th right side
# Solution:
#   - Do a recursion, only on the right side
#   - Append the values to a list

# python specific
# self.dfs(node: TreeNode)
# self.list.append(node)
# self.dfs(node.right)


class Solution:  
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        def dfs(node: TreeNode, list: List[int]):
            if not node:
                return
            
            list.append(node.val)

            dfs(node.right, list)
        
        list = []
        dfs(root, list)
        return list
        