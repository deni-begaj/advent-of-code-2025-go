package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func countZeroRotations(num int16, zeros int16, op Dir, amount int16) int16 {
	zeros += amount / Max
	normalizedAmount := amount % Max
	if num == 0 {
		return zeros
	}
	if op == Left && num-normalizedAmount <= 0 {
		zeros += 1
	} else if op == Right && num+normalizedAmount >= Max {
		zeros += 1
	}
	return zeros
}

func Day1Part2() {
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
		zeros = countZeroRotations(head, zeros, dir, int16(count))
		head = int16(finiteOp(head, dir, int16(count)))
		fmt.Println("Head is:", head, "| Zero counts:", zeros)
	}
}
