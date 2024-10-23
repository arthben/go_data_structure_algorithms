package latihan

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestMinLength(t *testing.T) {
	s := "ABFCACDB"
	fmt.Printf("Input: %v\n", s)
	fmt.Printf("minLength: %v\n", minLength(s))
}

func TestAddMakeValid(t *testing.T) {
	// s := "())"
	// s := "((("
	s := "()))(("
	fmt.Printf("Input: %v\n", s)
	fmt.Printf("minAddToMakeValid: %v\n", minAddToMakeValid(s))
}

func TestMaxWidht(t *testing.T) {
	// nums := []int{6, 0, 8, 2, 1, 5}
	// nums := []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}
	nums := []int{2, 2, 1}
	fmt.Printf("Input: %v\n", nums)
	fmt.Printf("maxWidhtRamp: %v\n", maxWidthRamp(nums))
}

func TestSmallestChair(t *testing.T) {
	times := [][]int{
		{3, 10}, {1, 5}, {2, 6},
	}
	targetFriend := 1
	fmt.Printf("Input times: %v\n", times)
	fmt.Printf("Input target: %v\n", targetFriend)
	fmt.Printf("smallestChair: %v\n", smallestChair(times, targetFriend))
}

func TestCipher(t *testing.T) {
	// s := "middle-Outz"
	// k := 2
	// s := "Always-Look-on-the-Bright-Side-of-Life"
	// k := 5
	// output := "Fqbfdx-Qttp-ts-ymj-Gwnlmy-Xnij-tk-Qnkj"

	s := "Hello_World!"
	k := 4
	output := "Lipps_Asvph!"

	// s := "www.abc.xy"
	// k := 87
	// output := "fff.jkl.gh"

	final := ""
	for _, v := range s {
		switch {
		case v >= 'a' && v <= 'z':
			i := 'a' + ((v - 'a' + rune(k)) % 26)
			final += string(rune(i))

		case v >= 'A' && v <= 'Z':
			i := 'A' + ((v - 'A' + rune(k)) % 26)
			final += string(rune(i))
		default:
			final += string(v)
		}
	}
	fmt.Printf("Input: %v\n", s)
	fmt.Printf("Output: %v\n", output)
	fmt.Printf("final : %v\n", final)

	// for i := 'A'; i <= 'z'; i++ {
	// 	fmt.Printf("i: %v, char: %v\n", i, string(i))
	// }

}

func sortString(s string) string {
	t := strings.Split(s, "")
	sort.Strings(t)
	return strings.Join(t, "")
}

func TestGrid(t *testing.T) {
	// grid := []string{"ebacd", "fghij", "olmkn", "trpqs", "xywuv"}
	grid := []string{"abc", "lmp", "qrt"}
	// grid := []string{"mpxz", "abcd", "wlmf"}
	var newGrid []string

	for g := range grid {
		newGrid = append(newGrid, sortString(grid[g]))
	}

	for i := range len(newGrid[0]) {
		concatCol := ""
		for j := 0; j < len(newGrid); j++ {
			c := newGrid[j][i]
			concatCol += string(c)
			// fmt.Printf("i: %v, j: %v, c: %v\n", i, j, string(c))
		}
		if concatCol != sortString(concatCol) {
			fmt.Println("Tidak Cocok")
			return
		}
	}

	fmt.Println("Cocok")
}

func TestSuperDigit(t *testing.T) {
	k := int32(4)
	n := "9875"
	// k := int32(3)
	// n := "123"
	fmt.Printf("superDigit: %v\n", superDigit(n, k))
}

func superDigit(n string, k int32) int32 {
	sum := 0
	for i := 0; i < len(n); i++ {
		strNum := string(n[i])
		num, _ := strconv.Atoi(strNum)
		sum += (num * int(k))
	}

	if sum < 10 {
		return int32(sum)
	}

	strN := strconv.Itoa(sum)
	return superDigit(strN, 1)
}
