package utils

import (
	"fmt"
	"os"
	"strconv"
)

func StrToInt(src string) int {
	convInt, err := strconv.ParseInt(src, 10, 64)
	if err != nil {
		fmt.Println("Failed to convert ", src, " to an int ", err.Error())
		os.Exit(1)
	}

	return int(convInt)
}
