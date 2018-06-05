package main

import "fmt"

func main() {
	fmt.Println(replace(""))
	fmt.Println(replace(" "))
	fmt.Println(replace("abc "))
	fmt.Println(replace(" abc"))
	fmt.Println(replace("  a  b  c  "))
	fmt.Println(replace("  aa  b  c  "))
}

//Write a method to replace all spaces in a string with ‘%20’.
func replace(input string) (output string) {
	//找出所有的space的位置
	for i := 0; i < len(input); i++ {
		item := string(input[i])
		if string(input[i]) == " " {
			item = "%20"
		}
		output = output + item
	}
	return output
}
