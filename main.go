package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter code: ")
	in := bufio.NewReader(os.Stdin)
	keycode, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	for _, x := range strings.Split(strings.TrimSpace(keycode), " ") {
		fmt.Print(ConvertCode(x))
	}
}

var letters = "abcdefghijklmnopqrstuvwxyz"

// converts phonekeys to string
// 2 = a
// 22 = b
// 222 = c
// 3 = d
// 9999 = z
func ConvertCode(numberSet string) string {
	if len(numberSet) < 3 && len(numberSet) > 4 {
		return ""
	}
	letterStart, err := strconv.Atoi(string(numberSet[0]))
	if err != nil {
		return ""
	}
	if letterStart == 0 {
		return " "
	}
	mod := 0
	if letterStart > 7 {
		mod = 1
	}
	return string(letters[((letterStart-2)*3)+(len(numberSet)-1)+mod])
}
