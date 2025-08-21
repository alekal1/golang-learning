package main

func ReverseString(word string) string {
	var res string
	for i := len(word) - 1; i >= 0; i-- {
		res += string(word[i])
	}
	return res
}
