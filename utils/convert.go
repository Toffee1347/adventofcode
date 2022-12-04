package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StrToInt(src string) int {
	convInt, err := strconv.ParseInt(src, 10, 64)
	if err != nil {
		fmt.Println("Failed to convert ", src, " to an int ", err.Error())
		os.Exit(1)
	}

	return int(convInt)
}

func StrToUint(src string) uint {
	convInt, err := strconv.ParseUint(src, 10, 64)
	if err != nil {
		fmt.Println("Failed to convert ", src, " to an int ", err.Error())
		os.Exit(1)
	}

	return uint(convInt)
}

func SplitStrToInt(data string, sep string) (newData []int) {
	splitData := strings.Split(data, sep)

	for _, singleData := range splitData {
		newData = append(newData, StrToInt(singleData))
	}

	return
}

func IntToStr(src int) string {
	return strconv.FormatInt(int64(src), 10)
}
