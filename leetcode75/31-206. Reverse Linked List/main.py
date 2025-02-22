"""
206. Reverse Linked List
[1,2,3,4,5]

1
2 points to prev which is one
prev = 2
3 points to prev, which is 2
prev = 3
"""


class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution(object):
    def reverseList(self, head):
        """
        :type head: Optional[ListNode]
        :rtype: Optional[ListNode]
        [1, n2]
        [2,n3]
        [3,n4]
        [4,n5]
        [5,None]
        """
        prev = None
        current = new_head = head

        while current:
            temp = (
                current.next
            )  # hold the value temporarily becasue we have to use it in next iteration
            current.next = prev  # the next value will be the prev value
            prev = current  # make this current value is prev because the next iteration will use this as next, therefore reversing it
            if temp:  # if there is some value in temp, update new_head, we need check here because we do not want it to be None
                new_head = temp
            current = temp
        return new_head


s = Solution()
l = ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(5, None)))))
c = s.reverseList(l)
current = c
while current:
    print(current.val)
    current = current.next
