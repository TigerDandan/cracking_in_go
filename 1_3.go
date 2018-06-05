package main

import "fmt"

func main() {
	fmt.Println(RemoveDuplicatedChars_Me(""))
	fmt.Println(RemoveDuplicatedChars_Me("aaaa"))
	fmt.Println(RemoveDuplicatedChars_Me("ab b "))
	fmt.Println(RemoveDuplicatedChars_Me("aaabbb"))
	fmt.Println(RemoveDuplicatedChars_Me("abcabdcabcdmmmmmm"))
}

//Design an algorithm and write code to remove the duplicate characters in a string without using any additional buffer.
//NOTE: One or two additional variables are fine. An extra copy of the array is not.
//FOLLOW UP: Write the test cases for this method.
func RemoveDuplicatedChars_Me(input string) (output string) {
	duplicated := make([]bool, len(input))
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				duplicated[j] = true
				continue
			}

		}
	}
	for i := 0; i < len(input); i++ {
		if duplicated[i] == false {
			output = output + string(input[i])
		}
	}
	return output
}
