package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter file name")
		return
	}

	fileName := os.Args[1]
	resultName := os.Args[2]

	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	// converting file (byte) to string
	fileText := string(file)

	// string manipulation
	word := strings.Fields(fileText)

	// iterating over each word
	i := 0
	for i < len(word) {
		if i > 0 {
			word = ChangeArticles(word)

			if word[i] == "(up)" {
				word[i-1] = strings.ToUpper(word[i-1])
				// Remove the marker "(up)" from the slice
				word = append(word[:i], word[i+1:]...)

			} else if word[i] == "(low)" {
				word[i-1] = strings.ToLower(word[i-1])
				word = append(word[:i], word[i+1:]...)

			} else if word[i] == "(cap)" {
				word[i-1] = strings.Title(word[i-1])
				word = append(word[:i], word[i+1:]...)

			} else if word[i] == "(hex)" {
				word[i-1] = HexToInt(word[i-1])
				word = append(word[:i], word[i+1:]...)

			} else if word[i] == "(bin)" {
				word[i-1] = BinToInt(word[i-1])
				word = append(word[:i], word[i+1:]...)

				// upper with number
			} else if word[i] == "(up," {
				b := strings.Trim(string(word[i+1]), ")")
				number, _ := strconv.Atoi(string(b))
				for j := 1; j <= number; j++ {
					word[i-j] = strings.ToUpper(word[i-j])
				}
				word = append(word[:i], word[i+2:]...)

				// lower with number
			} else if word[i] == "(low," {
				b := strings.Trim(string(word[i+1]), ")")
				number, _ := strconv.Atoi(string(b))
				for j := 1; j <= number; j++ {
					word[i-j] = strings.ToLower(word[i-j])
				}
				word = append(word[:i], word[i+2:]...)

				// cap with number
			} else if word[i] == "(cap," {
				b := strings.Trim(string(word[i+1]), ")")
				number, _ := strconv.Atoi(string(b))
				for j := 1; j <= number; j++ {
					word[i-j] = strings.Title(word[i-j])
				}
				word = append(word[:i], word[i+2:]...)
			} else {
				i++
			}
		} else {
			i++
		}
	}

	word = Punctuations(word)

	// join words back to string
	newText := strings.Join(word, " ") + "\n"
	trimmedText := strings.TrimSpace(newText)

	// write newText to file
	err = os.WriteFile(resultName, []byte(trimmedText), 0o644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func HexToInt(hexadecimal string) string {
	num, _ := strconv.ParseInt(hexadecimal, 16, 64)
	return fmt.Sprint(num)
}

func BinToInt(binary string) string {
	num, _ := strconv.ParseInt(binary, 2, 64)
	return fmt.Sprint(num)
}

func ChangeArticles(words []string) []string {
	vowels := []string{"a", "e", "i", "o", "u", "h"}
	for i, word := range words {
		for _, vowel := range vowels {
			if word == "a" && string(words[i+1][0]) == vowel {
				words[i] = "an"
			} else if word == "A" && string(words[i+1][0]) == vowel {
				words[i] = "An"
			}
		}
	}
	return words
}

func Punctuations(words []string) []string {
	puncs := []string{",", ".", "!", "?", ":", ";"}
	// for punctuations at the end of a string
	for i, word := range words {
		for _, punc := range puncs {
			if string(word[0]) == punc && words[i] == words[len(words)-1] {
				words[i-1] = words[i-1] + word
				words = words[:len(words)-1]
			}
		}
	}
	// punc in the beginning of a word but in the middle of the string
	for i, word := range words {
		for _, punc := range puncs {
			if string(word[0]) == punc && string(word[len(word)-1]) != punc {
				words[i-1] = words[i-1] + punc
				words[i] = word[1:]
			}
		}
	}
	// punc in the middle of the string
	for i, word := range words {
		for _, punc := range puncs {
			if string(word[0]) == punc && string(word[len(word)-1]) == punc {
				words[i-1] = words[i-1] + word
				words = append(words[:i], words[i+1:]...)
			}
		}
	}
	// first apostrophe
	count := 0
	for i, word := range words {
		if word == "'" && count == 0 {
			words[i+1] = word + words[i+1]
			words = append(words[:i], words[i+1:]...)
			count++
		}
	}
	// for the second apostrophe
	for i, word := range words {
		if word == "'" {
			words[i-1] = words[i-1] + word
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}
