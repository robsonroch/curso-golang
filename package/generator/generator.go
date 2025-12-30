package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - canal somente-leitura
func titulo(urls ...string) <-chan string {
    c := make(chan string)
    done := make(chan bool)
    for _, url := range urls {
        go func(url string) {
            resp, err := http.Get(url)
            if err != nil {
                c <- fmt.Sprintf("Erro ao buscar %s: %v", url, err)
                done <- true
                return
            }
            defer resp.Body.Close()
            html, err := io.ReadAll(resp.Body)
            if err != nil {
                c <- fmt.Sprintf("Erro ao ler %s: %v", url, err)
                done <- true
                return
            }

            r, _ := regexp.Compile("<title>(.*?)</title>")
            matches := r.FindStringSubmatch(string(html))
            if len(matches) > 1 {
                c <- matches[1]
            } else {
                c <- fmt.Sprintf("Título não encontrado para %s", url)
            }
            done <- true
        }(url)
    }
    go func() {
        for i := 0; i < len(urls); i++ {
            <-done
        }
        close(c)
    }()
    return c
}

func main() {
    t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
    t2 := titulo("https://www.amazon.com", "https://www.youtube.com")
    fmt.Println("Primeiros:", <-t1, "|", <-t2)
    fmt.Println("Segundos:", <-t1, "|", <-t2)
}