package main

import "fmt"

func main() {
	//s1 := "ab cd"
	//reverseString2(s1)
	//fmt.Println(s1)
	fmt.Println(ReverseString2("ab-+cd"))
	fmt.Println(ReverseString2("a aa"))
	fmt.Println(ReverseString2("a +aa"))
	fmt.Println(ReverseString2("a +中文"))
}

//Write code to reverse a C-Style String.
//c-style 是一个指针，或者char数组，字符串的末尾有\0字符

//(C-String means that “abcd” is represented as five characters, including the null character.)
// 这种写法不能处理unicode，中文字符，fmt.Println(reverseString1("a +中文"))报错
func ReverseString1(input string) (output string) {
	l := len(input)
	for i := 1; i <= len(input); i++ {
		output = output + string(input[l-i])
	}
	return output
}

//我们知道golang中的string类型存储的字符串是不可变的，
// 如果要修改string内容需要将string转换为[]byte或[]rune，并且修改后的string内容是重新分配的。
//In Go1 rune is a builtin type.
func ReverseString2(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//This is a classic interview question. The only “gotcha” is to try to do it in place, and to be careful for the null character.
//func reverseString2(input *string) {
//}
