package main

import "fmt"
import "os"
import "io/ioutil"
import "strings"
import "regexp"
import "bytes"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: words_concise <input_file_of_words>")
		os.Exit(1)
	}
	file_contents, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	file_contents_s := string(file_contents[:len(file_contents)])
	words := strings.Split(file_contents_s, "\n")
	var pattern bytes.Buffer
	pattern.WriteString("\\A(")
	or_str := ""
	for _, word := range words {
		if len(word) > 0 {
			pattern.WriteString(or_str + word)
			or_str = "|"
		}
	}
	pattern.WriteString("){2,}\\z")
	dict_regexp := regexp.MustCompile(pattern.String())
	longest_word_so_far := ""
	for _, word := range words {
		if len(word) > len(longest_word_so_far) && dict_regexp.MatchString(word) {
			longest_word_so_far = word
		}
	}
	if len(longest_word_so_far) == 0 {
		fmt.Println("There is no word in the list that can be composed from the other words.")
	} else {
		fmt.Println("The longest compound word is \"" + longest_word_so_far + "\".")
	}
}
