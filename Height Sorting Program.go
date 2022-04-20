package main

import "fmt"

const NMAX = 100

type mahasiswa struct {
	inisial string
	tinggi  int
}

type dataMhs struct {
	tabel [NMAX]mahasiswa
	n     int
}

func main() {
	var data dataMhs

	bacaData(&data, &data.n)
	urutData(&data, data.n)
	fmt.Println("")
	cetakData(data, data.n)
}

func bacaData(t *dataMhs, n *int) {
	var banyakmhs, tinggi int
	var inisial string

	fmt.Scanln(&banyakmhs)

	for i := 0; i < banyakmhs; i++ {
		fmt.Scanln(&inisial, &tinggi)
		t.tabel[i].inisial = inisial
		t.tabel[i].tinggi = tinggi
		*n++
	}

}

func cetakData(t dataMhs, n int) {

	for i := 0; i < n; i++ {
		fmt.Println(t.tabel[i].inisial, t.tabel[i].tinggi)
	}

}

func urutData(t *dataMhs, n int) {

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if t.tabel[i].tinggi > t.tabel[j].tinggi {
				temp := t.tabel[i]
				t.tabel[i] = t.tabel[j]
				t.tabel[j] = temp
			}
		}
	}
}
