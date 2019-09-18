package main

import "fmt"

func main() {
	tasks := []byte{'A', 'A', 'A', 'D', 'E', 'A'}
	fmt.Println(leastInterval(tasks, 1))
}
