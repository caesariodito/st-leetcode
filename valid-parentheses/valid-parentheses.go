// convertion from gpt
func isValid(s string) bool {
    parentheses := map[rune]rune{'(': ')', '{': '}', '[': ']'}
    stack := make([]rune, 0)
    for _, char := range s {
        if _, ok := parentheses[char]; ok {
            stack = append(stack, char)
        } else if len(stack) == 0 || parentheses[stack[len(stack)-1]] != char {
            return false
        } else {
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}