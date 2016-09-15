package helloworld

import "fmt"

func Greet(name string) string {
    if len(name) == 0 {
        name = "World"
    }

    return fmt.Sprintf("Hello, %s!", name)
}
