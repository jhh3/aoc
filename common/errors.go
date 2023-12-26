package common

import "fmt"

func CheckErr(e error, msg string) {
	if e != nil {
		fmt.Println(msg)
		panic(e)
	}
}
