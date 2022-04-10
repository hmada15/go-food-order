package helper

import (
	"fmt"
	"regexp"
	"strings"
)

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

	if str == "true" || str == "TRUE" || str == "True" || str == "1" {
		boolean = true
	} else if str == "false" || str == "FALSE" || str == "False" || str == "0" {
		boolean = false
	}

	return boolean
}

func StrToSlug(str string) string {
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	regexChars := regexp.MustCompile("[^a-zA-Z0-9-_]")
	regexpSign := regexp.MustCompile("-+")
	str = regexChars.ReplaceAllString(str, "-")
	str = regexpSign.ReplaceAllString(str, "-")
	str = strings.Trim(str, "-_")

	return str
}
