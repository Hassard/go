package main

import "fmt"

func main() {
	// Page 52
	fmt.Println("Page 52")
	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}

	fmt.Println(teams["Orcas"])
	fmt.Println(len(teams))

	// Page 53
	fmt.Println("Page 53")
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])

	// Page 54
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)

	fmt.Println(m)
	delete(m, "hello")
	fmt.Println(m)

	// Page 55
	inSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10, 100}
	for _, v := range vals {
		inSet[v] = true
	}
	fmt.Println(len(vals), len(inSet))
	fmt.Println(inSet[5])
	fmt.Println(inSet[500])
	if inSet[100] {
		fmt.Println("100 is in the set")
	}

}
