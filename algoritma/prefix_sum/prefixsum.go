package prefixsum

type PrefixSum interface {
	Proses() []int
}

type dataPrefix struct {
	data []int
}

func NewPrefixSum(data []int) PrefixSum {
	return &dataPrefix{data: data}
}

func (this *dataPrefix) Proses() []int {
	// contoh :
	// data = [1, 2, 3, 4, 5, 6]
	// final = [1, 3, 6, 10, 15, 21]

	final := make([]int, len(this.data))
	for x := 0; x < len(this.data); x++ {
		if x == 0 {
			final[x] += this.data[x]
			continue
		}
		final[x] = final[x-1] + this.data[x]
	}
	return final
}
