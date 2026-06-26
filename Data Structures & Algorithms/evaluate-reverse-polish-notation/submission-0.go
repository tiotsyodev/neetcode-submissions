func evalRPN(tokens []string) int {
	stack := []int{}
	for _, v := range tokens {
	
		switch v {
			case "+", "-", "*", "/":
				a := stack[len(stack) - 2]
				b := stack[len(stack) - 1]
				stack = stack[:len(stack) - 2]

				var res int
				switch v {
				case "+":
					res = a + b
				case "-":
					res = a - b
				case "/":
					res = a / b
				case "*":
					res = a * b
				}

				stack = append(stack, res)
			
			default:
				val, _ := strconv.Atoi(v)
				stack = append(stack, val)
		}
		
	}

	return stack[0]
}