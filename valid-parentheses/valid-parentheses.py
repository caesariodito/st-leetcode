class Solution:
    def isValid(self, s: str) -> bool:
        parantheses = {'(': ')', '{': '}', '[': ']'}
        stack = []
        for char in s:
            if char in parantheses.keys():
                stack.append(char)
            else:
                if len(stack) != 0 and parantheses[stack[-1]] == char:
                    stack.pop()
                else:
                    return False
        if len(stack) == 0:
            return True
        else:
            return False
