# Hangman-game

## What is it?

This project is based on the famous HangMan Game.

## How to play?


This HangMan Game contains all the basic features of the classic games, and some others are added, they are all enumerated just **below**.

* The game will start almost like a basic hangman game, you will get a word to find yet some letters will be visible to help you. Do your best to win by typing letters one by one or the entire word if you feel like you've got it. If a letter isnt in the word you will lose **one** attempt, if your word isnt the good one you will lose **two** attempts. Overall you have 10 attempts good luck.

* If you don't have time to finish your game don't worry! You can **pause** it with the command `STOP` and start from this save by using `hangman/main.go --startWith save.txt`.


## Installation

Just copy and paste the following commands in your terminal
```shell
##to download the game
git clone https://ytrack.learn.ynov.com/git/fmael/hangman-classic.git
##to start a brand new game 
cd hangman-classic
go run hangman/main.go files/words.txt
##to start from a save
go run hangman/main.go --startWith save.txt
```

## What does it look like?

![](https://i.imgur.com/RC8pmmy.png)
![](https://i.imgur.com/83rAV0c.png)


## Team

- Maël FATH 
- Alexandre BEL
- Rémy BEHAGUE
- Joel ANCEL