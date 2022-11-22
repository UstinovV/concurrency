package printinorder

import (
	"fmt"
	"testing"
)

func TestPrintInOrder(t *testing.T) {
	var tests = [][3]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("{ %d, %d, %d }", tt[0], tt[1], tt[2])
		t.Run(testname, func(t *testing.T) {
			PrintInOrder(tt)
		})
	}
}
