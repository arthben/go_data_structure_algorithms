package permutasi

import (
	"fmt"
	"strings"
)

/*
Permutasi pada intinya membagi menjadi 2 bagian yaitu kanan dan kiri
Jika elemen sudah ada disebelah kiri, maka elemen tersebut
tidak bisa berada disebelah kanan.

contoh : ABC

Kiri   Kanan
A       BC
B       AC
C       AB
A       CB
B       CA
C       BA
*/
type Permutasi interface {
	Proses() []string
}

type param struct {
	kata string
}

func NewPermutasi(kata string) Permutasi {
	return &param{kata: kata}
}

func faktorial(n int) int {
	total := 1
	for i := 1; i <= n; i++ {
		total *= i
	}
	fmt.Printf("Faktorial dari %v (%v!) adalah %v\n\n", n, n, total)
	return total
}

func (p *param) Proses() []string {
	final := make([]string, 0)

	arr := strings.Split(p.kata, "")

	x := 0
	m := len(p.kata)
	for range faktorial(m) {
		if x >= m-1 {
			x = 0
		}

		final = append(final, p.kata)
		arr[x], arr[x+1] = arr[x+1], arr[x]
		p.kata = strings.Join(arr, "")
		x++
	}

	return final
}
