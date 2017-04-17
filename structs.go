package main

import (
        "fmt"
        "encoding/json"
        "os"
)

type Book struct {
        Name string
        Release int
        Author string
        Genre string
        Avaiable bool
}

type Library struct {
        Name string
        Books []Book
}

func main() {
      bookOne := Book{
              Name: "The Hunger Games",
              Release: 2008,
              Author: "Suzanne Collins",
              Genre: "Fiction",
              Avaiable: true,
      }

      bookTwo := Book{
              Name: "A Song of Ice and Fire",
              Release: 2011,
              Author: "George R. R. Martin",
              Genre: "Fantasy",
              Avaiable: false,
      }

      library := Library{Name: "Biblioteca da Faccat", Books: []Book{bookOne, bookTwo}}

      b, err := json.MarshalIndent(library, "", "    ")

      if err != nil {
              fmt.Println(err)
              return
      }

      os.Stdout.Write(b)
}
