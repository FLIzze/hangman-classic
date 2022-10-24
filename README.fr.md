# Hangman-game

![english🇬🇧](https://ytrack.learn.ynov.com/git/fmael/hangman-classic/src/branch/master/README.md)

## Qu'est-ce que c'est?

Ce projet est basé sur le jeu du pendu.

## Comment jouer?

Ce pendu contient toutes les fonctionnalités de base du jeu normal, et quelques unes de plus, elles sont toutes énumérées juste en **dessous**.

* Le jeu commence presque comme une partie de pendu basique, vous allez avoir un mot a trouver et quelques lettres seront visibles dès le début pour vous aidez. Faites de votre mieux pour gagner en trouvant les lettres une par une ou le mot en entier si vous pensez l'avoir. Si vous vous trompez sur une lettre vous perdez **un** essai, si vous vous trompez sur un mot vous perdez **deux** essais. En tout vous avez 10 essais bonne chance.

* Si vous n'avez pas le temps de finir votre parti ne vous en faites pas! Vous pouvez la mettre en pause avec la commande `STOP` et repartir de cette sauvegarde en utilisant `hangman/main.go --startWith save.txt`.


## Installation

Vous n'avez qu'à copier/coller ces commandes dans votre terminal.
```shell
##pour télécharger le jeu
git clone github.com/FLIzze/hangman-game 
##pour commencer une nouvelle partie
go run main.go files/words1.txt
##pour continuer depuis une sauvegarde
hangman/main.go --startWith save.txt
```

## À quoi ça ressemble?

![](https://i.imgur.com/RC8pmmy.png)
![](https://i.imgur.com/83rAV0c.png)


## Team

- Maël FATH 
- Alexandre BEL
- Rémy BEHAGUE
- Joel ANCEL