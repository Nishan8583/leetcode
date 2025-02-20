"""You are given the head of a linked list. Delete the middle node, and return the head of the modified linked list.

The middle node of a linked list of size n is the ⌊n / 2⌋th node from the start using 0-based indexing, where ⌊x⌋ denotes the largest integer less than or equal to x.

    For n = 1, 2, 3, 4, and 5, the middle nodes are 0, 1, 1, 2, and 2, respectively.



Example 1:

Input: head = [1,3,4,7,1,2,6]
Output: [1,3,4,1,2,6]
Explanation:
The above figure represents the given linked list. The indices of the nodes are written below.
Since n = 7, node 3 with value 7 is the middle node, which is marked in red.
We return the new list after removing this node.

[1,2,3,4]
total = 4
mid = 2
when counter = 0, prev = None, current  = 1
counter = 1, prev = 1, current = 2
counter = 2, prev = 2, current = 3// remove current
Solution:
1. Go through the node, traverse it
2. Keep counting
3. Then traverse again and delete it

"""


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteMiddle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        current = head
        total = 0
        while current:
            current = current.next
            total += 1
        if total == 1:
            return None
        mid = int(total / 2)
        counter = 0
        prev = None
        current = head
        while True:
            print(counter, mid)
            if counter == mid:
                prev.next = current.next
                break
            prev = current
            current = current.next
            counter += 1
        return head
