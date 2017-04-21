package main

import (
        "fmt"
        "math"
)

type Circle struct {
        x, y, r float64
}

func (c *Circle) getArea() float64 {
        return math.Pi * c.r * c.r
}

func main() {
        c := Circle{10, 20, 10}

        fmt.Println(c.getArea())
}
