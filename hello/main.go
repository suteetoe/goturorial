package hello

import "fmt"

func main() {
	fmt.Println("Hello GO Turotial")
}

func sumOfFirst(num int) int {
	result := 0
	for n := num; n >= 1; n-- {
		result += n
	}
	return result
}

func double(n *int) {
	*n = *n * 2
}

func appendGreeting(s *string) {
	*s = fmt.Sprintf("Hi, %s", *s)
}
