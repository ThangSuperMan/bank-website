package helper

import "fmt"

func CheckErr(err error) {
    fmt.Println("checkErr")
	if err != nil {
		panic(err)
	}
}
