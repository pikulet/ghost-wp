package gpack

import (
    "bufio"
    "os"
    "strings"
    "math/rand"
)

var fileName = "easy"

// Downloads the wordpack from github.com/pikulet/ghost-wp.
// The wordpack is saved in a file 'easy'.
// Can be skipped if the user has a pre-defined wordpack
func DownloadWords() error {
    fileURL := "https://github.com/pikulet/ghost-wp/easy"
    resp, err := http.Get(fileURL)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    out, err := os.Create(fileName)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, resp.Body)
    return err
}

var words []string

// Initialises the word list
func Init() { 
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

func SetFileName(name string) {
    fileName = name
}

func GetRandomPair() (string, string) {
    index := rand.Intn(len(words))
    pick := strings.Split(words[index], ",")
    return pick[0], pick[1]
}
