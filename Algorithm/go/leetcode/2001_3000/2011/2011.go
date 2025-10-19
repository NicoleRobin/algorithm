package main

func finalValueAfterOperations(operations []string) int {
	var ans int
	for _, op := range operations {
		if op == "--X" || op == "X--" {
			ans--
		}
		if op == "++X" || op == "X++" {
			ans++
		}
	}
	return ans
}

func main() {}
