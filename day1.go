package main

import "fmt"
import "unicode"

func main() {
    fmt.Println("Hello, Joshy")

    result := get_calibration_code("a1bcdefg 1ABCDEFG1")
    fmt.Println(result)
}

func get_calibration_code(s string) int {
    // Find the first digit
    first := 0;
    for _, c := range s {
        if unicode.IsDigit(c) {
            //https://stackoverflow.com/a/21322694/4184410
            first = int(c - '0')
            break
        }
    }

    // It is not trivial to reverse a string, so loop the whole string instead
    // https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go/
    last := 0
    for _, c := range(s) {
        if unicode.IsDigit(c) {
            last = int(c - '0')
        }
    }

    return first * 10 + last
}