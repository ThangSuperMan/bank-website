package helper

import "fmt"

func CheckErr(err error) {
    fmt.Println("checkErr")
	if err != nil {
		panic(err)
	}
}

func PrintMessages(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}
