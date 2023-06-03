class Solution:
    def isValid(self, s: str) -> bool:
        parantheses = {'(': ')', '{': '}', '[': ']'}
        stack = []
        for char in s:
            if char in parantheses:
                stack.append(char)
            elif len(stack) == 0 or parantheses[stack.pop()] != char:
                return False
        return len(stack) == 0
