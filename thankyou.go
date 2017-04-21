package main

import (
        "fmt"
)

func main() {
        urls := []string{"github.com",
                "twitter.com",
                "facebook.com",
                "instagram.com"}

        for _, url := range urls {
                fmt.Printf("%s/fernandoporazzi\n", url)
        }

}
