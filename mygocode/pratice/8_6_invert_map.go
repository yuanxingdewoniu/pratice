package main

import (
	"fmt"
	"sort"
)
var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		invMap[v] = k
		fmt.Printf("Key: %v, Value: %v \n", k, v)// the telephone alphabet:
		package main
		import (
			"fmt"
		"sort"
		)

		var (
			barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
				"delta": 87, "echo": 56, "foxtrot": 12,
				"golf": 34, "hotel": 16, "indio": 87,
				"juliet": 65, "kili": 43, "lima": 98}
		)

		func main() {
			fmt.Println("unsorted:")
			for k, v := range barVal {
				fmt.Printf("Key: %v, Value: %v / ", k, v)
			}
			keys := make([]string, len(barVal))
			i := 0
			for k, _ := range barVal {
				keys[i] = k
				i++
			}
			sort.Strings(keys)
			fmt.Println()
			fmt.Println("sorted:")
			for _, k := range keys {
				fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
			}
		}
	}
	fmt.Println("inverted:")
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v \n", k, v)
	}
}