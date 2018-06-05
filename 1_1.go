package main

import "fmt"

//Implement an algorithm to determine if a string has all unique characters. What if you can not use additional data structures?
func main() {
	//fmt.Println(HasDuplicatedChar_On1(""))
	//fmt.Println(HasDuplicatedChar_On1("abc"))
	//fmt.Println(HasDuplicatedChar_On1("!@#$%^&*()"))
	//fmt.Println(HasDuplicatedChar_On1("  "))
	fmt.Println(HasDuplicatedChar_On1("abdced"))
	fmt.Println(HasDuplicatedChar_On1("abcdefghijklmnopqrstuvwxyza"))
	//fmt.Println(HasDuplicatedChar_On1("abcdega"))
	//fmt.Println(HasDuplicatedChar_On1("abcdefe"))
	//fmt.Println(HasDuplicatedChar_On1("121"))
	//fmt.Println(HasDuplicatedChar_On1("1223"))
	//fmt.Println(HasDuplicatedChar_On1("1#a@#"))
	//
	//fmt.Println(HasDuplicatedChar_On1("!@#$*^&*()"))
	//fmt.Println(HasDuplicatedChar_On1(" !@#$a^&*() "))
	//fmt.Println(HasDuplicatedChar_On1("!@#$a^&* () "))
}

// Check every char of the string with every other char of the string for duplicate occurrences.
// This will take O(n^2) time and no space.
func HasDuplicatedChar_On2(input string) bool {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				return true
			}
		}
	}
	return false
}

// For simplicity, assume char set is ASCII(if not, we need to increase the storage size.
// The rest of the logic would be the same). NOTE: This is a great thing to point out to your interviewer!
// ASCII 一共有256个， http://ee.hawaii.edu/~tep/EE160/Book/chap4/subsection2.1.1.1.html
// Time complexity is O(n), where n is the length of the string, and space complexity is O(n).
func HasDuplicatedChar_On1On1(input string) bool {
	charSet := make([]bool, 256)
	for i := 0; i < len(input); i++ {
		local := input[i]
		if charSet[local] == true {
			return true
		}
		charSet[local] = true
	}
	return false
}

// 不使用任何其他的数据结构检查一个字符串中是否有重复字符
// 实现空间有效算法以确定（从‘a’到’z’的字符）字符串是否具有所有唯一字符。不允许使用数组，散列等附加数据结构。
// 时间复杂度：O（n）
// 思路：因为从‘a’-‘z’共有26个字符，一个int有32位，如果将‘a’对应到int的第一位，‘b’放到第二位，只要判断从0到31位没有重复，字符串字符串中的所有字符都是唯一的，否则不唯一。
// We can reduce our space usage a little bit by using a bit vector.
// We will assume, in the below code, that the string is only lower case ‘a’ through ‘z’.
// This will allow us to use just a single int
func HasDuplicatedChar_On1(input string) bool {
	checker := 0
	for i := 0; i < len(input); i++ {
		val := input[i] - 'a'
		//左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。
		//把1左移val位，即 1 << val = 2的val次方
		fmt.Println("check is:", checker)
		fmt.Println("1<<val is:", 1<<val)
		//将val放到checker中对应的位置
		//如果有重复则返回true
		if (checker & (1 << val)) > 0 {
			return true
		}
		//将val放到checker中对应的位置
		checker = checker | (1 << val)
		fmt.Println()
	}
	return false
}

//If we are allowed to destroy the input string,
// we could sort the string in O(n log n) time and then linearly check the string for neighboring characters that are identical.
// Careful, though - many sorting algorithms take up extra space.
