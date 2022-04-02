package helper

import "fmt"

func StrToInt32(str string) int32 {
	var integer int32 = 0
	fmt.Sscan(str, &integer)
	return integer
}

func StrToInt(str string) int {
	var integer int = 0
	fmt.Sscan(str, &integer)
	return integer
}

func StrToBool(str string) bool {
	var boolean bool

	if str == "true" || str == "TRUE" || str == "True" {
		boolean = true
	} else if str == "false" || str == "FALSE" || str == "False" {
		boolean = false
	}

	return boolean
}
