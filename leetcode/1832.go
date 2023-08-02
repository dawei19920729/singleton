package main

import "fmt"

func main() {
	sentence := "thequickbrownfoxjumpsoverthelazydog"
	bool := checkIfPangram(sentence)
	fmt.Println(bool)

}

type void struct {
}

func checkIfPangram(sentence string) bool {
	var member void
	var pangramMap = make(map[byte]void)
	for i := 0; i < len(sentence); i++ {
		pangramMap[sentence[i]] = member
	}
	if len(pangramMap) == 26 {
		return true
	} else {
		return false
	}

}
