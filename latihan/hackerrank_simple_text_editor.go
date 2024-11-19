package latihan

import (
	"fmt"
	"strconv"
)

type myText struct {
	txt  string
	undo []string
}

func NewMyText() *myText {
	return &myText{
		txt:  "",
		undo: make([]string, 0),
	}
}

func (m *myText) Append(str string) {
	m.txt += str
	m.undo = append(m.undo, str)
}

func (m *myText) Delete(k int) {
	// m.undo = append(m.undo, m.txt[k:])
	m.txt = m.txt[k:]
}

func (m *myText) Print(k int) {
	fmt.Printf("%v\n", string(m.txt[k-1]))
}

func (m *myText) Undo() {
	// fmt.Printf("before: m.undo: %v, m.txt: %v\n", m.undo, m.txt)

	n := len(m.undo) - 1
	m.txt = m.txt[:len(m.txt)-len(m.undo[n])]
	m.undo = m.undo[:n]
	if len(m.undo) > 0 && len(m.txt) == 0 {
		m.txt = m.undo[n-1]
	}
	// fmt.Printf("after: m.undo: %v, m.txt: %v\n", m.undo, m.txt)
}

func MyTextEditor(instructions [][]string) {
	txt := NewMyText()
	for _, instruction := range instructions {
		switch instruction[0] {
		case "1":
			txt.Append(instruction[1])

		case "2":
			k, _ := strconv.Atoi(instruction[1])
			txt.Delete(k)

		case "3":
			k, _ := strconv.Atoi(instruction[1])
			txt.Print(k)

		case "4":
			txt.Undo()
		}
	}
}
