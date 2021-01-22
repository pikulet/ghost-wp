package ghost

import (
    "bufio"
    "os"
    "strings"
    "math/rand"
)

var init = false
var fileName = "easy"
var words []string

func InitPairs(fileName string) { 
    file, err := os.Open(fileName)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        *words = append(*words, scanner.Text())
    }
}

func GetRandomPair() (string) {
    index := rand.Intn(len(words))
    pick := strings.Split(words[index], ",")
    return pick[0], pick[1]
}
