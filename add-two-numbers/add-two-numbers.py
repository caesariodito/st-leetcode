# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        l1head = l1
        l2head = l2

        l1number = ""
        l2number = ""

        while l1head != None:
            # print(l1head.val)
            l1number += f"{l1head.val}"
            l1head = l1head.next

        while l2head != None:
            # print(l2head.val)
            l2number += f"{l2head.val}"
            l2head = l2head.next

        l1number = int(l1number[::-1])
        l2number = int(l2number[::-1])
        
        result = l1number + l2number

        result = str(result)[::-1]

        # print(result)

        rhead = ListNode(result[0])
        current_node = rhead
        for char in result[1:]:
            current_node.next = ListNode(int(char))
            current_node = current_node.next

        return rhead


