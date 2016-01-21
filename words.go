//------------------------------------------------------------------------------

/*

Author : Ross Williams.
Date   : 21 January 2016.

This program was created as a solution to a test
provided by NodePrime. Here is the exact specification:

   Given a list of words like
      https://github.com/NodePrime/quiz/blob/master/word.list
   find the longest compound-word in the list,
   which is also a concatenation of other sub-words
   that exist in the list. The program should allow
   the user to input different data. The finished
   solution shouldn't take more than one hour. Any
   programming language can be used, but Go is
   preferred.

This is a completely self-contained program. You can run it with simply:

   go run words.go <argument_file>

*/

//------------------------------------------------------------------------------

package main

import "fmt"
import "os"
import "io/ioutil"
import "strings"

//------------------------------------------------------------------------------

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//------------------------------------------------------------------------------

// Checks that there is exactly one argument and terminates with an error if not.
// Returns the first argument.
//
func check_and_get_argument() string {
	if len(os.Args) != 2 {
		fmt.Println("Usage: words <input_file_of_words>")
		os.Exit(1)
	}
	return os.Args[1]
}

//------------------------------------------------------------------------------

// Reads in the argument file and returns a list of strings
// being the lines in the file. If the input is OK, this is a list
// of words, but it doesn't actually matter what the lines contain.
//
func read_in_word_list(fileName string) []string {
	file_contents, err := ioutil.ReadFile(fileName)
	check(err)
	file_contents_s := string(file_contents[:len(file_contents)])
	return strings.Split(file_contents_s, "\n")
}

//------------------------------------------------------------------------------

// Accepts an array of words and returns a set of words implemented as a map.
//
func make_dictionary(words []string) map[string]bool {
	dict := make(map[string]bool)
	for _, word := range words {
		dict[word] = true
	}
	return dict
}

//------------------------------------------------------------------------------

// Accepts a word and a dictionary and attempts to compose the word from
// words in the dictionary. If it can, it returns a list of words from
// the dictionary that form the argument word. If it can't, it returns
// the empty list.
//
// If a word has multiple parsings, this function will return a parsing
// whose first word has the minimum length. This ensures that this
// function won't return a parse of just the word itself unless there
// are no multiple-word parsings.
//
// Assumption: A particular word in the dictionary
// may be used more than once within a composition.
// i.e. words in the dictionary are not "consumed" during composition.
//
func compose(dict map[string]bool, word []byte) []string {

	// The approach is to find all words in the dictionary
	// that are prefixes of the argument word and then recurse
	// to perform the rest of the word's parsing.

	// Search for prefixes by prefix length.
	//
	word_len := len(word)
	for prefix_len := 1; prefix_len <= word_len; prefix_len++ {

		prefix := string(word[0:prefix_len])

		// Lookup the prefix in the dictionary.
		// If the prefix does not match, try the next length of prefix.
		if !dict[prefix] {
			continue
		}

		prefix_word_list := []string{prefix}

		// If the prefix consumed the entire string, then the parse succeeded.
		if prefix_len == word_len {
			return prefix_word_list
		}

		// If the prefix did not consume the entire string,
		// attempt to parse the suffix by recursing.
		suffix_word_list := compose(dict, []byte(word[prefix_len:word_len]))

		// If the suffix parsed, then the parse succeeded.
		if len(suffix_word_list) > 0 {
			return append(prefix_word_list, suffix_word_list...)
		}

		// If the suffix parse failed, this parse has failed, but we
		// still need to go around to try the remaining prefix lengths.
	}

	return []string{}
}

//------------------------------------------------------------------------------

// Prints its argument list of words to standard output with "+" formatting.
//
func print_explanatory_word_list(word_list []string) string {
	r := ""
	join := ""
	for _, word := range word_list {
		r += join + "\"" + word + "\""
		join = " + "
	}
	return r
}

//------------------------------------------------------------------------------

func main() {

	argumentFileName := check_and_get_argument()

	words := read_in_word_list(argumentFileName)

	dict := make_dictionary(words)

	longest_word_so_far := ""
	longest_word_list := []string{}

	// Analyse each word in the list.
	//
	for _, word := range words {

		// This is a speed optimisation only. There's no need to
		// assess words that are not longer than the current best.
		if len(word) <= len(longest_word_so_far) {
			continue
		}

		// Can the longer-than-best word be composed of other words?
		// If so, it's the new best.
		//
		list := compose(dict, []byte(word))
		if len(list) > 1 {
			longest_word_so_far = word
			longest_word_list = list
		}
	}

	if len(longest_word_so_far) == 0 {
		fmt.Println("There is no word in the list that can be composed from the other words.")
		os.Exit(1)
	}

	fmt.Println("The longest compound word is \"" + longest_word_so_far + "\".")
	fmt.Println("It can be composed as " + print_explanatory_word_list(longest_word_list) + ".")
}

//------------------------------------------------------------------------------
