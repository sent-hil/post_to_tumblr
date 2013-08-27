Example of using struct tags in Go:

    package main

    import (
      "fmt"
      "reflect"
    )

    type My struct {
      A int `abc:"imatag"` // define a custom tag
    }

    func main() {
      x := &My{A: 34}

      // get field with reflect
      f, _ := reflect.TypeOf(x).Elem().FieldByName("A")

      // check out its tags
      fmt.Println(f.Tag)
    }
