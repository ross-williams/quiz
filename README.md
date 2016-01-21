# quiz


Q: Given a list of words like https://github.com/NodePrime/quiz/blob/master/word.list find the longest compound-word in the list, which is also a concatenation of other sub-words that exist in the list. The program should allow the user to input different data. The finished solution shouldn't take more than one hour. Any programming language can be used, but Go is preferred.


Fork this repo, add your solution and documentation on how to compile and run your solution, and then issue a Pull Request. 

Obviously, we are looking for a fresh solution, not based on others' code.


## Solution by Ross Williams

21 January 2016

The solution is implemented by the program `words.go`.

To run it, use the command:

    go run words.go word.list
    
It should yield the following output:

```
The longest compound word is "antidisestablishmentarianisms".
It can be composed as "an" + "ti" + "dis" + "establishmentarian" + "isms".
```

## A Shorter Solution

22 January 2016

I had an idea and managed to reduce my solution to 43 lines of Go code!
This solution will only work if the input file contains only letters
(or more specifically non pattern matching characters),
but it works on the example input!

The solution is implemented by the program `words_concise.go`.

To run it, use the command:

    go run words_concise.go word.list
    
It should yield the following output:

```
The longest compound word is "antidisestablishmentarianisms".
```

## Comments On This Test

* This test says it should take less than an hour. It took me about
five hours because I am still somewhat inexperienced at Go (I've only
written about 1000 lines total so far) and I had to keep looking things up.

* The fact that it took me five hours makes me wonder whether I missed
a truly simple way to do this. Perhaps there is a Go library that
parses strings using dictionaries. Please let me know if I have
missed something obvious!

* I am a little confused about the interaction of type `string` and
type `[]byte` in Go, and this shows in the code. The code could
probably be simplified by sticking to one or the other.

* Normally I would add checks to check the input file,
but as the specification imposes no constraints on the
input file, apart from that it be a text file, I haven't
bothered.

* I haven't bothered to add a test suite.

## Defects In This Specification

* The term "compound-word" is a compound adjective that is being used
as a noun. It is also overqualified by being undefined. It would be
better to replace it with just "word".

* The term "sub-words" is similarly overqualified and should be
replaced with just "words".

* It is not clear whether a word in the list may be used more than
once when composing a word out of words in the list. For example
if "woolloomooloo", "wool", "loo", and "moo" are in the list,
can "woolloomooloo" be considered to be composed of other words
in the list? It's not clear because it uses "loo" twice.
If this answer to this question is "no", then would it be OK
if "loo" appeared twice in the input list? The specification
is ambiguous without answers to these questions.

* The specification doesn't say whether the input file is in ASCII
or Unicode. I have assumed ASCII, but I think my program will probably
work with Unicode anyway. The nature of the input should be specified
explicitly.

* The sentence "the finished solution shouldn't take more than an hour"
doesn't clarify whether it is referring to programming time or program
execution time. I am assuming programming time.

