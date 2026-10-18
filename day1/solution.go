package day1

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const Max int16 = 100

type Dir string

const (
	Left  Dir = "l"
	Right Dir = "r"
)

func finiteOp(num int16, op Dir, amount int16) int16 {
	if amount > Max {
		amount = amount % Max
	}
	if op == Left {
		diff := num - amount
		if diff < 0 {
			return Max - int16(math.Abs(float64(diff)))
		} else {
			return diff
		}
	} else if op == Right {
		sum := num + amount
		remainderToEnd := Max - num
		if amount > remainderToEnd {
			return amount - remainderToEnd
		} else if sum == Max {
			return 0
		} else {
			return sum
		}
	} else {
		return num
	}
}

func Day1() {
	fmt.Println("Hello from day1")

	otherInput := `L60
R10
R60
L60
R250`
	fmt.Println(otherInput)
	inputSlice := strings.Split(Input, "\n")
	var head int16 = 50
	var zeroHit int16 = 0
	for i := range inputSlice {
		elem := inputSlice[i]
		fmt.Println("Elem", elem)
		if elem == "" {
			continue
		}
		dir := Dir(strings.ToLower(elem[:1]))
		count, err := strconv.ParseInt(elem[1:], 10, 16)
		if err != nil {
			fmt.Println("Unable to parse num:", err)
			continue
		}
		head = int16(finiteOp(head, dir, int16(count)))
		if head == 0 {
			zeroHit += 1
		}
		fmt.Println("Head is:", head, "| Zero counts:", zeroHit)
	}
}
