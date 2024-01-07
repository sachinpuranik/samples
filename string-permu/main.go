package main

import (
	"fmt"
)

func substring(content string, part string,index int) {
	if (index >= len(content)) {
		return;
	}
	sub := part + content[index:index+1]


	substring(content, sub, index + 1)

	permutation("", sub)
}

func permutation(prefix, str string) {
	n := len(str)
	if (n == 0) {
		fmt.Println(prefix)
	} else {
		for i := 0; i < n; i++ {
			permutation(prefix + str[i:i+1], str[0: i] + str[i + 1: n])
		}
	}
}


func main(){
	var s = "abcd";
	// for i := 0; i < len(s); i++ {
		substring(s, "", 2)
	// }
}