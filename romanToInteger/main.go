package main

import (
	"fmt"
	"strings"
)

type IntWithRoman struct {
	arabian int
	roman   string
}

func intToRoman(num int) string {
	thousands := []string{"", "M", "MM", "MMM"}
	hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	builder := strings.Builder{}

	builder.WriteString(thousands[num/1000])
	builder.WriteString(hundreds[num%1000/100])
	builder.WriteString(tens[num%1000%100/10])
	builder.WriteString(ones[num%1000%100%10/1])

	return builder.String()
}

func main() {
	fmt.Print(intToRoman(90))
}
