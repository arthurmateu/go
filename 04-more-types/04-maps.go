package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex // Like struct literals but the keys are required

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])
}

/* Map exercise
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)

	for _, w := range W {
		m[w] += 1
	}

	return m
}
*/
