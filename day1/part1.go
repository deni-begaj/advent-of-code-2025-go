package day1

import (
	"fmt"
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
	amount = amount % Max
	switch op {
	case Left:
		return (num - amount + Max) % Max
	case Right:
		return (num + amount) % Max
	default:
		return num
	}
}

func Day1() {
	fmt.Println("Hello from day1")

	otherInput := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	fmt.Println(otherInput)
	inputSlice := strings.Split(Input, "\n")
	var head int16 = 50
	var zeros int16 = 0
	for i := range inputSlice {
		elem := inputSlice[i]
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
			zeros += 1
		}
		fmt.Println("Head is:", head, "| Zero counts:", zeros)
	}
}
