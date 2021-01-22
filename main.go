package ghostwords

import (
    "bufio"
    "os"
    "strings"
    "math/rand"
)

var isInit = false
var fileName = "easy"
var words []string

func initPairs() { 
    file, err := os.Open(fileName)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        words = append(words, scanner.Text())
    }
}

func GetRandomPair() (string, string) {
    if !isInit {
        initPairs()
        isInit = true
    }

    index := rand.Intn(len(words))
    pick := strings.Split(words[index], ",")
    return pick[0], pick[1]
}
