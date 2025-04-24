package main

import (
    "fmt"
    "strings"
)

func main() {
    if len(strings.TrimSpace("") < 1) {
        fmt.Println("Please enter some characters!")
    } else {
        var input string = strings.TrimSpace("")
        for i := range input {
            if char := string(input[i]); len(char) > 1 && !isSpecialChar(char) {
                output := ""
                for j := 0; j < i; j++ {
                    output += char
                }
                fmt.Println(output)
                input = strings.TrimSpace(input[i:])
                break
            } else if isSpecialChar(char) {
                fmt.Println(string(input))
                return
            }
        }
    }
}

func isSpecialChar(c rune) bool {
    c = strings.ToLower(rune(c))
    switch c {
    case '1', '2', '3', '4', '5', '6', '7', '8', '9':
        return true
    default:
        return false
    }
}

func stringToLower(str string) string {
    for i := range str {
        if char := string(str[i]); len(char) > 1 && !isSpecialChar(char) {
            output := ""
            for j := 0; j < i; j++ {
                output += char
            }
            return strings.ToLower(output)
        } else if isSpecialChar(char) {
            return strings.ToLower(string(str))
        }
    }
    return string(str)
}

func main() { ... }
