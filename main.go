package ghost

import (
    "bufio"
    "os"
    "strings"
    "math/rand"
)

func InitPairs(filename string, *words []string) { 
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
    index := rand.Intn(len(words))
    pick := strings.Split(words[index], ",")
    return pick[0], pick[1]
}
