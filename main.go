package main

import (
    "github.com/gin-gonic/gin"
    "github.com/tjarratt/babble"
    
    "bufio"
    "os"
    "strings"
    "math/rand"
)

func main() {
    initWords()
    initBabbler()

    r := gin.Default()
    r.GET("/", pair)
    r.GET("/splash", splash)
    r.Run()
}

func pair(c* gin.Context) {
    tWord, fWord := getRandomPair()
    c.JSON(200, gin.H {
        "town"  : tWord,
        "fool"  : fWord,
    })
}

func splash(c* gin.Context) {
    c.JSON(200, gin.H {
        "splash"    : getRandomSplash(),
    })
}

var fileName = "easy"
var words []string

var babbler = babble.NewBabbler() 

func initWords() { 
    file, err := os.Open(fileName)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    words = lines
}

func getRandomPair() (string, string) {
    index := rand.Intn(len(words))
    pick := strings.Split(words[index], ",")
    return pick[0], pick[1]
}


func initBabbler() {
    babbler.Count = 10
}

func getRandomSplash() []string {
    return strings.Split(babbler.Babble(), babbler.Separator)
}
