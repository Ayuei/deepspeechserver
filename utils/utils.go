package utils

import (
	"fmt"
)

func Chkerror(err error) bool {
	if err != nil{
		fmt.Println("Error: ", err)
		return true
	}

	return false
}
