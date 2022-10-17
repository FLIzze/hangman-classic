package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"bufio"
	"os"
	"time"
)

type HangManData struct {
	Word             string // Word composed of '_', ex: H_ll_
	ToFind           string // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	hangmandata := HangManData{}

	hangmandata.ToFind = chooseWord();
	hangmandata.Word = createWord(hangmandata.ToFind);
	hangmandata.Attempts = 0

	fmt.Println("Good Luck, you have 10 attempts.")
	for {
		var goodLetter bool
		hangmandata.Word, goodLetter = checkLetter(askLetter(hangmandata.Word), hangmandata.Word, hangmandata.ToFind)
		if !goodLetter {
			hangmandata.Attempts++
		}
		if hangmandata.Word == hangmandata.ToFind {
			break
		}
		if hangmandata.Attempts >= 10 {
			loose()
		}
	}
	fmt.Println("The word was " + hangmandata.ToFind)
	win()
}

func chooseWord() string { // Function to select word needed to find in a special file
	args := os.Args[1:]
	
	if len(args) != 1 {
		fmt.Println("/!\\ You need to select a file...")
		os.Exit(0)
	}

	content, _ := ioutil.ReadFile(args[0])

	var words []string
	var word string 

	for _, i := range content {
		if i == 10 {
			words = append(words, word)
			word = ""
		} else {
			word = word + string(i)
		}
	}
	return words[rand.Intn(len(words))]
}

func createWord(word string) string { // Function to create the word we display
	var secretword []string
	var words []string

	for _, i := range word {
		words = append(words, string(i))
		secretword = append(secretword, "_")
	}

	var checkreveal bool

	for i := (len(word) / 2 - 1); i >= 0; i-- {
		rdm := rand.Intn(len(secretword) - 1)
		for _, j := range secretword {
			if words[rdm] == j {
				i++
				checkreveal = true
				break
			}
		}
		if !checkreveal {
			for _, j := range words {
				if words[rdm] == j {
					i--
				}
			}
			if i >= -1 {
				for k, j := range words {
					if words[rdm] == j {
						secretword[k] = j
					}
				}
			}
		} 
	}

	word = ""

	for _, j := range secretword {
		word = word + string(j)
	}

	return word
}

func askLetter(word string) string {
	reader := bufio.NewReader(os.Stdin)
	for _, j := range word {
		fmt.Print(string(j) + " ")
	}
	fmt.Println("\n")
	fmt.Print("Choose: ")
	letter, _ := reader.ReadString('\n')
	for {
		if len(letter) != 3 {
			fmt.Println("You need to choose a letter, not a word")
			fmt.Print("Choose: ")
			letter, _ = reader.ReadString('\n')
		} else {
			break
		}
	}
	return letter
}

func checkLetter(letter string, word string, tofind string)(string, bool) {
	var newWord string
	var goodLetter bool

	for i, j := range word {
		if string(tofind[i]) == string(letter[0]) && string(j) != letter {
			newWord += string(letter[0])
			goodLetter = true
		} else {
			newWord += string(j)
		}
	}
	return newWord, goodLetter
}

func win() {
	fmt.Println("\n-------------------------------")
	fmt.Println("            YOU WIN            ")
	fmt.Println("-------------------------------\n")
}

func loose() {
	fmt.Println("\n-------------------------------")
	fmt.Println("           YOU LOOSE           ")
	fmt.Println("-------------------------------\n")
	os.Exit(0)
}