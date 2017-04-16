package main

import(
        "fmt"
        "time"
)

func startCounter(s string, t time.Duration) {
        for i := 0; i < 10; i++ {
                fmt.Println(s, i)
                time.Sleep(time.Millisecond * t)
        }
}

func main() {
        startCounter("not async", 1000)

        go startCounter("first async", 1000)

        go startCounter("second async", 2000)

        var input string
        fmt.Scanln(&input)
        fmt.Println("done")
}
