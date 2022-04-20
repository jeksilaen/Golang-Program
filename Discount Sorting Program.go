package main

import "fmt"

type belanja [100]int

func main() {
	var t belanja
	var n, harga, banyak, total int
	var diskon float64

	// Proses Input Data
	fmt.Scanln(&harga, &banyak)
	for harga != 0 || banyak != 0 {
		t[n] = harga * banyak
		total = total + harga*banyak
		n++
		fmt.Scanln(&harga, &banyak)
	}

	if total > 100000 {
		urut(&t, n)
		promo(t, n, &diskon)
		fmt.Println(total, diskon)
	} else {
		fmt.Println(total, diskon)
	}
}

func urut(t *belanja, n int) {
	var temp int
	for x := 0; x < n-1; x++ {
		for y := x + 1; y < n; y++ {
			if t[y] < t[x] {
				temp = t[x]
				t[x] = t[y]
				t[y] = temp
			}
		}
	}
}

func promo(t belanja, n int, diskon *float64) {
	for i := 0; i < n; i++ {
		*diskon += float64(t[i] - t[i]*(i+1)/100)
	}
}
