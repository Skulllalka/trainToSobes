package stroki

import (
	"fmt"
	"strconv"
	"strings"
)

func Task1() {

	tempString := strings.Builder{}

	for i := 0; i < 1000; i++ {
		tempString.WriteString(strconv.Itoa(i))
	}
	fmt.Println(tempString.String())
}
