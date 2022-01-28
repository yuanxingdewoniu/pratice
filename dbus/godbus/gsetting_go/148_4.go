package main

func main() {
	Map := make(map[int]int)

	for i := 0; i < 10; i++ {
		go writeMap(Map, i, i)
		go readMap(Map, i)
	}

}

func readMap(Map map[int]int, key int) int {
	return Map[key]
}

func writeMap(Map map[int]int, key int, value int) {
	Map[key] = value
}
