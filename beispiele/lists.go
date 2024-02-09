package beispiele

import (
	"fmt"

	"github.com/tel23a-inf/aufgaben-rekursion/lists"
)

func PrintList(list []int) {
	if lists.Empty(list) {
		return
	}

	fmt.Println(list[0])
	PrintList(list[1:])

}

func AddListElements(list []int) int {
	if lists.Empty(list) {
		return 0
	}

	return AddListElements(list[1:]) + list[0]
}
