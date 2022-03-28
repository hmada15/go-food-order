package helper

import "fmt"

func StrToInt(str string) int32 {
	var integer int32 = 0
	fmt.Sscan(str, &integer)
	return integer
}
