package twosum

type TwoSum interface {
	SolutionOne() []int
	SolutionTwo() []int
	SolutionThree() []int
}

type twoData struct {
	data   []int
	target int
}

func NewTwoSum(data []int, target int) TwoSum {
	return &twoData{data: data, target: target}
}

func (this *twoData) SolutionOne() []int {
	// pendekatnya adalah menambah satu per satu hingga
	// ketemu angka yang pas
	final := make([]int, 0)
	for i := range this.data {
		for j := i + 1; j < len(this.data); j++ {
			sums := this.data[i] + this.data[j]
			if sums == this.target {
				final = append(final, i, j)
			}
		}
	}

	return final
}

func (this *twoData) SolutionTwo() []int {
	// pendekatanya adalah cari angka selisih pada array dimana
	// angka selisih didapat dari target - angka iterasi.
	// jika angka selisih tidak ditemukan dalam array,
	// simpan angka iterasi dalam array.
	temp := make(map[int]int)
	for i, num := range this.data {
		key := this.target - num
		if j, ok := temp[key]; ok {
			return []int{j, i}
		}

		temp[num] = i
	}
	return []int{}
}

func (this *twoData) SolutionThree() []int {
	left, right := 0, len(this.data)-1
	for left < right {
		if this.data[left]+this.data[right] > this.target {
			right--
		} else if this.data[left]+this.data[right] < this.target {
			left++
		} else {
			return []int{left, right}
		}
	}
	return []int{left, right}
}
