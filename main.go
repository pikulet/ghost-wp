package ghostwords

import (
    "bufio"
    "os"
    "strings"
    "math/rand"
)

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

func SetFileName(name string) {
    fileName = name
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
