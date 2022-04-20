package main

import "fmt"

func main() {
	var t [100]int
	var n, d, u int

	fmt.Scanln(&n, &d, &u)

	isiArray(&t, n)
	sorting(&t, u, d, n)
	tampilArray(t, n)
}

func isiArray(t *[100]int, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
	}
}

func tampilArray(t [100]int, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(t[i], " ")
	}
}

func sorting(t *[100]int, u, d, n int) {
	var a, temp int
	for x := d + 1; x < u; x++ {
		a = x - 1
		for y := x; y > d; y-- {
			if t[y] > t[a] {
				temp = t[a]
				t[a] = t[y]
				t[y] = temp
			}
			a--
		}
	}
}
