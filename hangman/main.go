package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type HangManData struct {
	Word        string     // Word composed of '_', ex: H_ll_
	ToFind      string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts    int        // Number of attempts left
	CountMin    int        // Line where we'll print the hangman in hangman.txt.
	CountMax    int        // Same thing.
	UsedLetters [10]string //Letters we already used.
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	hangmandata := &HangManData{}

	args := os.Args[1:]

	if args[0] == "--startWith" {
		getSavedFile(hangmandata, args[1])
		fmt.Println("-------------------------------------")
		fmt.Println("Welcome back you have", 10-hangmandata.Attempts, "attempts.")
	} else {
		hangmandata.ToFind = chooseWord(hangmandata)
		hangmandata.Word = createWord(hangmandata.ToFind)
		hangmandata.Attempts = 0
		hangmandata.CountMax = 7
		hangmandata.CountMin = 0
		fmt.Println("-------------------------------------")
		fmt.Println("Good Luck, you have", 10-hangmandata.Attempts, "attempts.")
	}

	for {
		goodLetter := 1
		hangmandata.Word, goodLetter = checkLetter(askLetter(hangmandata.Word, hangmandata), hangmandata.Word, hangmandata.ToFind, hangmandata)
		if goodLetter != 0 {
			printHangmanFile(hangmandata)
			hangmandata.Attempts += goodLetter
			hangmandata.CountMax += 8
			hangmandata.CountMin += 8
		}
		if hangmandata.Word == hangmandata.ToFind {
			win(hangmandata.ToFind)
			break
		}
		if hangmandata.Attempts >= 10 {
			loose(hangmandata.ToFind)
			break
		}
		fmt.Println("-------------------------------------")
		fmt.Println("\nGood Luck, you have:", 10-hangmandata.Attempts, "attemps left.")
	}
}

func chooseWord(hangmandata *HangManData) string { // Function to select word needed to find in a special file
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

	for i := (len(word)/2 - 1); i >= 0; i-- {
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

func askLetter(word string, hangmandata *HangManData) string {
	var letter string
	for _, j := range word {
		fmt.Print(string(j) + " ")
	}
	fmt.Println()
	fmt.Print("Choose: ")
	fmt.Scanln(&letter)
	if letter == "STOP" {
		saveFile(hangmandata)
		os.Exit(0)
	}
	for {
		if letter >= "a" && letter <= "z" || len(letter) > 2 {
			break
		} else {
			fmt.Println("\nThis character is not allowed. Try again !\n")
			fmt.Println("Choose: ")
			fmt.Scanln(&letter)
			if letter == "STOP" {
				saveFile(hangmandata)
				os.Exit(0)
			}
		}
	}
	return letter
}

func askWord(letters string, word string, tofind string) string {
	var newletters string
	for _, j := range letters {
		letters = ""
		if string(j) != " " {
			newletters += string(j)
		}
	}
	if newletters == tofind+"\n" {
		win(tofind)
		os.Exit(0)
	}
	return word
}

func checkLetter(letter string, word string, tofind string, hangmandata *HangManData) (string, int) {
	var newWord string
	goodLetter := 1

	if len(letter) > 1 {
		return askWord(letter, word, tofind), 2
	} else {
		for i, j := range word {
			if string(tofind[i]) == string(letter[0]) && string(j) != letter {
				newWord += string(letter[0])
				goodLetter = 0
			} else {
				newWord += string(j)
			}
		}
	}
	for _, j := range hangmandata.UsedLetters {
		if letter == j {
			goodLetter = 0
			break
		}
	}
	if len(letter) == 1 {
		usedLetters(hangmandata, letter)
	}
	return newWord, goodLetter
}

func printHangmanFile(hangmandata *HangManData) { //Print the current state of the hangman
	file, err := os.Open("files/hangman.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line int
	fmt.Println()
	for scanner.Scan() {
		if line >= hangmandata.CountMin && line <= hangmandata.CountMax {
			fmt.Println(scanner.Text())
		}
		line++
	}
}

func usedLetters(hangmandata *HangManData, letter string) { //Print an array of letters we've used
	var count bool
	for _, i := range hangmandata.UsedLetters {
		if i == letter {
			count = true
		}
	}
	if !count {
		hangmandata.UsedLetters[hangmandata.Attempts] = letter
	}
	fmt.Println(" ___")
	for _, j := range hangmandata.UsedLetters {
		if j >= "a" && j <= "z" {
			fmt.Println("|", j, "|")
		}
	}
	fmt.Println(" ---")
}

func win(tofind string) {
	fmt.Println("The word was " + tofind)
	fmt.Println("-------------------------------------")
	fmt.Println("|           - CONGRATS !            |")
	fmt.Println("-------------------------------------")

}

func loose(tofind string) {
	fmt.Println("The word was " + tofind)
	fmt.Println("------------------------------------")
	fmt.Println("|          - TRY AGAIN !           |")
	fmt.Println("-------------------------------------")
}

func saveFile(hangmandata *HangManData) { // STOP to use the function, will save progress
	content, _ := json.Marshal(hangmandata)
	err := ioutil.WriteFile("save.txt", content, 0777)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Game saved in save.txt.")
}

func getSavedFile(hangmandata *HangManData, save string) { // Start from a saved file
	content, err := ioutil.ReadFile(save)
	if err != nil {
		fmt.Print(err)
	}
	json.Unmarshal(content, hangmandata)
}
