package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "unicode"
)

func main() {
    // Opening and reading file from
    // https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
    file, err := os.Open("./day1-input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)

    sum := 0
    for scanner.Scan() {
        code := get_calibration_code(scanner.Text())
        fmt.Println("adding code: ", code)
        sum += code;
    }

    fmt.Println("The sum of all calibration codes is: ", sum)
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