# :ghost: Ghost Wordpack :ghost:
Just a wordpack since I couldn't find any around :) 

Ghost is a word game that involves hiding a secret word from a minority in the group. There will also be a second secret word, but this word is just a decoy.

In most Ghost games, players want a relation between the two words, for example antonyms or words in the same semantic field. As such, the repository aims to provide two related words for the Ghost game.

Currently, the word pairs are hard-coded in the `easy` file. In future, it would be nice to use Natural Language Processing to quickly decide on what a good word pair for the game would be.

If you want to host your own wordpack server, this repository was configured to deploy with Heroku :)

## :firecracker: Usage

If you just need the words, you can directly use the `easy` file, or make an API call to 

`pikulet.herokuapp.com/ghost`

If you want to use this go mod, do

`go get -u github.com/pikulet/ghost-wp`

```
import github.com/pikulet/ghost-wp
...
townWord, foolWord := ghostpack.GetRandomPair()
```

## :memo: Contributing

I'd definitely look forward to suggestions for the word pack. You can make a pull request or create an issue. Due credit will be given!

## :seedling: Extensions

Currently working on a word splash the gives a few random words (with no relation). I'm looking for a random word generator for English words, but there aren't many around. I've played with https://github.com/tjarratt/babble but this gomod isn't compatible with hosting on a remote server. Do let me know if you have great suggestions.
