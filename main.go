package main

import (
    "github.com/gin-gonic/gin"
    
    "bufio"
    "os"
    "strings"
    "math/rand"
)

func main() {
    initWords()

    r := gin.Default()
    r.GET("/", index)
    r.Run()
}

func index(c* gin.Context) {
    tWord, fWord := getRandomWords()
    c.JSON(200, gin.H {
        "town"  : tWord,
        "fool"  : fWord,
    })
}

var fileName = "easy"
var words []string

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

func getRandomWords() (string, string) {
    index := rand.Intn(len(words))
    pick := strings.Split(words[index], ",")
    return pick[0], pick[1]
}
