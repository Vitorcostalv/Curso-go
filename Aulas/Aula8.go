package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    // Escrever em um arquivo
    data := []byte("Hello, file!")
    ioutil.WriteFile("hello.txt", data, 0644)

    // Ler de um arquivo
    content, _ := ioutil.ReadFile("hello.txt")
    fmt.Println("File Content:", string(content))

    // Fazer uma requisição HTTP
    resp, _ := http.Get("https://api.github.com")
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("GitHub API Response:", string(body))
    resp.Body.Close()
}
